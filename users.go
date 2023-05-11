package cybergo

import (
	"encoding/json"
	"log"
)

func (p *PVWA) GetUsers(options ...GetUsersOption) ([]*UserDetails, error) {
	var data GetUsersResponse

	path := buildGetUsersPath(options...)

	log.Println(path.String())

	res, err := p.Get(path.String())
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res).Decode(&data); err != nil {
		return nil, err
	}

	LogIfError(res.Close)

	return data.Users, nil
}
