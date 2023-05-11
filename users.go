package cybergo

import (
	"encoding/json"
)

func (p *PVWA) GetUsers(options ...ApiOption) ([]*UserDetails, error) {
	var data GetUsersResponse

	path, err := p.buildPath("Users", options...)
	if err != nil {
		return nil, err
	}

	p.logIfEnabled(path)

	res, err := p.Get(path)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res).Decode(&data); err != nil {
		return nil, err
	}

	p.logIfError(res.Close)

	return data.Users, nil
}
