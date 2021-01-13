package models

import (
	"fmt"
	"time"
)

// ServicePrincipal describes a Service Principal object.
type ServicePrincipal struct {
	ID                                  *string                       `json:"id,omitempty"`
	AccountEnabled                      *bool                         `json:"accountEnabled,omitempty"`
	AddIns                              *[]AddIn                      `json:"addIns,omitempty"`
	AlternativeNames                    *[]string                     `json:"alternativeNames,omitempty"`
	AppDisplayName                      *string                       `json:"appDisplayName,omitempty"`
	AppId                               *string                       `json:"appId,omitempty"`
	ApplicationTemplateId               *string                       `json:"applicationTemplateId,omitempty"`
	AppOwnerOrganizationId              *string                       `json:"appOwnerOrganizationId,omitempty"`
	AppRoleAssignmentRequired           *bool                         `json:"appRoleAssignmentRequired,omitempty"`
	AppRoles                            *[]AppRole                    `json:"appRoles,omitempty"`
	DeletedDateTime                     *time.Time                    `json:"deletedDateTime,omitempty"`
	DisplayName                         *string                       `json:"displayName,omitempty"`
	Homepage                            *string                       `json:"homepage,omitempty"`
	Info                                *InformationalUrl             `json:"info,omitempty"`
	KeyCredentials                      *[]KeyCredential              `json:"keyCredentials,omitempty"`
	LoginUrl                            *string                       `json:"loginUrl,omitempty"`
	LogoutUrl                           *string                       `json:"logoutUrl,omitempty"`
	NotificationEmailAddresses          *[]string                     `json:"notificationEmailAddresses,omitempty"`
	PasswordCredentials                 *[]PasswordCredential         `json:"passwordCredentials,omitempty"`
	PasswordSingleSignOnSettings        *PasswordSingleSignOnSettings `json:"passwordSingleSignOnSettings,omitempty"`
	PreferredSingleSignOnMode           *string                       `json:"preferredSingleSignOnMode,omitempty"`
	PreferredTokenSigningKeyEndDateTime *time.Time                    `json:"preferredTokenSigningKeyEndDateTime,omitempty"`
	PublishedPermissionScopes           *[]PermissionScope            `json:"publishedPermissionScopes,omitempty"`
	ReplyUrls                           *[]string                     `json:"replyUrls,omitempty"`
	SamlSingleSignOnSettings            *SamlSingleSignOnSettings     `json:"samlSingleSignOnSettings,omitempty"`
	ServicePrincipalNames               *[]string                     `json:"servicePrincipalNames,omitempty"`
	ServicePrincipalType                *string                       `json:"servicePrincipalType,omitempty"`
	SignInAudience                      SignInAudience                `json:"signInAudience,omitempty"`
	Tags                                *[]string                     `json:"tags,omitempty"`
	TokenEncryptionKeyId                *string                       `json:"tokenEncryptionKeyId,omitempty"`
	VerifiedPublisher                   *VerifiedPublisher            `json:"verifiedPublisher,omitempty"`

	Owners *[]string `json:"owners@odata.bind,omitempty"`
}

// AppendOwner appends a new owner object URI to the Owners slice.
func (a *ServicePrincipal) AppendOwner(endpoint string, apiVersion string, id string) {
	val := fmt.Sprintf("%s/%s/directoryObjects/%s", endpoint, apiVersion, id)
	var owners []string
	if a.Owners != nil {
		owners = *a.Owners
	}
	owners = append(owners, val)
	a.Owners = &owners
}

type PasswordSingleSignOnSettings struct {
	Fields *[]SingleSignOnField `json:"fields,omitempty"`
}

type SamlSingleSignOnSettings struct {
	RelayState *string `json:"relayState,omitempty"`
}

type SingleSignOnField struct {
	CustomizedLabel *string `json:"customizedLabel,omitempty"`
	DefaultLabel    *string `json:"defaultLabel,omitempty"`
	FieldId         *string `json:"fieldId,omitempty"`
	Type            *string `json:"type,omitempty"`
}

type VerifiedPublisher struct {
	AddedDateTime       *time.Time `json:"addedDateTime,omitempty"`
	DisplayName         *string    `json:"displayName,omitempty"`
	VerifiedPublisherId *string    `json:"verifiedPublisherId,omitempty"`
}
