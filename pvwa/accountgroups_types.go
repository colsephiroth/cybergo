package pvwa

type AccountGroup struct {
	ID         string `json:"GroupID,omitempty"`
	Name       string `json:"GroupName,omitempty"`
	PlatformID string `json:"GroupPlatformID,omitempty"`
	Safe       string `json:"Safe"`
}

type AccountGroupMember struct {
	AccountID  string `json:"AccountID,omitempty"`
	SafeName   string `json:"SafeName,omitempty"`
	PlatformID string `json:"PlatformID,omitempty"`
	Address    string `json:"Address,omitempty"`
	UserName   string `json:"UserName,omitempty"`
}
