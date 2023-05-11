package cybergo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type PVWA struct {
	base          string
	username      string
	password      string
	authorization string
	*http.Client
}

type PVWAOption func(p *PVWA)

func WithPassword(s string) PVWAOption {
	return func(p *PVWA) {
		p.password = s
	}
}

func NewPVWA(subdomain, username string, options ...PVWAOption) (*PVWA, error) {
	pvwa := &PVWA{
		base:     fmt.Sprintf("https://%s.privilegecloud.cyberark.com/PasswordVault/", subdomain),
		username: username,
		Client:   http.DefaultClient,
	}

	for _, option := range options {
		option(pvwa)
	}

	if err := pvwa.Logon(); err != nil {
		return nil, err
	}

	return pvwa, nil
}

func (p *PVWA) Logon() error {
	password := p.password

	if p.password == "" {
		fmt.Printf("Password for %s: ", p.username)

		passwordBytes, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			return err
		}
		fmt.Println()

		password = string(passwordBytes)
	}

	data, err := json.Marshal(map[string]any{
		"username":          p.username,
		"password":          password,
		"concurrentSession": true,
	})
	if err != nil {
		return err
	}

	res, err := p.Client.Post(p.base+"API/auth/Cyberark/Logon", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("%d: %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	p.authorization = strings.ReplaceAll(string(body), "\"", "")

	return nil
}

func (p *PVWA) Logoff() error {
	req, err := http.NewRequest("POST", p.base+"API/auth/Logoff", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", p.authorization)

	res, err := p.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("%d: %s", res.StatusCode, res.Status)
	}

	return nil
}

func (p *PVWA) Get(path string) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, p.base+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", p.authorization)

	res, err := p.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}

func (p *PVWA) Post(path string, data []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodPost, p.base+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", p.authorization)
	req.Header.Set("Content-Type", "application/json")

	res, err := p.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}

func (p *PVWA) Patch(path string, data []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodPatch, p.base+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", p.authorization)
	req.Header.Set("Content-Type", "application/json")

	res, err := p.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}
