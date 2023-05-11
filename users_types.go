package cybergo

type GetUsersResponse struct {
	Users []*UserDetails `json:"Users"`
	Total int            `json:"Total"`
}

type UserDetails struct {
	ID                 int                       `json:"id,omitempty"`
	Username           string                    `json:"username"`
	Source             string                    `json:"source,omitempty"`
	UserType           string                    `json:"userType,omitempty"`
	ComponentUser      bool                      `json:"componentUser,omitempty"`
	GroupsMembership   []*GroupMembershipDetails `json:"groupsMembership,omitempty"`
	UserDN             string                    `json:"userDN,omitempty"`
	VaultAuthorization []string                  `json:"vaultAuthorization,omitempty"`
	Location           string                    `json:"location,omitempty"`
	PersonalDetails    *UserPersonalDetails      `json:"personalDetails,omitempty"`
	EnableUser         bool                      `json:"enableUser,omitempty"`
	Suspended          bool                      `json:"suspended,omitempty"`
}

type GroupMembershipDetails struct {
	GroupID   int    `json:"groupID,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	GroupType string `json:"groupType,omitempty"`
}

type UserPersonalDetails struct {
	FirstName    string `json:"firstName,omitempty"`
	MiddleName   string `json:"middleName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	Organization string `json:"organization,omitempty"`
	Department   string `json:"department,omitempty"`
}
