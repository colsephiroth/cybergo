package cybergo

import (
	"encoding/json"
	"log"
)

func (p *PVWA) GetUsers(options ...ApiOption) ([]*UserDetails, error) {
	var data GetUsersResponse

	path, err := buildPath("Users", options...)
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

	return data.Users, nil
}
