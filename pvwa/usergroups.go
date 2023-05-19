package pvwa

import (
	"encoding/json"
)

func (p *PVWA) GetUserGroups(options ...ApiOption) ([]*UserGroup, error) {
	path, err := p.buildPath("UserGroups", options...)
	if err != nil {
		return nil, err
	}

	var groups []*UserGroup

	for {
		p.logIfEnabled(path)

		data := new(GenericResponse[*UserGroup])

		res, err := p.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		p.logIfError(res.Close)

		groups = append(groups, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return groups, nil
}
