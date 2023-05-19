package ccp

import (
	"net/url"
	"strconv"
)

type ApiOption func(v *url.Values)

func (c *CCP) buildPath(path string, options ...ApiOption) (string, error) {
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

func WithSafe(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("Safe", s)
	}
}

func WithFolder(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("Folder", s)
	}
}

func WithObject(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("Object", s)
	}
}

func WithUserName(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("UserName", s)
	}
}

func WithAddress(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("Address", s)
	}
}

func WithDatabase(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("Database", s)
	}
}

func WithPolicyID(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("PolicyID", s)
	}
}

func WithReason(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("Reason", s)
	}
}

func WithConnectionTimeout(i int) ApiOption {
	return func(v *url.Values) {
		v.Add("ConnectionTimeout", strconv.Itoa(i))
	}
}

func WithQuery(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("Query", s)
	}
}

func WithQueryFormat(s string) ApiOption {
	return func(v *url.Values) {
		v.Add("QueryFormat", s)
	}
}

func WithFailRequestOnPasswordChange(b bool) ApiOption {
	return func(v *url.Values) {
		v.Add("FailRequestOnPasswordChange", strconv.FormatBool(b))
	}
}
