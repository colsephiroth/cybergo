package pvwa

import (
	"encoding/json"
)

func (p *PVWA) GetSafes(options ...ApiOption) ([]*Safe, error) {
	path, err := p.buildPath("Safes", options...)
	if err != nil {
		return nil, err
	}

	var safes []*Safe

	for {
		p.logIfEnabled(path)

		data := new(GenericResponse[*Safe])

		res, err := p.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		p.logIfError(res.Close)

		safes = append(safes, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return safes, nil
}

func (p *PVWA) GetSafeMembers(safeUrlId string, options ...ApiOption) ([]*SafeMember, error) {
	path, err := p.buildPath("Safes/"+safeUrlId+"/Members", options...)
	if err != nil {
		return nil, err
	}

	var members []*SafeMember

	for {
		p.logIfEnabled(path)

		data := new(GenericResponse[*SafeMember])

		res, err := p.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		p.logIfError(res.Close)

		members = append(members, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return members, nil
}
