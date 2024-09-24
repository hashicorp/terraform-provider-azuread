package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserRegistrationDetails{}

type UserRegistrationDetails struct {
	// The method the user or admin selected as default for performing multifactor authentication for the user. The possible
	// values are: none, mobilePhone, alternateMobilePhone, officePhone, microsoftAuthenticatorPush,
	// softwareOneTimePasscode, unknownFutureValue.
	DefaultMfaMethod *DefaultMfaMethodType `json:"defaultMfaMethod,omitempty"`

	// Indicates whether the user has an admin role in the tenant. This value can be used to check the authentication
	// methods that privileged accounts are registered for and capable of.
	IsAdmin nullable.Type[bool] `json:"isAdmin,omitempty"`

	// Indicates whether the user has registered a strong authentication method for multifactor authentication. The method
	// must be allowed by the authentication methods policy. Supports $filter (eq).
	IsMfaCapable *bool `json:"isMfaCapable,omitempty"`

	// Indicates whether the user has registered a strong authentication method for multifactor authentication. The method
	// may not necessarily be allowed by the authentication methods policy. Supports $filter (eq).
	IsMfaRegistered *bool `json:"isMfaRegistered,omitempty"`

	// Indicates whether the user has registered a passwordless strong authentication method (including FIDO2, Windows Hello
	// for Business, and Microsoft Authenticator (Passwordless)) that is allowed by the authentication methods policy.
	// Supports $filter (eq).
	IsPasswordlessCapable *bool `json:"isPasswordlessCapable,omitempty"`

	// Indicates whether the user has registered the required number of authentication methods for self-service password
	// reset and the user is allowed to perform self-service password reset by policy. Supports $filter (eq).
	IsSsprCapable *bool `json:"isSsprCapable,omitempty"`

	// Indicates whether the user is allowed to perform self-service password reset by policy. The user may not necessarily
	// have registered the required number of authentication methods for self-service password reset. Supports $filter (eq).
	IsSsprEnabled *bool `json:"isSsprEnabled,omitempty"`

	// Indicates whether the user has registered the required number of authentication methods for self-service password
	// reset. The user may not necessarily be allowed to perform self-service password reset by policy. Supports $filter
	// (eq).
	IsSsprRegistered *bool `json:"isSsprRegistered,omitempty"`

	// Indicates whether system preferred authentication method is enabled. If enabled, the system dynamically determines
	// the most secure authentication method among the methods registered by the user. Supports $filter (eq).
	IsSystemPreferredAuthenticationMethodEnabled nullable.Type[bool] `json:"isSystemPreferredAuthenticationMethodEnabled,omitempty"`

	// The date and time (UTC) when the report was last updated. The DateTimeOffset type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// Collection of authentication methods registered, such as mobilePhone, email, passKeyDeviceBound. Supports $filter
	// (any with eq).
	MethodsRegistered *[]string `json:"methodsRegistered,omitempty"`

	// Collection of authentication methods that the system determined to be the most secure authentication methods among
	// the registered methods for second factor authentication. Possible values are: push, oath, voiceMobile,
	// voiceAlternateMobile, voiceOffice, sms, none, unknownFutureValue. Supports $filter (any with eq).
	SystemPreferredAuthenticationMethods *[]string `json:"systemPreferredAuthenticationMethods,omitempty"`

	// The user display name, such as Adele Vance. Supports $filter (eq, startsWith) and $orderby.
	UserDisplayName *string `json:"userDisplayName,omitempty"`

	// The method the user selected as the default second-factor for performing multifactor authentication. Possible values
	// are: push, oath, voiceMobile, voiceAlternateMobile, voiceOffice, sms, none, unknownFutureValue.
	UserPreferredMethodForSecondaryAuthentication *UserDefaultAuthenticationMethod `json:"userPreferredMethodForSecondaryAuthentication,omitempty"`

	// The user principal name, such as AdeleV@contoso.com. Supports $filter (eq, startsWith) and $orderby.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`

	// Identifies whether the user is a member or guest in the tenant. The possible values are: member, guest,
	// unknownFutureValue.
	UserType *SignInUserType `json:"userType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserRegistrationDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserRegistrationDetails{}

func (s UserRegistrationDetails) MarshalJSON() ([]byte, error) {
	type wrapper UserRegistrationDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserRegistrationDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserRegistrationDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userRegistrationDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserRegistrationDetails: %+v", err)
	}

	return encoded, nil
}
