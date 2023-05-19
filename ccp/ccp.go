package ccp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CCP struct {
	servers []string
	path    string
	logging bool

	*http.Client
}

type Response struct {
	Content                 string `json:"Content"`
	CreationMethod          string `json:"CreationMethod"`
	Address                 string `json:"Address"`
	Safe                    string `json:"Safe"`
	UserName                string `json:"UserName"`
	Name                    string `json:"Name"`
	PolicyID                string `json:"PolicyID"`
	DeviceType              string `json:"DeviceType"`
	CPMDisabled             string `json:"CPMDisabled"`
	Folder                  string `json:"Folder"`
	PasswordChangeInProcess string `json:"PasswordChangeInProcess"`
}

type Option func(p *CCP)

func WithLogging(b bool) Option {
	return func(p *CCP) {
		p.logging = b
	}
}

func New(servers []string, appID string, options ...Option) (*CCP, error) {
	ccp := &CCP{
		servers: servers,
		path:    fmt.Sprintf("/AIMWebService/api/Accounts?AppId=%s", appID),
		Client:  http.DefaultClient,
	}

	for _, option := range options {
		option(ccp)
	}

	return ccp, nil
}

func (c *CCP) GetPassword(options ...ApiOption) (*Response, error) {
	ctx, cancel := context.WithCancel(context.Background())

	result := make(chan *Response)
	fails := make(chan error)

	path, err := c.buildPath(c.path, options...)
	if err != nil {
		c.logIfEnabled(err.Error())
		cancel()
		return nil, err
	}

	for _, server := range c.servers {
		go func(ctx context.Context, server string) {
			_url, err := url.Parse("https://" + server + path)
			if err != nil {
				fails <- err
				return
			}

			c.logIfEnabled(_url.String())

			res, err := c.Get(_url.String())
			if err != nil {
				fails <- err
				return
			}
			defer c.logIfError(res.Body.Close)

			if res.StatusCode != http.StatusOK {
				b, err := io.ReadAll(res.Body)
				if err != nil {
					fails <- err
					return
				}

				fails <- fmt.Errorf("%d: %s: %s: %s\n", res.StatusCode, res.Status, _url.String(), string(b))
				return
			}

			r := new(Response)
			if err := json.NewDecoder(res.Body).Decode(r); err != nil {
				fails <- err
				return
			}

			result <- r

			c.logIfEnabled(fmt.Sprintf("got result from %s\n", server))

			cancel()
		}(ctx, server)
	}

	var numErrors int

	for {
		select {
		case r := <-result:
			cancel()
			return r, nil
		case err := <-fails:
			numErrors++
			c.logIfEnabled(fmt.Sprintf("errnum (%d/%d): %s", numErrors, len(c.servers), err))
			if numErrors == len(c.servers) {
				cancel()
				return nil, fmt.Errorf("all requests failed")
			}
		}
	}
}

func (c *CCP) logIfEnabled(s string) {
	if c.logging {
		log.Println(s)
	}
}

func (c *CCP) logIfError(f func() error) {
	if err := f(); err != nil {
		c.logIfEnabled(err.Error())
	}
}
