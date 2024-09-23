package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomClaimBase = SamlNameIdClaim{}

type SamlNameIdClaim struct {
	NameIdFormat *SamlNameIDFormat `json:"nameIdFormat,omitempty"`

	// Allows the specification of a service provider name qualifier reflected in the sAML response. The value provided must
	// match one of the service provider names configured for the application and is only applicable for IdP-initiated
	// applications (the sign-on URL should be empty for the IdP-initiated applications), in all other cases this value is
	// ignored.
	ServiceProviderNameQualifier nullable.Type[string] `json:"serviceProviderNameQualifier,omitempty"`

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

func (s SamlNameIdClaim) CustomClaimBase() BaseCustomClaimBaseImpl {
	return BaseCustomClaimBaseImpl{
		Configurations: s.Configurations,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

var _ json.Marshaler = SamlNameIdClaim{}

func (s SamlNameIdClaim) MarshalJSON() ([]byte, error) {
	type wrapper SamlNameIdClaim
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SamlNameIdClaim: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SamlNameIdClaim: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.samlNameIdClaim"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SamlNameIdClaim: %+v", err)
	}

	return encoded, nil
}
