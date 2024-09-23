package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareAction struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identities the item was shared with in this action.
	Recipients *[]IdentitySet `json:"recipients,omitempty"`
}

var _ json.Unmarshaler = &ShareAction{}

func (s *ShareAction) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling ShareAction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["recipients"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Recipients into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentitySet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentitySetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Recipients' for 'ShareAction': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Recipients = &output
	}

	return nil
}
