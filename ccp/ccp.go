package ccp

import (
	"fmt"
	"log"
	"net/http"
)

type CCP struct {
	servers []string
	path    string
	logging bool

	*http.Client
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
