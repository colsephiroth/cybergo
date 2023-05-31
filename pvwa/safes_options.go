package pvwa

import "net/url"

type GetSafesOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

type GetSafeMembersOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

// WithLimit The maximum number of Safes that are returned.
func (s *GetSafesOptions) WithLimit(i int) *GetSafesOptions {
	withLimit(i)(s.query)
	return s
}

// WithOffset Offset of the first Safe that is returned in the collection of results.
func (s *GetSafesOptions) WithOffset(i int) *GetSafesOptions {
	withOffset(i)(s.query)
	return s
}

// WithUseCache Whether to retrieve the cache from a session.
func (s *GetSafesOptions) WithUseCache(b bool) *GetSafesOptions {
	withUseCache(b)(s.query)
	return s
}

// WithSort Sorts according to the safeName property in ascending order (default) or
// descending order to control the sort direction.
func (s *GetSafesOptions) WithSort(_s []string) *GetSafesOptions {
	withSort(_s)(s.query)
	return s
}

// WithSearch Searches according to the Safe name. Search is performed according to the
// REST standard (search="search word").
func (s *GetSafesOptions) WithSearch(_s string) *GetSafesOptions {
	withSearch(_s)(s.query)
	return s
}

// WithIncludeAccounts Whether to return accounts for each Safe as part of the response.
// If not sent, the value will be false.
func (s *GetSafesOptions) WithIncludeAccounts(b bool) *GetSafesOptions {
	withIncludeAccounts(b)(s.query)
	return s
}

// WithExtendedDetails Whether to return all Safe details or only safeName as part of the response.
// If not sent, the value is True.
func (s *GetSafesOptions) WithExtendedDetails(b bool) *GetSafesOptions {
	withExtendedDetails(b)(s.query)
	return s
}

// WithLimit The maximum number of members that are returned.
func (s *GetSafeMembersOptions) WithLimit(i int) *GetSafeMembersOptions {
	withLimit(i)(s.query)
	return s
}

// WithOffset Offset of the first member that is returned in the collection of results.
func (s *GetSafeMembersOptions) WithOffset(i int) *GetSafeMembersOptions {
	withOffset(i)(s.query)
	return s
}

// WithUseCache Whether to retrieve the cache from a session.
func (s *GetSafeMembersOptions) WithUseCache(b bool) *GetSafeMembersOptions {
	withUseCache(b)(s.query)
	return s
}

// WithSort Sorts according to the memberName property in ascending order (default) or descending order to control
// the sort direction. asc: ascending desc:descending
func (s *GetSafeMembersOptions) WithSort(_s []string) *GetSafeMembersOptions {
	withSort(_s)(s.query)
	return s
}

// WithSearch Searches according to the Safe name. Search is performed according to the REST standard (search="search word").
func (s *GetSafeMembersOptions) WithSearch(_s string) *GetSafeMembersOptions {
	withSearch(_s)(s.query)
	return s
}

// WithFilter Filters are according to the REST standard. Search for Safe members using the following filters.
// Multiple filters can be applied using the AND operator. • memberType - returns all members according to the
// type(user or group) Default: both Example: filter= memberType eq user • membershipExpired - returns either
// expired members or members that are not expired. Default: both Example: filter= membershipExpired eq true
// • includePredefinedUsers - includes predefined users in the returned list. Default: False, non-predefined
// users only Example: filter= includePredefinedUsers eq true
func (s *GetSafeMembersOptions) WithFilter(_s string) *GetSafeMembersOptions {
	withFilter(_s)(s.query)
	return s
}
