package cybergo

import (
	"encoding/json"
	"log"
)

func (p *PVWA) GetAccounts(options ...ApiOption) ([]*AccountDetails, error) {
	path, err := buildPath("Accounts", options...)
	if err != nil {
		return nil, err
	}

	var accounts []*AccountDetails

	for {
		log.Println(path)

		data := new(GetAccountsResponse)

		res, err := p.Get(path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		LogIfError(res.Close)

		accounts = append(accounts, data.Value...)

		if data.NextLink != "" {
			path = data.NextLink
		} else {
			break
		}
	}

	return accounts, nil
}

func (p *PVWA) GetAccountDetails(id string) (*AccountDetails, error) {
	path, err := buildPath("Accounts/" + id)
	if err != nil {
		return nil, err
	}

	account := new(AccountDetails)

	res, err := p.Get(path)
	if err != nil {
		return nil, err
	}
	defer LogIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&account); err != nil {
		return nil, err
	}

	return account, nil
}

func (p *PVWA) UpdateAccount(id string, ops []UpdateAccountOperation) (*AccountDetails, error) {
	path, err := buildPath("Accounts/" + id)
	if err != nil {
		return nil, err
	}

	account := new(AccountDetails)

	data, err := json.Marshal(ops)
	if err != nil {
		return nil, err
	}

	res, err := p.Patch(path, data)
	if err != nil {
		return nil, err
	}
	defer LogIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&account); err != nil {
		return nil, err
	}

	return account, nil
}
