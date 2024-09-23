package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CustomClaimsPolicy{}

type CustomClaimsPolicy struct {
	// If specified, it overrides the content of the audience claim for WS-Federation and SAML2 protocols. A custom signing
	// key must be used for audienceOverride to be applied, otherwise, the audienceOverride value is ignored. The value
	// provided must be in the format of an absolute URI.
	AudienceOverride nullable.Type[string] `json:"audienceOverride,omitempty"`

	// Defines which claims are present in the tokens affected by the policy, in addition to the basic claim and the core
	// claim set. Inherited from customclaimbase.
	Claims *[]CustomClaimBase `json:"claims,omitempty"`

	// Indicates whether the application ID is added to the claim. It is relevant only for SAML2.0 and if a custom signing
	// key is used. the default value is true. Optional.
	IncludeApplicationIdInIssuer nullable.Type[bool] `json:"includeApplicationIdInIssuer,omitempty"`

	// Determines whether the basic claim set is included in tokens affected by this policy. If set to true, all claims in
	// the basic claim set are emitted in tokens affected by the policy. By default the basic claim set isn't in the tokens
	// unless they're explicitly configured in this policy.
	IncludeBasicClaimSet nullable.Type[bool] `json:"includeBasicClaimSet,omitempty"`

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

func (s CustomClaimsPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomClaimsPolicy{}

func (s CustomClaimsPolicy) MarshalJSON() ([]byte, error) {
	type wrapper CustomClaimsPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomClaimsPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomClaimsPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customClaimsPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomClaimsPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CustomClaimsPolicy{}

func (s *CustomClaimsPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AudienceOverride             nullable.Type[string] `json:"audienceOverride,omitempty"`
		IncludeApplicationIdInIssuer nullable.Type[bool]   `json:"includeApplicationIdInIssuer,omitempty"`
		IncludeBasicClaimSet         nullable.Type[bool]   `json:"includeBasicClaimSet,omitempty"`
		Id                           *string               `json:"id,omitempty"`
		ODataId                      *string               `json:"@odata.id,omitempty"`
		ODataType                    *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AudienceOverride = decoded.AudienceOverride
	s.IncludeApplicationIdInIssuer = decoded.IncludeApplicationIdInIssuer
	s.IncludeBasicClaimSet = decoded.IncludeBasicClaimSet
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CustomClaimsPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["claims"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Claims into list []json.RawMessage: %+v", err)
		}

		output := make([]CustomClaimBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCustomClaimBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Claims' for 'CustomClaimsPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Claims = &output
	}

	return nil
}
