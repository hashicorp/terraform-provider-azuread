package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VerticalSection{}

type VerticalSection struct {
	// Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural,
	// soft, strong, unknownFutureValue.
	Emphasis *SectionEmphasisType `json:"emphasis,omitempty"`

	// The set of web parts in this section.
	Webparts *[]WebPart `json:"webparts,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s VerticalSection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VerticalSection{}

func (s VerticalSection) MarshalJSON() ([]byte, error) {
	type wrapper VerticalSection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VerticalSection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VerticalSection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.verticalSection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VerticalSection: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &VerticalSection{}

func (s *VerticalSection) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Emphasis  *SectionEmphasisType `json:"emphasis,omitempty"`
		Id        *string              `json:"id,omitempty"`
		ODataId   *string              `json:"@odata.id,omitempty"`
		ODataType *string              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Emphasis = decoded.Emphasis
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VerticalSection into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["webparts"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Webparts into list []json.RawMessage: %+v", err)
		}

		output := make([]WebPart, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWebPartImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Webparts' for 'VerticalSection': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Webparts = &output
	}

	return nil
}
