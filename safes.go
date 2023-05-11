package cybergo

import (
	"encoding/json"
	"io"
	"log"
)

func (p *PVWA) GetSafes(options ...ApiOption) ([]*Safe, error) {
	path, err := buildPath("Safes", options...)
	if err != nil {
		return nil, err
	}

	var safes []*Safe

	for {
		log.Println(path)

		data := new(GenericResponse[*Safe])

		res, err := p.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		LogIfError(res.Close)

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
	var data []*SafeMember

	path, err := buildPath("Safes/"+safeUrlId+"/Members", options...)
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

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	LogIfError(res.Close)

	return data, nil
}
