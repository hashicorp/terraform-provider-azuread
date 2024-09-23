package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodConfiguration = Fido2AuthenticationMethodConfiguration{}

type Fido2AuthenticationMethodConfiguration struct {
	// A collection of groups that are enabled to use the authentication method.
	IncludeTargets *[]PasskeyAuthenticationMethodTarget `json:"includeTargets,omitempty"`

	// Determines whether attestation must be enforced for FIDO2 security key registration.
	IsAttestationEnforced nullable.Type[bool] `json:"isAttestationEnforced,omitempty"`

	// Determines if users can register new FIDO2 security keys.
	IsSelfServiceRegistrationAllowed nullable.Type[bool] `json:"isSelfServiceRegistrationAllowed,omitempty"`

	// Controls whether key restrictions are enforced on FIDO2 security keys, either allowing or disallowing certain key
	// types as defined by Authenticator Attestation GUID (AAGUID), an identifier that indicates the type (e.g. make and
	// model) of the authenticator.
	KeyRestrictions *Fido2KeyRestrictions `json:"keyRestrictions,omitempty"`

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

func (s Fido2AuthenticationMethodConfiguration) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return BaseAuthenticationMethodConfigurationImpl{
		ExcludeTargets: s.ExcludeTargets,
		State:          s.State,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s Fido2AuthenticationMethodConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Fido2AuthenticationMethodConfiguration{}

func (s Fido2AuthenticationMethodConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Fido2AuthenticationMethodConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Fido2AuthenticationMethodConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Fido2AuthenticationMethodConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.fido2AuthenticationMethodConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Fido2AuthenticationMethodConfiguration: %+v", err)
	}

	return encoded, nil
}
