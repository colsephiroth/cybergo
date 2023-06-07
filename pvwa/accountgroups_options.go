package pvwa

import "net/url"

type GetAccountGroupsOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type NewAccountGroupOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
	data  *AccountGroup
}

type GetAccountGroupMembersOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type NewAccountGroupMemberOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
	data  *AccountGroupMember
}

type DeleteAccountGroupMemberOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}
