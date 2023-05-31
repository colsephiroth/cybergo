package pvwa

import "net/url"

type GetUserGroupsOptions struct {
	path  string
	query *url.Values
	pvwa  *PVWA
}

// WithSearch Searches according to the REST standard (searching with "contains"). Search matches when
// all search terms appear in the group name.
func (u *GetUserGroupsOptions) WithSearch(s string) *GetUserGroupsOptions {
	withSearch(s)(u.query)
	return u
}

// WithSort Property or properties by which to sort returned users, followed by asc (default) or desc to control sort direction.
// Separate multiple properties with commas, up to a maximum of three properties: "groupname", "directory", "location"
func (u *GetUserGroupsOptions) WithSort(s []string) *GetUserGroupsOptions {
	withSort(s)(u.query)
	return u
}

// WithFilter Filters according to the REST standard. Search for groups using the following filters:
//
//   - groupType eq <Directory|Vault>
//   - groupName eq <Group_Name>
func (u *GetUserGroupsOptions) WithFilter(s string) *GetUserGroupsOptions {
	withFilter(s)(u.query)
	return u
}

// WithIncludeMembers Whether to return members for each user group as part of the response.
// If not sent, the value will be False.
func (u *GetUserGroupsOptions) WithIncludeMembers(b bool) *GetUserGroupsOptions {
	withIncludeMembers(b)(u.query)
	return u
}
