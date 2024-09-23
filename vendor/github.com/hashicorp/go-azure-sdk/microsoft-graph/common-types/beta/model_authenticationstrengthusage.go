package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthUsage struct {
	Mfa  *[]ConditionalAccessPolicy `json:"mfa,omitempty"`
	None *[]ConditionalAccessPolicy `json:"none,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AuthenticationStrengthUsage{}

func (s *AuthenticationStrengthUsage) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling AuthenticationStrengthUsage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["mfa"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Mfa into list []json.RawMessage: %+v", err)
		}

		output := make([]ConditionalAccessPolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalConditionalAccessPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Mfa' for 'AuthenticationStrengthUsage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Mfa = &output
	}

	if v, ok := temp["none"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling None into list []json.RawMessage: %+v", err)
		}

		output := make([]ConditionalAccessPolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalConditionalAccessPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'None' for 'AuthenticationStrengthUsage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.None = &output
	}

	return nil
}
