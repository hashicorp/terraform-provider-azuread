package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationCombinationConfiguration = Fido2CombinationConfiguration{}

type Fido2CombinationConfiguration struct {
	// A list of AAGUIDs allowed to be used as part of the specified authentication method combinations.
	AllowedAAGUIDs *[]string `json:"allowedAAGUIDs,omitempty"`

	// Fields inherited from AuthenticationCombinationConfiguration

	// Which authentication method combinations this configuration applies to. Must be an allowedCombinations object defined
	// for the authenticationStrengthPolicy. For fido2combinationConfigurations use 'fido2', for
	// x509certificatecombinationconfiguration use 'x509CertificateSingleFactor' or 'x509CertificateMultiFactor'.
	AppliesToCombinations *[]AuthenticationMethodModes `json:"appliesToCombinations,omitempty"`

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

func (s Fido2CombinationConfiguration) AuthenticationCombinationConfiguration() BaseAuthenticationCombinationConfigurationImpl {
	return BaseAuthenticationCombinationConfigurationImpl{
		AppliesToCombinations: s.AppliesToCombinations,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s Fido2CombinationConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Fido2CombinationConfiguration{}

func (s Fido2CombinationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Fido2CombinationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Fido2CombinationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Fido2CombinationConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.fido2CombinationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Fido2CombinationConfiguration: %+v", err)
	}

	return encoded, nil
}
