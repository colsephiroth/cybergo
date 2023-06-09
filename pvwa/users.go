package pvwa

import (
	"net/url"
)

// GetUsers This method returns a list of all existing users in the Vault except for the Master and Batch built-in users.
// The user who runs this web service must have Audit Users permissions in the Vault. Groups on the same level as your
// user or lower in the Vault hierarchy are retrieved.
//
//   - The groups that are returned depends on the HideVaultUsersTree parameter is defined in the dpbaram.ini file.
//   - If HideVaultUsersTree is set to No, all groups are returned(instead of only those on the same level or lower in the Vault hierarchy).
//   - If HideVaultUsersTree is set to Yes, only auditors and managers are allowed to get all groups.
//
// Note: This Web service returns up to 6000 users in up to 20 seconds. If the number of users is higher, the response time may be higher.
func (p *PVWA) GetUsers() *GetUsersOptions {
	return &GetUsersOptions{
		path:  "API/Users",
		query: &url.Values{},
		pvwa:  p,
	}
}

func (u *GetUsersOptions) Run() ([]*UserDetails, error) {
	res, err := getReturnSingle[GetUsersResponse](u.pvwa, u.path, u.query)
	return res.Users, err
}
