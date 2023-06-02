package pvwa

import "net/url"

type GetAccountGroupsOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}
