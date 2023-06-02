package pvwa

import "net/url"

// GetAccountGroups This method returns all the account groups in a specific Safe. The user performing
// this task must have the following permissions in the Safe: List Accounts.
func (p *PVWA) GetAccountGroups(safe string) *GetAccountGroupsOptions {
	return &GetAccountGroupsOptions{
		path:  "API/AccountGroups",
		query: &url.Values{"safe": []string{safe}},
		pvwa:  p,
	}
}

func (a *GetAccountGroupsOptions) Run() ([]*AccountGroup, error) {
	return getSingleRequestReturnSlice[AccountGroup](a.pvwa, a.path, a.query)
}
