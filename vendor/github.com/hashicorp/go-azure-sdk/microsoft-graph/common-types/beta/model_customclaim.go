package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomClaimBase = CustomClaim{}

type CustomClaim struct {
	// The name of the claim to be emitted.
	Name *string `json:"name,omitempty"`

	// An optional namespace to be included as part of the claim name.
	Namespace nullable.Type[string] `json:"namespace,omitempty"`

	// If specified, it sets the nameFormat attribute associated with the claim in the SAML response. The possible values
	// are: unspecified, uri, basic, unknownFutureValue.
	SamlAttributeNameFormat *SamlAttributeNameFormat `json:"samlAttributeNameFormat,omitempty"`

	// List of token formats for which this claim should be emitted. The possible values are: saml,jwt, unknownFutureValue
	TokenFormat *[]TokenFormat `json:"tokenFormat,omitempty"`

	// Fields inherited from CustomClaimBase

	// One or more configurations that describe how the claim is sourced and under what conditions.
	Configurations *[]CustomClaimConfiguration `json:"configurations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CustomClaim) CustomClaimBase() BaseCustomClaimBaseImpl {
	return BaseCustomClaimBaseImpl{
		Configurations: s.Configurations,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

var _ json.Marshaler = CustomClaim{}

func (s CustomClaim) MarshalJSON() ([]byte, error) {
	type wrapper CustomClaim
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomClaim: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomClaim: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customClaim"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomClaim: %+v", err)
	}

	return encoded, nil
}
