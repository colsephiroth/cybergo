package cybergo

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func (p *PVWA) GetUsers(options *GetUsersOptions) ([]*UserDetails, error) {
	path, err := url.Parse("API/Users")
	if err != nil {
		return nil, err
	}

	var users []*UserDetails

	if options != nil {
		if options.ExtendedDetails != nil {
			path.Query().Add("ExtendedDetails", strconv.FormatBool(*options.ExtendedDetails))
		}
		if options.Search != nil {
			path.Query().Add("search", *options.Search)
		}
		if options.Sort != nil {
			path.Query().Add("sort", strings.Join(*options.Sort, ","))
		}
		if options.UserName != nil {
			path.Query().Add("userName", *options.UserName)
		}
		if options.UserType != nil {
			path.Query().Add("userType", *options.UserType)
		}
		if options.ComponentUser != nil {
			path.Query().Add("componentUser", strconv.FormatBool(*options.ComponentUser))
		}

		path.RawQuery = path.Query().Encode()
	}

	_path := path.String()

	log.Println(_path)

	res, err := p.Get(_path)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res).Decode(&users); err != nil {
		return nil, err
	}

	LogIfError(res.Close)

	return users, nil
}
