package pvwa

import (
	"net/url"
	"strconv"
	"strings"
)

type ApiOption func(v *url.Values)

func (p *PVWA) buildPath(path string, options ...ApiOption) (string, error) {
	_path, err := url.Parse("API/" + path)
	if err != nil {
		return "", err
	}

	v := _path.Query()

	for _, option := range options {
		option(&v)
	}

	_path.RawQuery = v.Encode()

	return _path.String(), nil
}

func WithExtendedDetails(b bool) ApiOption {
	return func(v *url.Values) {
		v.Add("ExtendedDetails", strconv.FormatBool(b))
	}
}

func WithSearch(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("search", s)
	}
}

func WithSort(s []string) ApiOption {
	return func(v *url.Values) {
		v.Add("sort", strings.Join(s, ","))
	}
}

func WithUserName(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("userName", s)
	}
}

func WithUserType(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("userType", s)
	}
}

func WithComponentUser(b bool) ApiOption {
	return func(v *url.Values) {
		v.Add("componentUser", strconv.FormatBool(b))
	}
}

func WithSavedFilter(i int) ApiOption {
	return func(v *url.Values) {
		v.Add("savedFilter", strconv.Itoa(i))
	}
}

func WithOffset(i int) ApiOption {
	return func(v *url.Values) {
		v.Add("offset", strconv.Itoa(i))
	}
}

func WithLimit(i int) ApiOption {
	return func(v *url.Values) {
		v.Add("limit", strconv.Itoa(i))
	}
}

func WithFilter(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("filter", s)
	}
}

func WithSearchType(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("searchType", s)
	}
}

func WithIncludeMembers(b bool) ApiOption {
	return func(v *url.Values) {
		v.Add("includeMembers", strconv.FormatBool(b))
	}
}

func WithUseCache(b bool) ApiOption {
	return func(v *url.Values) {
		v.Add("useCache", strconv.FormatBool(b))
	}
}

func WithIncludeAccounts(b bool) ApiOption {
	return func(v *url.Values) {
		v.Add("includeAccounts", strconv.FormatBool(b))
	}
}
