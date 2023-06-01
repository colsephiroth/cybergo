package pvwa

import (
	"net/url"
)

// GetUserGroups This method returns a list of all existing user groups in the Vault. The user who runs this web
// service must have Audit users permissions in the Vault. Groups on the same level as your user or
// lower in the Vault hierarchy are retrieved.
//
//   - The groups that are returned depends on how the HideVaultUsersTree parameter is defined in the dpbaram.ini file.
//   - If HideVaultUsersTree is set to No, all groups are returned(instead of only those on the same level or lower in the Vault hierarchy).
//   - If HideVaultUsersTree is set to Yes, only auditors and managers are allowed to get all groups.
//   - Retrieving more than 1,000 groups may cause a slowdown in the response.
func (p *PVWA) GetUserGroups() *GetUserGroupsOptions {
	return &GetUserGroupsOptions{
		path:  "/Api/UserGroups",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (u *GetUserGroupsOptions) Run() ([]*UserGroup, error) {
	return genericGetReturnSlice[UserGroup](u.pvwa, u.path, u.query)
}
