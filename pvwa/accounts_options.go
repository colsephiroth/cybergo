package pvwa

import "net/url"

type GetAccountsOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type NewAccountOptions struct {
	path    string
	query   *url.Values
	pvwa    *PVWA
	account *AccountDetails
}

type UpdateAccountOptions struct {
	path       string
	query      *url.Values
	pvwa       *PVWA
	operations []UpdateAccountOperation
}

type ChangePasswordOptions struct {
	path             string
	query            *url.Values
	pvwa             *PVWA
	changeProperties *ChangeCredentialsNow
}

type CheckInOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type GrantAdministrativeAccessOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type DeleteLinkedAccountOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type RetrievePasswordOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type GetAccountDetailsOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

// WithSavedFilter Search for accounts using a saved filter(s).
func (a *GetAccountsOptions) WithSavedFilter(i int) *GetAccountsOptions {
	withSavedFilter(i)(a.query)
	return a
}

// WithSearch List of keywords separated with space to search in accounts.
func (a *GetAccountsOptions) WithSearch(s string) *GetAccountsOptions {
	withSearch(s)(a.query)
	return a
}

// WithSort Property or properties to sort returned accounts, followed by asc (default) or desc to control sort direction.
// Multiple sorts are comma-separated. Maximum number of properties is 3.
func (a *GetAccountsOptions) WithSort(s []string) *GetAccountsOptions {
	withSort(s)(a.query)
	return a
}

// WithOffset Offset of the first returned account into the collection of results.
func (a *GetAccountsOptions) WithOffset(i int) *GetAccountsOptions {
	withOffset(i)(a.query)
	return a
}

// WithLimit Maximum number of returned accounts. If not specified, the default value is 50. The maximum number
// that can be specified is 1000.
func (a *GetAccountsOptions) WithLimit(i int) *GetAccountsOptions {
	withLimit(i)(a.query)
	return a
}

// WithFilter Search for accounts filtered by specific Safe. Filter=safename eq [safe name]
func (a *GetAccountsOptions) WithFilter(s string) *GetAccountsOptions {
	withFilter(s)(a.query)
	return a
}

// WithSearchType Type of search to perform - if keywords are contained in the relevant account properties values or in
// the start of the properties values (the latter enhances performance)
func (a *GetAccountsOptions) WithSearchType(s string) *GetAccountsOptions {
	withSearchType(s)(a.query)
	return a
}
