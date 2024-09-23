package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Location interface {
	Location() BaseLocationImpl
}

var _ Location = BaseLocationImpl{}

type BaseLocationImpl struct {
	// The street address of the location.
	Address *PhysicalAddress `json:"address,omitempty"`

	// The geographic coordinates and elevation of the location.
	Coordinates *OutlookGeoCoordinates `json:"coordinates,omitempty"`

	// The name associated with the location.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Optional email address of the location.
	LocationEmailAddress nullable.Type[string] `json:"locationEmailAddress,omitempty"`

	// The type of location. Possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates,
	// streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
	LocationType *LocationType `json:"locationType,omitempty"`

	// Optional URI representing the location.
	LocationUri nullable.Type[string] `json:"locationUri,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// For internal use only.
	UniqueId nullable.Type[string] `json:"uniqueId,omitempty"`

	// For internal use only.
	UniqueIdType *LocationUniqueIdType `json:"uniqueIdType,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseLocationImpl) Location() BaseLocationImpl {
	return s
}

var _ Location = RawLocationImpl{}

// RawLocationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLocationImpl struct {
	location BaseLocationImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawLocationImpl) Location() BaseLocationImpl {
	return s.location
}

var _ json.Marshaler = BaseLocationImpl{}

func (s BaseLocationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseLocationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseLocationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseLocationImpl: %+v", err)
	}

	delete(decoded, "locationType")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseLocationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalLocationImplementation(input []byte) (Location, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Location into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.locationConstraintItem") {
		var out LocationConstraintItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocationConstraintItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseLocationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseLocationImpl: %+v", err)
	}

	return RawLocationImpl{
		location: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
