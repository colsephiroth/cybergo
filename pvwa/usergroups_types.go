package pvwa

type UserGroup struct {
	ID          int                 `json:"id,omitempty"`
	GroupType   string              `json:"groupType,omitempty"`
	Directory   string              `json:"directory,omitempty"`
	DN          string              `json:"dn,omitempty"`
	Members     []*UserGroupMembers `json:"members,omitempty"`
	GroupName   string              `json:"groupName,omitempty"`
	Description string              `json:"description,omitempty"`
	Location    string              `json:"location,omitempty"`
}

type UserGroupMembers struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}
