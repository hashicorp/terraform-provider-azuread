package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomClaimConfiguration struct {
	// The attribute on which we source this property.
	Attribute CustomClaimAttributeBase `json:"attribute"`

	// The condition, if any, associated with this configuration.
	Condition CustomClaimConditionBase `json:"condition"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// An ordered list of transformations that are applied in sequence.
	Transformations *[]CustomClaimTransformation `json:"transformations,omitempty"`
}

var _ json.Unmarshaler = &CustomClaimConfiguration{}

func (s *CustomClaimConfiguration) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling CustomClaimConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attribute"]; ok {
		impl, err := UnmarshalCustomClaimAttributeBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Attribute' for 'CustomClaimConfiguration': %+v", err)
		}
		s.Attribute = impl
	}

	if v, ok := temp["condition"]; ok {
		impl, err := UnmarshalCustomClaimConditionBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Condition' for 'CustomClaimConfiguration': %+v", err)
		}
		s.Condition = impl
	}

	if v, ok := temp["transformations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Transformations into list []json.RawMessage: %+v", err)
		}

		output := make([]CustomClaimTransformation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCustomClaimTransformationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Transformations' for 'CustomClaimConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Transformations = &output
	}

	return nil
}
