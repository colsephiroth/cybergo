package pvwa

import (
	"net/url"
)

// GetAccounts This method returns a list of all the accounts in the Vault. The user who runs this web
// service requires List Accounts permission in the Safe.
func (p *PVWA) GetAccounts() *GetAccountsOptions {
	return &GetAccountsOptions{
		path:  "API/Accounts",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (a *GetAccountsOptions) Run() ([]*AccountDetails, error) {
	return getMultipleRequestReturnSlice[AccountDetails](a.pvwa, a.path, a.query)
}

// NewAccount This method adds a new privileged account or SSH key to the Vault The user who runs this web service
// requires the following permission in the Vault
//   - Add account AND ( update password OR update password properties )
func (p *PVWA) NewAccount(account *AccountDetails) *NewAccountOptions {
	return &NewAccountOptions{
		path:    "API/Accounts",
		query:   &url.Values{},
		pvwa:    p,
		account: account,
	}
}

func (a *NewAccountOptions) Run() (*AccountDetails, error) {
	return postReturnSingle[*AccountDetails, AccountDetails](a.pvwa, a.path, a.query, a.account)
}

// UpdateAccount This method updates an existing account's details. It isn't mandatory to send all
// the account’s details. Any values sent in the request that were changed will be updated. All other
// properties values will remain the same.
//
// On each property, the following are the allowed operations:
//   - Replace (to replace the existing value of that property)
//   - Remove (to remove the property from the account)
//   - Add (to add that property to the account)
//
// It is possible to set several properties using the same command using the following structure:
//
//	{
//		"op": "replace",
//		"path": "/platformaccountproperties",
//		"value": "{\"{PropertyID1}\":\"{Value}\",\"{PropertyID2}\":\"{Value}\",\"{PropertyID3}\":\"{Value}\"}"
//	}
//
// When sending several operations on the same property – only the last operation will affect.
func (p *PVWA) UpdateAccount(id string, ops []UpdateAccountOperation) *UpdateAccountOptions {
	return &UpdateAccountOptions{
		path:       "API/Accounts/" + id,
		query:      &url.Values{},
		pvwa:       p,
		operations: ops,
	}
}

func (a *UpdateAccountOptions) Run() (*AccountDetails, error) {
	return patchReturnSingle[[]UpdateAccountOperation, AccountDetails](a.pvwa, a.path, a.query, a.operations)
}

// ChangePassword This method marks the account for an immediate credentials change by the CPM to a
// new random password. The user who runs this web service requires the following permission in the Safe
// where the privileged account is stored: 'Initiate CPM password management operations'
func (p *PVWA) ChangePassword(id string, changeEntireGroup bool) *ChangePasswordOptions {
	return &ChangePasswordOptions{
		path:             "API/Accounts/" + id + "/Change",
		query:            &url.Values{},
		pvwa:             p,
		changeProperties: &ChangeCredentialsNow{ChangeEntireGroup: changeEntireGroup},
	}
}

func (a *ChangePasswordOptions) Run() error {
	return postReturnNone[*ChangeCredentialsNow](a.pvwa, a.path, a.query, a.changeProperties)
}

// CheckIn This method checks an exclusive account into the Vault. If the account is managed automatically
// by the CPM, after it is checked in, the password is changed immediately. If the account is managed manually,
// a notification is sent to a user who is authorized to change the password. The account is checked in automatically
// after it has been changed. The user who runs this web service requires the following permission in the Safe where
// the privileged account is stored: 'Initiate CPM password management operations'
func (p *PVWA) CheckIn(id string) *CheckInOptions {
	return &CheckInOptions{
		path:  "API/Accounts/" + id + "/CheckIn",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (a *CheckInOptions) Run() error {
	return postReturnNone(a.pvwa, a.path, a.query, nil)
}

// GrantAdministrativeAccess This method requests and receives access to a target Windows machine with
// administrative rights. The domain user who runs this web service will be added to the local Administrators
// group of the target machine. Make sure that all users who want to request access to the target Windows
// machine must have the following permission in the Safe: List accounts, Use accounts. Make sure that the
// platform is enabled for ad hoc access at platform level. For more information, see the product documentation.
// End user machine environment that will be able to get Ad-Hoc access:
//   - Windows Server 2012/2012R2/2016
//   - Windows 8, Windows 10
func (p *PVWA) GrantAdministrativeAccess(id string) *GrantAdministrativeAccessOptions {
	return &GrantAdministrativeAccessOptions{
		path:  "API/Accounts/" + id + "/GrantAdministrativeAccess",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (a *GrantAdministrativeAccessOptions) Run() error {
	return postReturnNone(a.pvwa, a.path, a.query, nil)
}

// DeleteLinkedAccount This method enables a user to remove association between linked account and source account.
// To run this service, the user must have the following Safe member authorizations for the Safe of the source
// account: List accounts Update account properties Manage Safe - This authorization is needed only in case
// "RequireManageSafeToClearLinkedAccount" is enabled in the configuration.
func (p *PVWA) DeleteLinkedAccount(id string, index string) *DeleteLinkedAccountOptions {
	return &DeleteLinkedAccountOptions{
		path:  "API/Accounts/" + id + "/LinkAccount/" + index,
		query: &url.Values{},
		pvwa:  p,
	}
}

func (a *DeleteLinkedAccountOptions) Run() error {
	return deleteReturnNone(a.pvwa, a.path, a.query)
}

// GetAccountDetails This method returns information about an account identified by its id. The user who
// runs this web service requires List Accounts permission in the Vault.
func (p *PVWA) GetAccountDetails(id string) *GetAccountDetailsOptions {
	return &GetAccountDetailsOptions{
		path:  "API/Accounts/" + id,
		query: &url.Values{},
		pvwa:  p,
	}
}

func (a *GetAccountDetailsOptions) Run() (*AccountDetails, error) {
	return getReturnSingle[AccountDetails](a.pvwa, a.path, a.query)
}
