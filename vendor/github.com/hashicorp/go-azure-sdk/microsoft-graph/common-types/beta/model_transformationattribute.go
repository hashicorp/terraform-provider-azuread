package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransformationAttribute struct {
	Attribute CustomClaimAttributeBase `json:"attribute"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// This flag is only relevant in the case where the attribute is multivalued. By default, transformations are only
	// applied to the first element in a multi-valued claim, however setting this flag to true ensures the transformation is
	// applied to all values, resulting in a multivalued output.
	TreatAsMultiValue *bool `json:"treatAsMultiValue,omitempty"`
}

var _ json.Unmarshaler = &TransformationAttribute{}

func (s *TransformationAttribute) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId           *string `json:"@odata.id,omitempty"`
		ODataType         *string `json:"@odata.type,omitempty"`
		TreatAsMultiValue *bool   `json:"treatAsMultiValue,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.TreatAsMultiValue = decoded.TreatAsMultiValue

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TransformationAttribute into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attribute"]; ok {
		impl, err := UnmarshalCustomClaimAttributeBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Attribute' for 'TransformationAttribute': %+v", err)
		}
		s.Attribute = impl
	}

	return nil
}
