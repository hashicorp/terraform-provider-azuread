package models

// User describes a User object.
type User struct {
	ID                           *string   `json:"id,omitempty"`
	AboutMe                      *string   `json:"aboutMe,omitempty"`
	AccountEnabled               *bool     `json:"accountEnabled,omitempty"`
	BusinessPhones               *[]string `json:"businessPhones,omitempty"`
	City                         *string   `json:"city,omitempty"`
	CompanyName                  *string   `json:"companyName,omitempty"`
	Country                      *string   `json:"country,omitempty"`
	CreationType                 *string   `json:"creationType,omitempty"`
	Department                   *string   `json:"department,omitempty"`
	DisplayName                  *string   `json:"displayName,omitempty"`
	EmployeeId                   *string   `json:"employeeId,omitempty"`
	ExternalUserState            *string   `json:"externalUserState,omitempty"`
	FaxNumber                    *string   `json:"faxNumber,omitempty"`
	GivenName                    *string   `json:"givenName,omitempty"`
	ImAddresses                  *[]string `json:"imAddresses,omitempty"`
	Interests                    *[]string `json:"interests,omitempty"`
	JobTitle                     *string   `json:"jobTitle,omitempty"`
	Mail                         *string   `json:"mail,omitempty"`
	MailNickname                 *string   `json:"mailNickname,omitempty"`
	MobilePhone                  *string   `json:"mobilePhone,omitempty"`
	MySite                       *string   `json:"mySite,omitempty"`
	OfficeLocation               *string   `json:"officeLocation,omitempty"`
	OnPremisesDistinguishedName  *string   `json:"onPremisesDistinguishedName,omitempty"`
	OnPremisesDomainName         *string   `json:"onPremisesDomainName,omitempty"`
	OnPremisesImmutableId        *string   `json:"onPremisesImmutableId,omitempty"`
	OnPremisesSamAccountName     *string   `json:"onPremisesSamAccountName,omitempty"`
	OnPremisesSecurityIdentifier *string   `json:"onPremisesSecurityIdentifier,omitempty"`
	OnPremisesSyncEnabled        *bool     `json:"onPremisesSyncEnabled,omitempty"`
	OnPremisesUserPrincipalName  *string   `json:"onPremisesUserPrincipalName,omitempty"`
	OtherMails                   *[]string `json:"otherMails,omitempty"`
	PasswordPolicies             *string   `json:"passwordPolicies,omitempty"`
	PastProjects                 *[]string `json:"pastProjects,omitempty"`
	PostalCode                   *string   `json:"postalCode,omitempty"`
	PreferredDataLocation        *string   `json:"preferredDataLocation,omitempty"`
	PreferredLanguage            *string   `json:"preferredLanguage,omitempty"`
	PreferredName                *string   `json:"preferredName,omitempty"`
	ProxyAddresses               *[]string `json:"proxyAddresses,omitempty"`
	Responsibilities             *[]string `json:"responsibilities,omitempty"`
	Schools                      *[]string `json:"schools,omitempty"`
	ShowInAddressList            *bool     `json:"showInAddressList,omitempty"`
	Skills                       *[]string `json:"skills,omitempty"`
	State                        *string   `json:"state,omitempty"`
	StreetAddress                *string   `json:"streetAddress,omitempty"`
	Surname                      *string   `json:"surname,omitempty"`
	UsageLocation                *string   `json:"usageLocation,omitempty"`
	UserPrincipalName            *string   `json:"userPrincipalName,omitempty"`
	UserType                     *string   `json:"userType,omitempty"`

	PasswordProfile *UserPasswordProfile `json:"passwordProfile,omitempty"`
}

type UserPasswordProfile struct {
	ForceChangePasswordNextSignIn        *bool   `json:"forceChangePasswordNextSignIn,omitempty"`
	ForceChangePasswordNextSignInWithMfa *bool   `json:"forceChangePasswordNextSignInWithMfa,omitempty"`
	Password                             *string `json:"password,omitempty"`
}
