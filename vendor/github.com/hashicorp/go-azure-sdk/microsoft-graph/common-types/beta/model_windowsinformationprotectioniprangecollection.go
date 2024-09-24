package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionIPRangeCollection struct {
	// Display name
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Collection of ip ranges
	Ranges *[]IPRange `json:"ranges,omitempty"`
}

var _ json.Unmarshaler = &WindowsInformationProtectionIPRangeCollection{}

func (s *WindowsInformationProtectionIPRangeCollection) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName *string `json:"displayName,omitempty"`
		ODataId     *string `json:"@odata.id,omitempty"`
		ODataType   *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsInformationProtectionIPRangeCollection into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["ranges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Ranges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Ranges' for 'WindowsInformationProtectionIPRangeCollection': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Ranges = &output
	}

	return nil
}
