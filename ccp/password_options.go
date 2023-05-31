package ccp

import "net/url"

type GetPasswordOptions struct {
	path  string
	query *url.Values
	ccp   *CCP
}

// WithSafe Specifies the name of the Safe where the password is stored.
func (p *GetPasswordOptions) WithSafe(s string) *GetPasswordOptions {
	withSafe(s)(p.query)
	return p
}

// WithFolder Specifies the name of the folder where the password is stored.
func (p *GetPasswordOptions) WithFolder(s string) *GetPasswordOptions {
	withFolder(s)(p.query)
	return p
}

// WithObject Specifies the name of the password object to retrieve.
func (p *GetPasswordOptions) WithObject(s string) *GetPasswordOptions {
	withObject(s)(p.query)
	return p
}

// WithUserName Defines search criteria according to the UserName account property.
func (p *GetPasswordOptions) WithUserName(s string) *GetPasswordOptions {
	withUserName(s)(p.query)
	return p
}

// WithAddress Defines search criteria according to the Address account property.
func (p *GetPasswordOptions) WithAddress(s string) *GetPasswordOptions {
	withAddress(s)(p.query)
	return p
}

// WithDatabase Defines search criteria according to the Database account property.
func (p *GetPasswordOptions) WithDatabase(s string) *GetPasswordOptions {
	withDatabase(s)(p.query)
	return p
}

// WithPolicyID Defines the format that will be used in the setPolicyID method.
func (p *GetPasswordOptions) WithPolicyID(s string) *GetPasswordOptions {
	withPolicyID(s)(p.query)
	return p
}

// WithReason The reason for retrieving the password. This reason will be audited in the Credential Provider audit log.
func (p *GetPasswordOptions) WithReason(s string) *GetPasswordOptions {
	withReason(s)(p.query)
	return p
}

// WithConnectionTimeout The number of seconds that the Central Credential Provider will try to retrieve the password.
// The timeout is calculated when the request is sent from the web service to the Vault and returned back to the web service.
func (p *GetPasswordOptions) WithConnectionTimeout(i int) *GetPasswordOptions {
	withConnectionTimeout(i)(p.query)
	return p
}

// WithQuery Defines a free query using account properties, including Safe, folder, and object.
// When this method is specified, all other search criteria (Safe/Folder/ Object/UserName/Address/PolicyID/Database)
// are ignored and only the account properties that are specified in the query are passed to the
// Central Credential Provider in the password request.
func (p *GetPasswordOptions) WithQuery(s string) *GetPasswordOptions {
	withQuery(s)(p.query)
	return p
}

// WithQueryFormat Defines the query format, which can optionally use regular expressions. Possible values are:
//   - Exact
//   - Regexp
func (p *GetPasswordOptions) WithQueryFormat(s string) *GetPasswordOptions {
	withQueryFormat(s)(p.query)
	return p
}

// WithFailRequestOnPasswordChange Whether an error will be returned if this web service is called
// when a password change process is underway.
func (p *GetPasswordOptions) WithFailRequestOnPasswordChange(b bool) *GetPasswordOptions {
	withFailRequestOnPasswordChange(b)(p.query)
	return p
}
