package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodConfiguration = VoiceAuthenticationMethodConfiguration{}

type VoiceAuthenticationMethodConfiguration struct {
	// A collection of groups that are enabled to use the authentication method. Expanded by default.
	IncludeTargets *[]VoiceAuthenticationMethodTarget `json:"includeTargets,omitempty"`

	// true if users can register office phones, otherwise, false.
	IsOfficePhoneAllowed nullable.Type[bool] `json:"isOfficePhoneAllowed,omitempty"`

	// Fields inherited from AuthenticationMethodConfiguration

	// Groups of users that are excluded from a policy.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// The state of the policy. Possible values are: enabled, disabled.
	State *AuthenticationMethodState `json:"state,omitempty"`

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

func (s VoiceAuthenticationMethodConfiguration) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return BaseAuthenticationMethodConfigurationImpl{
		ExcludeTargets: s.ExcludeTargets,
		State:          s.State,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s VoiceAuthenticationMethodConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VoiceAuthenticationMethodConfiguration{}

func (s VoiceAuthenticationMethodConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper VoiceAuthenticationMethodConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VoiceAuthenticationMethodConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VoiceAuthenticationMethodConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.voiceAuthenticationMethodConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VoiceAuthenticationMethodConfiguration: %+v", err)
	}

	return encoded, nil
}
