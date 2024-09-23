package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodTarget = SmsAuthenticationMethodTarget{}

type SmsAuthenticationMethodTarget struct {
	// Determines if users can use this authentication method to sign in to Microsoft Entra ID. true if users can use this
	// method for primary authentication, otherwise false.
	IsUsableForSignIn *bool `json:"isUsableForSignIn,omitempty"`

	// Fields inherited from AuthenticationMethodTarget

	// Determines if the user is enforced to register the authentication method.
	IsRegistrationRequired *bool `json:"isRegistrationRequired,omitempty"`

	TargetType *AuthenticationMethodTargetType `json:"targetType,omitempty"`

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

func (s SmsAuthenticationMethodTarget) AuthenticationMethodTarget() BaseAuthenticationMethodTargetImpl {
	return BaseAuthenticationMethodTargetImpl{
		IsRegistrationRequired: s.IsRegistrationRequired,
		TargetType:             s.TargetType,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s SmsAuthenticationMethodTarget) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SmsAuthenticationMethodTarget{}

func (s SmsAuthenticationMethodTarget) MarshalJSON() ([]byte, error) {
	type wrapper SmsAuthenticationMethodTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SmsAuthenticationMethodTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SmsAuthenticationMethodTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.smsAuthenticationMethodTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SmsAuthenticationMethodTarget: %+v", err)
	}

	return encoded, nil
}
