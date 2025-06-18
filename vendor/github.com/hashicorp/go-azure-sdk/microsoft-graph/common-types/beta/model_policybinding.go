package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyBinding struct {
	// Specifies the users or groups to be explicitly excluded from this policy scope. Can be null or empty.
	Exclusions *[]ScopeBase `json:"exclusions,omitempty"`

	// Specifies the users or groups to be included in this policy scope. Often set to tenantScope for 'All users'.
	Inclusions *[]ScopeBase `json:"inclusions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &PolicyBinding{}

func (s *PolicyBinding) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling PolicyBinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["exclusions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Exclusions into list []json.RawMessage: %+v", err)
		}

		output := make([]ScopeBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalScopeBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Exclusions' for 'PolicyBinding': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Exclusions = &output
	}

	if v, ok := temp["inclusions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Inclusions into list []json.RawMessage: %+v", err)
		}

		output := make([]ScopeBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalScopeBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Inclusions' for 'PolicyBinding': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Inclusions = &output
	}

	return nil
}
