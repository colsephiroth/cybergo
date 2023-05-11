package cybergo

type GetAccountsOptions struct {
	SavedFilter *int
	Search      *string
	Sort        *[]string
	Offset      *int
	Limit       *int
	Filter      *string
	SearchType  *string
}

type UpdateAccountOperation struct {
	Op    string `json:"op,omitempty"`
	Path  string `json:"path,omitempty"`
	Value any    `json:"value,omitempty"`
	From  string `json:"from,omitempty"`
}

type AccountDetails struct {
	CategoryModificationTime  int                   `json:"categoryModificationTime"`
	PlatformId                string                `json:"platformId"`
	SafeName                  string                `json:"safeName"`
	Id                        string                `json:"id"`
	Name                      string                `json:"name"`
	Address                   string                `json:"address"`
	UserName                  string                `json:"userName"`
	SecretType                string                `json:"secretType"`
	Secret                    string                `json:"secret"`
	PlatformAccountProperties map[string]any        `json:"platformAccountProperties"`
	SecretManagement          *SecretManagement     `json:"secretManagement"`
	RemoteMachinesAccess      *RemoteMachinesAccess `json:"remoteMachinesAccess"`
	CreatedTime               int                   `json:"createdTime"`
	DeletionTime              int                   `json:"deletionTime"`
}

type SecretManagement struct {
	AutomaticManagementEnabled bool   `json:"automaticManagementEnabled"`
	ManualManagementReason     string `json:"manualManagementReason"`
	Status                     string `json:"status"`
	LastModifiedTime           int    `json:"lastModifiedTime"`
	LastReconciledTime         int    `json:"lastReconciledTime"`
	LastVerifiedTime           int    `json:"lastVerifiedTime"`
}

type RemoteMachinesAccess struct {
	RemoteMachines                   string `json:"remoteMachines"`
	AccessRestrictedToRemoteMachines bool   `json:"accessRestrictedToRemoteMachines"`
}
