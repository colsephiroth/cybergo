package cybergo

import (
	"encoding/json"
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

	if err := json.NewDecoder(res).Decode(&data); err != nil {
		return nil, err
	}

	LogIfError(res.Close)

	return data, nil
}
