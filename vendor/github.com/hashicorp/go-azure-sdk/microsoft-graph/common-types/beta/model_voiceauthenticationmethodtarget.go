package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodTarget = VoiceAuthenticationMethodTarget{}

type VoiceAuthenticationMethodTarget struct {

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

func (s VoiceAuthenticationMethodTarget) AuthenticationMethodTarget() BaseAuthenticationMethodTargetImpl {
	return BaseAuthenticationMethodTargetImpl{
		IsRegistrationRequired: s.IsRegistrationRequired,
		TargetType:             s.TargetType,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s VoiceAuthenticationMethodTarget) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VoiceAuthenticationMethodTarget{}

func (s VoiceAuthenticationMethodTarget) MarshalJSON() ([]byte, error) {
	type wrapper VoiceAuthenticationMethodTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VoiceAuthenticationMethodTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VoiceAuthenticationMethodTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.voiceAuthenticationMethodTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VoiceAuthenticationMethodTarget: %+v", err)
	}

	return encoded, nil
}
