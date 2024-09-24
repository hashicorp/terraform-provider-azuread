package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentWrapper struct {
	Classification *[]DetectedSensitiveContent `json:"classification,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &DetectedSensitiveContentWrapper{}

func (s *DetectedSensitiveContentWrapper) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling DetectedSensitiveContentWrapper into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["classification"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Classification into list []json.RawMessage: %+v", err)
		}

		output := make([]DetectedSensitiveContent, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDetectedSensitiveContentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Classification' for 'DetectedSensitiveContentWrapper': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Classification = &output
	}

	return nil
}
