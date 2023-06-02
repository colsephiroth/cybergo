package pvwa

import (
	"net/url"
	"strconv"
	"strings"
)

func buildPath(path string, query *url.Values) string {
	_path, _ := url.Parse(path)

	_path.RawQuery = query.Encode()

	return _path.String()
}

func withExtendedDetails(b bool) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("ExtendedDetails", strconv.FormatBool(b))
	}
}

func withSearch(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("search", s)
	}
}

func withSort(s []string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("sort", strings.Join(s, ","))
	}
}

func withUserName(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("userName", s)
	}
}

func withUserType(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("userType", s)
	}
}

func withComponentUser(b bool) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("componentUser", strconv.FormatBool(b))
	}
}

func withSavedFilter(i int) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("savedFilter", strconv.Itoa(i))
	}
}

func withOffset(i int) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("offset", strconv.Itoa(i))
	}
}

func withLimit(i int) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("limit", strconv.Itoa(i))
	}
}

func withFilter(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("filter", s)
	}
}

func withSearchType(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("searchType", s)
	}
}

func withIncludeMembers(b bool) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("includeMembers", strconv.FormatBool(b))
	}
}

func withUseCache(b bool) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("useCache", strconv.FormatBool(b))
	}
}

func withIncludeAccounts(b bool) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("includeAccounts", strconv.FormatBool(b))
	}
}

func withSafe(s string) func(v *url.Values) {
	return func(v *url.Values) {
		v.Add("safe", s)
	}
}
