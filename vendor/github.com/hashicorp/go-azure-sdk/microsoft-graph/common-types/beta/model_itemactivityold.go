package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ItemActivityOLD{}

type ItemActivityOLD struct {
	Action    *ItemActionSet       `json:"action,omitempty"`
	Actor     IdentitySet          `json:"actor"`
	DriveItem *DriveItem           `json:"driveItem,omitempty"`
	ListItem  *ListItem            `json:"listItem,omitempty"`
	Times     *ItemActivityTimeSet `json:"times,omitempty"`

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

func (s ItemActivityOLD) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemActivityOLD{}

func (s ItemActivityOLD) MarshalJSON() ([]byte, error) {
	type wrapper ItemActivityOLD
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemActivityOLD: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemActivityOLD: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.itemActivityOLD"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemActivityOLD: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ItemActivityOLD{}

func (s *ItemActivityOLD) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action    *ItemActionSet       `json:"action,omitempty"`
		DriveItem *DriveItem           `json:"driveItem,omitempty"`
		ListItem  *ListItem            `json:"listItem,omitempty"`
		Times     *ItemActivityTimeSet `json:"times,omitempty"`
		Id        *string              `json:"id,omitempty"`
		ODataId   *string              `json:"@odata.id,omitempty"`
		ODataType *string              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.DriveItem = decoded.DriveItem
	s.ListItem = decoded.ListItem
	s.Times = decoded.Times
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ItemActivityOLD into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actor"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Actor' for 'ItemActivityOLD': %+v", err)
		}
		s.Actor = impl
	}

	return nil
}
