package pvwa

type AccountGroup struct {
	ID         string `json:"GroupID,omitempty"`
	Name       string `json:"GroupName,omitempty"`
	PlatformID string `json:"GroupPlatformID,omitempty"`
	Safe       string `json:"Safe"`
}
