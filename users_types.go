package cybergo

import (
	"net/url"
	"strconv"
	"strings"
)

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

type GetUsersOption func(o *url.URL)

func buildGetUsersPath(options ...GetUsersOption) *url.URL {
	path, err := url.Parse("API/Users")
	if err != nil {
		return nil
	}

	for _, option := range options {
		option(path)
	}

	path.RawQuery = path.Query().Encode()

	return path
}

func WithExtendedDetails(b bool) GetUsersOption {
	return func(o *url.URL) {
		o.Query().Add("ExtendedDetails", strconv.FormatBool(b))
	}
}

func WithSearch(s string) GetUsersOption {
	return func(o *url.URL) {
		o.Query().Add("search", s)
	}
}

func WithSort(s []string) GetUsersOption {
	return func(o *url.URL) {
		o.Query().Add("sort", strings.Join(s, ","))
	}
}

func WithUserName(s string) GetUsersOption {
	return func(o *url.URL) {
		o.Query().Add("userName", s)
	}
}

func WithUserType(s string) GetUsersOption {
	return func(o *url.URL) {
		o.Query().Add("userType", s)
	}
}

func WithComponentUser(b bool) GetUsersOption {
	return func(o *url.URL) {
		o.Query().Add("componentUser", strconv.FormatBool(b))
	}
}
