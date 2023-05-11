package cybergo

import (
	"encoding/json"
	"log"
)

func (p *PVWA) GetUserGroups(options ...ApiOption) ([]*UserGroup, error) {
	path, err := buildPath("UserGroups", options...)
	if err != nil {
		return nil, err
	}

	var groups []*UserGroup

	for {
		log.Println(path)

		data := new(GetUserGroupsResponse)

		res, err := p.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		LogIfError(res.Close)

		groups = append(groups, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return groups, nil
}
