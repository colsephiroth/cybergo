package cybergo

import (
	"net/url"
	"strconv"
	"strings"
)

type ApiOption func(v *url.Values)

func buildPath(path string, options ...ApiOption) (string, error) {
	p, err := url.Parse("API/" + path)
	if err != nil {
		return "", err
	}

	v := p.Query()

	for _, option := range options {
		option(&v)
	}

	p.RawQuery = v.Encode()

	return p.String(), nil
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
