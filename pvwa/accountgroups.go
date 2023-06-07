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

// NewAccountGroup This method enables application managers to define a new account group automatically,
// and manage accounts as part of a group. To create an account group, users require the following permissions
// in the Safe where the group is created: Add accounts, Update account content, Update account properties, Create folders.
func (p *PVWA) NewAccountGroup(group *AccountGroup) *NewAccountGroupOptions {
	return &NewAccountGroupOptions{
		path:  "API/AccountGroups",
		query: &url.Values{},
		pvwa:  p,
		data:  group,
	}
}

func (a *NewAccountGroupOptions) Run() (*AccountGroup, error) {
	return postReturnSingle[*AccountGroup, AccountGroup](a.pvwa, a.path, a.query, a.data)
}

// GetAccountGroupMembers This method returns all the members of an existing account group.
// These accounts can be either password accounts or SSH Key accounts. The user performing
// this task must have the following permissions in the Safe: List Accounts.
func (p *PVWA) GetAccountGroupMembers(id string) *GetAccountGroupMembersOptions {
	return &GetAccountGroupMembersOptions{
		path:  "API/AccountGroups/" + id + "/members",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (a *GetAccountGroupMembersOptions) Run() ([]*AccountGroupMember, error) {
	return getSingleRequestReturnSlice[AccountGroupMember](a.pvwa, a.path, a.query)
}

// NewAccountGroupMember This method adds an account as a member to an existing account group. The account can
// contain either a password or an SSH key. All members of an account group must be stored in the same Safe as
// the group itself. To add an account as a member to an account group, users require the following permissions
// in the Safe where the group is created: Add accounts, Update account content, Update account properties.
func (p *PVWA) NewAccountGroupMember(id string, member *AccountGroupMember) *NewAccountGroupMemberOptions {
	return &NewAccountGroupMemberOptions{
		path:  "API/AccountGroups/" + id + "/members",
		query: &url.Values{},
		pvwa:  p,
		data:  member,
	}
}

func (a *NewAccountGroupMemberOptions) Run() (*AccountGroupMember, error) {
	return postReturnSingle[*AccountGroupMember, AccountGroupMember](a.pvwa, a.path, a.query, a.data)
}

// DeleteAccountGroupMember This method removes an account member from an account group. This account can be
// either a password account or an SSH Key account. The user performing this task must have the following
// permissions in the Safe: Add accounts, Update account content, Update account properties, Create folders.
func (p *PVWA) DeleteAccountGroupMember(groupID, accountID string) *DeleteAccountGroupMemberOptions {
	return &DeleteAccountGroupMemberOptions{
		path:  "API/AccountGroups/" + groupID + "/members/" + accountID,
		query: &url.Values{},
		pvwa:  p,
	}
}

func (a *DeleteAccountGroupMemberOptions) Run() error {
	return deleteReturnNone(a.pvwa, a.path, a.query)
}
