package pvwa

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

type ChangeCredentialsNow struct {
	// Whether the CPM will change the credentials in all the accounts that belong to the same group. This parameter
	// is only relevant for accounts that belong to an account group, and if this parameter does not belong to a group,
	// it will be ignored. If this account is part of an account group and this value is not specified,
	// the default value will be applied.
	ChangeEntireGroup bool `json:"ChangeEntireGroup"`
}

type AccountContentPrerequisites struct {
	// The reason that is required to retrieve the password/SSH key
	Reason string `json:"Reason,omitempty"`
	// The ticket ID of the ticketing system
	TicketId string `json:"TicketId,omitempty"`
	// The name of the ticketing system
	TicketingSystem string `json:"TicketingSystem,omitempty"`
	// Internal parameter (for PSMP only)
	IsUse bool `json:"isUse,omitempty"`
	// The action this password/SSH key will be used for. (Show/Copy/Connect for password; Retrieve for SSH key)
	ActionType string `json:"ActionType,omitempty"`
	// The address of the remote machine the user wants to connect to using the password/SSH key
	Machine string `json:"Machine,omitempty"`
	// The version number of the required password/SSH key. Must be a positive number. If the value is left empty
	// or the value passed does not exist, then the current password/SSH key version is returned
	Version int `json:"Version,omitempty"`
}
