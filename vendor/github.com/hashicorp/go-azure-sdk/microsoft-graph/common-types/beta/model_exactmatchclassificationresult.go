package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchClassificationResult struct {
	Classification *[]ExactMatchDetectedSensitiveContent `json:"classification,omitempty"`
	Errors         *[]ClassificationError                `json:"errors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &ExactMatchClassificationResult{}

func (s *ExactMatchClassificationResult) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Classification *[]ExactMatchDetectedSensitiveContent `json:"classification,omitempty"`
		ODataId        *string                               `json:"@odata.id,omitempty"`
		ODataType      *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Classification = decoded.Classification
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExactMatchClassificationResult into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["errors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Errors into list []json.RawMessage: %+v", err)
		}

		output := make([]ClassificationError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalClassificationErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Errors' for 'ExactMatchClassificationResult': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Errors = &output
	}

	return nil
}
