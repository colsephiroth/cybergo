package pvwa

import "net/url"

type GetUsersOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

// WithExtendedDetails Returns additional user details such as user groups and userDN for LDAP users.
func (u *GetUsersOptions) WithExtendedDetails(b bool) *GetUsersOptions {
	withExtendedDetails(b)(u.query)
	return u
}

// WithSearch Search is according to REST standards using the values, username, firstname, and lastname.
func (u *GetUsersOptions) WithSearch(s string) *GetUsersOptions {
	withSearch(s)(u.query)
	return u
}

// WithSort Property or properties by which to sort returned users, followed by asc (default) or desc to control
// sort direction. Separate multiple properties with commas, up to a maximum of three properties. (*include the properties)
func (u *GetUsersOptions) WithSort(s []string) *GetUsersOptions {
	withSort(s)(u.query)
	return u
}

// WithUserName The name of the user.
func (u *GetUsersOptions) WithUserName(s string) *GetUsersOptions {
	withUserName(s)(u.query)
	return u
}

// WithUserType The user type as defined in the license.
func (u *GetUsersOptions) WithUserType(s string) *GetUsersOptions {
	withUserType(s)(u.query)
	return u
}

// WithComponentUser If the user is a component, then the value is true. Otherwise, it is false.
func (u *GetUsersOptions) WithComponentUser(b bool) *GetUsersOptions {
	withComponentUser(b)(u.query)
	return u
}
