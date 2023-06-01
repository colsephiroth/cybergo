package pvwa

import (
	"net/url"
)

// GetSafes This method returns a list of all Safes in the Vault that the user has permissions for.
// The user who runs this web service must be a member of the Safes in the Vault that are returned in the list.
func (p *PVWA) GetSafes() *GetSafesOptions {
	return &GetSafesOptions{
		path:  "API/Safes",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (s *GetSafesOptions) Run() ([]*Safe, error) {
	return genericGetReturnSlice[Safe](s.pvwa, s.path, s.query)
}

// GetSafeMembers This method returns the list of members of a Safe. The user who run this web
// service must have View Safe Members permissions on the Safe.
func (p *PVWA) GetSafeMembers(safeUrlId string) *GetSafeMembersOptions {
	return &GetSafeMembersOptions{
		path:  "API/Safes" + safeUrlId + "/Members",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (s *GetSafeMembersOptions) Run() ([]*SafeMember, error) {
	return genericGetReturnSlice[SafeMember](s.pvwa, s.path, s.query)
}

// UpdateSafeMembers This method adds an existing user as a Safe member. The user who runs
// this web service must have Manage Safe Members permissions in the Vault.
func (p *PVWA) UpdateSafeMembers(safeUrlId string, user *SafeMember) *UpdateSafeMembersOptions {
	return &UpdateSafeMembersOptions{
		path:  "API/Safes/" + safeUrlId + "/members",
		query: &url.Values{},
		pvwa:  p,
		user:  user,
	}
}

func (s *UpdateSafeMembersOptions) Run() (*SafeMember, error) {
	return genericUpdateReturnSingle[*SafeMember, SafeMember](s.pvwa, s.path, s.query, s.user)
}
