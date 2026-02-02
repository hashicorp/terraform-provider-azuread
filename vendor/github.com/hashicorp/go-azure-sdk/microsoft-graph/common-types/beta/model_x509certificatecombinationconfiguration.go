package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationCombinationConfiguration = X509CertificateCombinationConfiguration{}

type X509CertificateCombinationConfiguration struct {
	// A list of allowed subject key identifier values.
	AllowedIssuerSkis *[]string `json:"allowedIssuerSkis,omitempty"`

	// A list of allowed policy OIDs.
	AllowedPolicyOIDs *[]string `json:"allowedPolicyOIDs,omitempty"`

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

func (s X509CertificateCombinationConfiguration) AuthenticationCombinationConfiguration() BaseAuthenticationCombinationConfigurationImpl {
	return BaseAuthenticationCombinationConfigurationImpl{
		AppliesToCombinations: s.AppliesToCombinations,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s X509CertificateCombinationConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = X509CertificateCombinationConfiguration{}

func (s X509CertificateCombinationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper X509CertificateCombinationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling X509CertificateCombinationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling X509CertificateCombinationConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.x509CertificateCombinationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling X509CertificateCombinationConfiguration: %+v", err)
	}

	return encoded, nil
}
