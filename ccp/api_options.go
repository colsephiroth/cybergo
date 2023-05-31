package ccp

import (
	"net/url"
	"strconv"
)

func buildPath(path string, query *url.Values) string {
	_path, _ := url.Parse(path)

	values := _path.Query()

	for k, v := range *query {
		values.Set(k, v[0])
	}

	_path.RawQuery = values.Encode()

	return _path.String()
}

func withSafe(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("Safe", s)
	}
}

func withFolder(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("Folder", s)
	}
}

func withObject(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("Object", s)
	}
}

func withUserName(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("UserName", s)
	}
}

func withAddress(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("Address", s)
	}
}

func withDatabase(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("Database", s)
	}
}

func withPolicyID(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("PolicyID", s)
	}
}

func withReason(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("Reason", s)
	}
}

func withConnectionTimeout(i int) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("ConnectionTimeout", strconv.Itoa(i))
	}
}

func withQuery(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("Query", s)
	}
}

func withQueryFormat(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("QueryFormat", s)
	}
}

func withFailRequestOnPasswordChange(b bool) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("FailRequestOnPasswordChange", strconv.FormatBool(b))
	}
}
