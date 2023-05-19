package pvwa

import (
	"encoding/json"
)

func (p *PVWA) GetAccounts(options ...ApiOption) ([]*AccountDetails, error) {
	path, err := p.buildPath("Accounts", options...)
	if err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}

	var accounts []*AccountDetails

	for {
		p.logIfEnabled(path)

		data := new(GenericResponse[*AccountDetails])

		res, err := p.Get(path)
		if err != nil {
			p.logIfEnabled(err.Error())
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			p.logIfEnabled(err.Error())
			return nil, err
		}

		p.logIfError(res.Close)

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
	path, err := p.buildPath("Accounts/" + id)
	if err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}

	account := new(AccountDetails)

	res, err := p.Get(path)
	if err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}
	defer p.logIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&account); err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}

	return account, nil
}

func (p *PVWA) UpdateAccount(id string, ops []UpdateAccountOperation) (*AccountDetails, error) {
	path, err := p.buildPath("Accounts/" + id)
	if err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}

	account := new(AccountDetails)

	data, err := json.Marshal(ops)
	if err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}

	res, err := p.Patch(path, data)
	if err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}
	defer p.logIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&account); err != nil {
		p.logIfEnabled(err.Error())
		return nil, err
	}

	return account, nil
}
