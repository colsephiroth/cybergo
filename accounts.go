package cybergo

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func (p *PVWA) GetAccounts(options *GetAccountsOptions) ([]*AccountDetails, error) {
	path, err := url.Parse("API/Accounts")
	if err != nil {
		return nil, err
	}

	var accounts []*AccountDetails

	if options != nil {
		if options.SavedFilter != nil {
			path.Query().Add("savedFilter", strconv.Itoa(*options.SavedFilter))
		}
		if options.Search != nil {
			path.Query().Add("search", *options.Search)
		}
		if options.Sort != nil {
			path.Query().Add("sort", strings.Join(*options.Sort, ","))
		}
		if options.Offset != nil {
			path.Query().Add("offset", strconv.Itoa(*options.Offset))
		}
		if options.Limit != nil {
			path.Query().Add("limit", strconv.Itoa(*options.Limit))
		}
		if options.Filter != nil {
			path.Query().Add("filter", *options.Filter)
		}
		if options.SearchType != nil {
			path.Query().Add("searchType", *options.SearchType)
		}

		path.RawQuery = path.Query().Encode()
	}

	_path := path.String()

	for {
		log.Println(_path)

		data := new(GetAccountsResponse)

		res, err := p.Get(_path)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res).Decode(&data); err != nil {
			return nil, err
		}

		LogIfError(res.Close)

		accounts = append(accounts, data.Value...)

		if data.NextLink != "" {
			_path = data.NextLink
		} else {
			break
		}
	}

	return accounts, nil
}

func (p *PVWA) GetAccountDetails(id string) (*AccountDetails, error) {
	path, err := url.Parse("API/Accounts/" + id)
	if err != nil {
		return nil, err
	}

	account := new(AccountDetails)

	res, err := p.Get(path.String())
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
	path, err := url.Parse("API/Accounts/" + id + "/")
	if err != nil {
		return nil, err
	}

	account := new(AccountDetails)

	data, err := json.Marshal(ops)
	if err != nil {
		return nil, err
	}

	res, err := p.Patch(path.String(), data)
	if err != nil {
		return nil, err
	}
	defer LogIfError(res.Close)

	if err := json.NewDecoder(res).Decode(&account); err != nil {
		return nil, err
	}

	return account, nil
}
