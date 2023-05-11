package cybergo

import (
	"encoding/json"
	"io"
	"log"
)

func (p *PVWA) GetSafes(options ...ApiOption) ([]*Safe, error) {
	var data []*Safe

	path, err := buildPath("Safes", options...)
	if err != nil {
		return nil, err
	}

	log.Println(path)

	res, err := p.Get(path)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res)
	if err != nil {
		return nil, err
	}

	log.Println(string(b))

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	LogIfError(res.Close)

	return data, nil
}
