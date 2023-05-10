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
	Base          string
	Username      string
	Authorization string
	*http.Client
}

func NewPVWA(base, username string) (*PVWA, error) {
	pvwa := &PVWA{
		Base:     base,
		Username: username,
		Client:   http.DefaultClient,
	}

	if err := pvwa.Logon(); err != nil {
		return nil, err
	}

	return pvwa, nil
}

func (p *PVWA) Logon() error {
	fmt.Printf("Password for %s: ", p.Username)

	passwordBytes, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return err
	}

	fmt.Println()

	data, err := json.Marshal(map[string]interface{}{
		"username":          p.Username,
		"password":          string(passwordBytes),
		"concurrentSession": true,
	})
	if err != nil {
		return err
	}

	res, err := p.Client.Post(p.Base+"API/auth/Cyberark/Logon", "application/json", bytes.NewBuffer(data))
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

	p.Authorization = strings.ReplaceAll(string(body), "\"", "")

	return nil
}

func (p *PVWA) Logoff() error {
	req, err := http.NewRequest("POST", p.Base+"API/auth/Logoff", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", p.Authorization)

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
	req, err := http.NewRequest(http.MethodGet, p.Base+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", p.Authorization)

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
	req, err := http.NewRequest(http.MethodPost, p.Base+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", p.Authorization)
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
	req, err := http.NewRequest(http.MethodPatch, p.Base+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", p.Authorization)
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
