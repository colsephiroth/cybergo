package cybergo

import (
	"encoding/json"
	"log"
)

func (p *PVWA) GetUsers(options ...GetUsersOption) ([]*UserDetails, error) {
	var users []*UserDetails

	path := buildGetUsersPath(options...)

	log.Println(path)

	res, err := p.Get(path.String())
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res).Decode(&users); err != nil {
		return nil, err
	}

	LogIfError(res.Close)

	return users, nil
}
