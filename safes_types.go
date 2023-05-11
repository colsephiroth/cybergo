package cybergo

type GetSafesResponse struct{}

type Safe struct {
	SafeNumber                int            `json:"safeNumber,omitempty"`
	Location                  string         `json:"location,omitempty"`
	Creator                   *SafeCreator   `json:"creator,omitempty"`
	Accounts                  []*SafeAccount `json:"accounts,omitempty"`
	OlacEnabled               bool           `json:"olacEnabled,omitempty"`
	NumberOfVersionsRetention int            `json:"numberOfVersionsRetention,omitempty"`
	NumberOfDaysRetention     int            `json:"numberOfDaysRetention,omitempty"`
	AutoPurgeEnabled          bool           `json:"autoPurgeEnabled,omitempty"`
	CreationTime              int            `json:"creationTime,omitempty"`
	LastModificationTime      int            `json:"lastModificationTime,omitempty"`
	SafeUrlId                 string         `json:"safeUrlId,omitempty"`
	SafeName                  string         `json:"safeName,omitempty"`
	Description               string         `json:"description,omitempty"`
	ManagingCPM               string         `json:"managingCPM,omitempty"`
	IsExpiredMember           bool           `json:"isExpiredMember,omitempty"`
}

type SafeCreator struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SafeAccount struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
