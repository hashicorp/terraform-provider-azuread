package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppManagementPolicyActorExemptions struct {
	CustomSecurityAttributes *[]CustomSecurityAttributeExemption `json:"customSecurityAttributes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AppManagementPolicyActorExemptions{}

func (s *AppManagementPolicyActorExemptions) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling AppManagementPolicyActorExemptions into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["customSecurityAttributes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CustomSecurityAttributes into list []json.RawMessage: %+v", err)
		}

		output := make([]CustomSecurityAttributeExemption, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCustomSecurityAttributeExemptionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CustomSecurityAttributes' for 'AppManagementPolicyActorExemptions': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CustomSecurityAttributes = &output
	}

	return nil
}
