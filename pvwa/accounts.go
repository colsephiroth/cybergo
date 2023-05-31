package pvwa

import (
	"encoding/json"
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
	path := buildPath(a.path, a.query)

	var accounts []*AccountDetails

	for {
		a.pvwa.logIfEnabled(path)

		data := new(GenericResponse[*AccountDetails])

		res, err := a.pvwa.Get(path)
		if err != nil {
			a.pvwa.logIfEnabled(err.Error())
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			a.pvwa.logIfEnabled(err.Error())
			return nil, err
		}

		a.pvwa.logIfError(res.Close)

		accounts = append(accounts, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return accounts, nil
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
	path := buildPath(a.path, a.query)

	account := new(AccountDetails)

	res, err := a.pvwa.Get(path)
	if err != nil {
		a.pvwa.logIfEnabled(err.Error())
		return nil, err
	}
	defer a.pvwa.logIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&account); err != nil {
		a.pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	return account, nil
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
	path := buildPath(a.path, a.query)

	account := new(AccountDetails)

	data, err := json.Marshal(a.operations)
	if err != nil {
		a.pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	res, err := a.pvwa.Patch(path, data)
	if err != nil {
		a.pvwa.logIfEnabled(err.Error())
		return nil, err
	}
	defer a.pvwa.logIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&account); err != nil {
		a.pvwa.logIfEnabled(err.Error())
		return nil, err
	}

	return account, nil
}
