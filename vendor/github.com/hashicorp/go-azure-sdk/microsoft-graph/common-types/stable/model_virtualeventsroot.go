package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VirtualEventsRoot{}

type VirtualEventsRoot struct {
	Events *[]VirtualEvent `json:"events,omitempty"`

	// A collection of town halls. Nullable.
	Townhalls *[]VirtualEventTownhall `json:"townhalls,omitempty"`

	// A collection of webinars. Nullable.
	Webinars *[]VirtualEventWebinar `json:"webinars,omitempty"`

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

func (s VirtualEventsRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventsRoot{}

func (s VirtualEventsRoot) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventsRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventsRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventsRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventsRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventsRoot: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &VirtualEventsRoot{}

func (s *VirtualEventsRoot) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Townhalls *[]VirtualEventTownhall `json:"townhalls,omitempty"`
		Webinars  *[]VirtualEventWebinar  `json:"webinars,omitempty"`
		Id        *string                 `json:"id,omitempty"`
		ODataId   *string                 `json:"@odata.id,omitempty"`
		ODataType *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Townhalls = decoded.Townhalls
	s.Webinars = decoded.Webinars
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VirtualEventsRoot into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["events"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Events into list []json.RawMessage: %+v", err)
		}

		output := make([]VirtualEvent, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalVirtualEventImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Events' for 'VirtualEventsRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Events = &output
	}

	return nil
}
