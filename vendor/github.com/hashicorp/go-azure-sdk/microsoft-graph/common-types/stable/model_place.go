package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Place interface {
	Entity
	Place() BasePlaceImpl
}

var _ Place = BasePlaceImpl{}

type BasePlaceImpl struct {
	// The street address of the place.
	Address *PhysicalAddress `json:"address,omitempty"`

	// The name associated with the place.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies the place location in latitude, longitude, and (optionally) altitude coordinates.
	GeoCoordinates *OutlookGeoCoordinates `json:"geoCoordinates,omitempty"`

	// The phone number of the place.
	Phone nullable.Type[string] `json:"phone,omitempty"`

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

func (s BasePlaceImpl) Place() BasePlaceImpl {
	return s
}

func (s BasePlaceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Place = RawPlaceImpl{}

// RawPlaceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPlaceImpl struct {
	place  BasePlaceImpl
	Type   string
	Values map[string]interface{}
}

func (s RawPlaceImpl) Place() BasePlaceImpl {
	return s.place
}

func (s RawPlaceImpl) Entity() BaseEntityImpl {
	return s.place.Entity()
}

var _ json.Marshaler = BasePlaceImpl{}

func (s BasePlaceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePlaceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePlaceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePlaceImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.place"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePlaceImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPlaceImplementation(input []byte) (Place, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Place into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.room") {
		var out Room
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Room: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roomList") {
		var out RoomList
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoomList: %+v", err)
		}
		return out, nil
	}

	var parent BasePlaceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlaceImpl: %+v", err)
	}

	return RawPlaceImpl{
		place:  parent,
		Type:   value,
		Values: temp,
	}, nil

}
