package pvwa

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

type SafeMember struct {
	SafeUrlId                 string                 `json:"safeUrlId,omitempty"`
	SafeName                  string                 `json:"safeName,omitempty"`
	SafeNumber                int                    `json:"safeNumber,omitempty"`
	MemberId                  string                 `json:"memberId,omitempty"`
	MemberName                string                 `json:"memberName,omitempty"`
	MemberType                string                 `json:"memberType,omitempty"`
	MembershipExpirationDate  int                    `json:"membershipExpirationDate,omitempty"`
	IsExpiredMembershipEnable bool                   `json:"isExpiredMembershipEnable,omitempty"`
	IsPredefinedUser          bool                   `json:"isPredefinedUser,omitempty"`
	IsReadOnly                bool                   `json:"isReadOnly,omitempty"`
	Permissions               *SafeMemberPermissions `json:"permissions,omitempty"`
}

// SafeMemberPermissions The permissions that the user or group has in this Safe
type SafeMemberPermissions struct {
	// Use accounts but not view passwords.
	UseAccounts bool `json:"useAccounts"`
	// Retrieve and view accounts in the Safe.
	RetrieveAccounts bool `json:"retrieveAccounts"`
	// View Accounts list.
	ListAccounts bool `json:"listAccounts"`
	// Add accounts in the Safe. Users who have this permission automatically
	// have UpdateAccountProperties permissions.
	AddAccounts bool `json:"addAccounts"`
	// Update existing account content.
	UpdateAccountContent bool `json:"updateAccountContent"`
	// Update existing account properties.
	UpdateAccountProperties bool `json:"updateAccountProperties"`
	// Initiate password management operations through CPM such as changing,
	// verifying, and reconciling passwords. When this parameter is set to False,
	// the SpecifyNextAccountContent parameter is also automatically set to False.
	InitiateCPMAccountManagementOperations bool `json:"initiateCPMAccountManagementOperations"`
	// Specify the password that is used when the CPM changes the password value.
	// This parameter can only be specified when the InitiateCPMAccountManagementOperations
	// parameter is set to True. When InitiateCPMAccountManagementOperations is set
	// to False this parameter is automatically set to False.
	SpecifyNextAccountContent bool `json:"specifyNextAccountContent"`
	// Rename existing accounts in the Safe.
	RenameAccounts bool `json:"renameAccounts"`
	// Delete existing passwords in the Safe.
	DeleteAccounts bool `json:"deleteAccounts"`
	// Unlock accounts that are locked by other users.
	UnlockAccounts bool `json:"unlockAccounts"`
	// Perform administrative tasks in the Safe, including:
	// 	- Update Safe properties
	// 	- Recover the Safe
	//	- Delete the Safe
	ManageSafe bool `json:"manageSafe"`
	// Add and remove Safe members, and update their authorizations in the Safe.
	ManageSafeMembers bool `json:"manageSafeMembers"`
	// Create a backup of a Safe and its contents, and store in another location.
	BackupSafe bool `json:"backupSafe"`
	// View account and user activity in the Safe.
	ViewAuditLog bool `json:"viewAuditLog"`
	// View permissions of Safe members.
	ViewSafeMembers bool `json:"viewSafeMembers"`
	// Request Authorization Level 1.
	RequestsAuthorizationLevel1 bool `json:"requestsAuthorizationLevel1"`
	// Request Authorization Level 2.
	RequestsAuthorizationLevel2 bool `json:"requestsAuthorizationLevel2"`
	// Access the Safe without confirmation from authorized users. This overrides
	// the Safe properties that specify that Safe members require confirmation to access the Safe.
	AccessWithoutConfirmation bool `json:"accessWithoutConfirmation"`
	// Create folders in the Safe.
	CreateFolders bool `json:"createFolders"`
	// Delete folders from the Safe.
	DeleteFolders bool `json:"deleteFolders"`
	// Move accounts and folders in the Safe to different folders and subfolders.
	MoveAccountsAndFolders bool `json:"moveAccountsAndFolders"`
}
