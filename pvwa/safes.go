package pvwa

import (
	"encoding/json"
	"net/url"
)

// GetSafes This method returns a list of all Safes in the Vault that the user has permissions for.
// The user who runs this web service must be a member of the Safes in the Vault that are returned in the list.
func (p *PVWA) GetSafes() *GetSafesOptions {
	return &GetSafesOptions{
		path:  "API/Safes",
		query: new(url.Values),
		pvwa:  p,
	}
}

func (s *GetSafesOptions) Run() ([]*Safe, error) {
	path := buildPath(s.path, s.query)

	var safes []*Safe

	for {
		s.pvwa.logIfEnabled(path)

		data := new(GenericResponse[*Safe])

		res, err := s.pvwa.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		s.pvwa.logIfError(res.Close)

		safes = append(safes, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return safes, nil
}

// GetSafeMembers This method returns the list of members of a Safe. The user who run this web
// service must have View Safe Members permissions on the Safe.
func (p *PVWA) GetSafeMembers(safeUrlId string) *GetSafeMembersOptions {
	return &GetSafeMembersOptions{
		path:  "API/Safes" + safeUrlId + "/Members",
		query: new(url.Values),
		pvwa:  p,
	}
}

func (s *GetSafeMembersOptions) Run() ([]*SafeMember, error) {
	path := buildPath(s.path, s.query)

	var members []*SafeMember

	for {
		s.pvwa.logIfEnabled(path)

		data := new(GenericResponse[*SafeMember])

		res, err := s.pvwa.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		s.pvwa.logIfError(res.Close)

		members = append(members, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return members, nil
}
