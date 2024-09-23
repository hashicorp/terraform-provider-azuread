package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CredentialUserRegistrationDetails{}

type CredentialUserRegistrationDetails struct {
	// Represents the authentication method that the user has registered. Possible values are: email, mobilePhone,
	// officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode,
	// alternateMobilePhone (supported only in registration), fido, appPassword, unknownFutureValue.
	AuthMethods *[]RegistrationAuthMethod `json:"authMethods,omitempty"`

	// Indicates whether the user is ready to perform self-service password reset or MFA.
	IsCapable *bool `json:"isCapable,omitempty"`

	// Indicates whether the user enabled to perform self-service password reset.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Indicates whether the user is registered for MFA.
	IsMfaRegistered *bool `json:"isMfaRegistered,omitempty"`

	// Indicates whether the user has registered any authentication methods for self-service password reset.
	IsRegistered *bool `json:"isRegistered,omitempty"`

	// Provides the user name of the corresponding user.
	UserDisplayName *string `json:"userDisplayName,omitempty"`

	// Provides the user principal name of the corresponding user.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`

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

func (s CredentialUserRegistrationDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CredentialUserRegistrationDetails{}

func (s CredentialUserRegistrationDetails) MarshalJSON() ([]byte, error) {
	type wrapper CredentialUserRegistrationDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CredentialUserRegistrationDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CredentialUserRegistrationDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.credentialUserRegistrationDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CredentialUserRegistrationDetails: %+v", err)
	}

	return encoded, nil
}
