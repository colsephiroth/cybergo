package cybergo

import (
	"encoding/json"
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
