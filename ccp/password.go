package ccp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetPassword This service enables applications to retrieve passwords from the Central Credential Provider.
func (c *CCP) GetPassword() *GetPasswordOptions {
	return &GetPasswordOptions{
		path:  c.path,
		query: &url.Values{},
		ccp:   c,
	}
}

func (p *GetPasswordOptions) Run() (*Response, error) {
	ctx, cancel := context.WithCancel(context.Background())

	result := make(chan *Response)
	fails := make(chan error)

	path := buildPath(p.path, p.query)

	for _, server := range p.ccp.servers {
		go func(ctx context.Context, server string) {
			_url, err := url.Parse("https://" + server + path)
			if err != nil {
				fails <- err
				return
			}

			p.ccp.logIfEnabled(_url.String())

			res, err := p.ccp.Get(_url.String())
			if err != nil {
				fails <- err
				return
			}
			defer p.ccp.logIfError(res.Body.Close)

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

			p.ccp.logIfEnabled(fmt.Sprintf("got result from %s\n", server))

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
			p.ccp.logIfEnabled(fmt.Sprintf("errnum (%d/%d): %s", numErrors, len(p.ccp.servers), err))
			if numErrors == len(p.ccp.servers) {
				cancel()
				return nil, fmt.Errorf("all requests failed")
			}
		}
	}
}
