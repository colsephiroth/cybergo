package cybergo

import (
	"encoding/json"
	"fmt"
	"io"
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

	b, err := io.ReadAll(res)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	if err := json.Unmarshal(b, &users); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res).Decode(&users); err != nil {
		return nil, err
	}

	LogIfError(res.Close)

	return users, nil
}
