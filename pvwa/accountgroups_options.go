package pvwa

import "net/url"

type GetAccountGroupsOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

func (a *GetAccountGroupsOptions) WithSafe(s string) *GetAccountGroupsOptions {
	withSafe(s)(a.query)
	return a
}
