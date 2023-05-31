package ccp

type Response struct {
	Content                 string `json:"Content"`
	CreationMethod          string `json:"CreationMethod"`
	Address                 string `json:"Address"`
	Safe                    string `json:"Safe"`
	UserName                string `json:"UserName"`
	Name                    string `json:"Name"`
	PolicyID                string `json:"PolicyID"`
	DeviceType              string `json:"DeviceType"`
	CPMDisabled             string `json:"CPMDisabled"`
	Folder                  string `json:"Folder"`
	PasswordChangeInProcess string `json:"PasswordChangeInProcess"`
}
