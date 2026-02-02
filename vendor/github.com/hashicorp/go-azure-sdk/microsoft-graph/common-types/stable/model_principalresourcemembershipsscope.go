package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewScope = PrincipalResourceMembershipsScope{}

type PrincipalResourceMembershipsScope struct {
	// Defines the scopes of the principals whose access to resources are reviewed in the access review.
	PrincipalScopes *[]AccessReviewScope `json:"principalScopes,omitempty"`

	// Defines the scopes of the resources for which access is reviewed.
	ResourceScopes *[]AccessReviewScope `json:"resourceScopes,omitempty"`

	// Fields inherited from AccessReviewScope

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PrincipalResourceMembershipsScope) AccessReviewScope() BaseAccessReviewScopeImpl {
	return BaseAccessReviewScopeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrincipalResourceMembershipsScope{}

func (s PrincipalResourceMembershipsScope) MarshalJSON() ([]byte, error) {
	type wrapper PrincipalResourceMembershipsScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrincipalResourceMembershipsScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrincipalResourceMembershipsScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.principalResourceMembershipsScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrincipalResourceMembershipsScope: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrincipalResourceMembershipsScope{}

func (s *PrincipalResourceMembershipsScope) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PrincipalResourceMembershipsScope into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["principalScopes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PrincipalScopes into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessReviewScope, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessReviewScopeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PrincipalScopes' for 'PrincipalResourceMembershipsScope': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PrincipalScopes = &output
	}

	if v, ok := temp["resourceScopes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ResourceScopes into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessReviewScope, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessReviewScopeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ResourceScopes' for 'PrincipalResourceMembershipsScope': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ResourceScopes = &output
	}

	return nil
}
