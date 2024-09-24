package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Location = LocationConstraintItem{}

type LocationConstraintItem struct {
	// If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set
	// to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache
	// without checking if it's free. Default is true.
	ResolveAvailability nullable.Type[bool] `json:"resolveAvailability,omitempty"`

	// Fields inherited from Location

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

func (s LocationConstraintItem) Location() BaseLocationImpl {
	return BaseLocationImpl{
		Address:              s.Address,
		Coordinates:          s.Coordinates,
		DisplayName:          s.DisplayName,
		LocationEmailAddress: s.LocationEmailAddress,
		LocationType:         s.LocationType,
		LocationUri:          s.LocationUri,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
		UniqueId:             s.UniqueId,
		UniqueIdType:         s.UniqueIdType,
	}
}

var _ json.Marshaler = LocationConstraintItem{}

func (s LocationConstraintItem) MarshalJSON() ([]byte, error) {
	type wrapper LocationConstraintItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LocationConstraintItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LocationConstraintItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.locationConstraintItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LocationConstraintItem: %+v", err)
	}

	return encoded, nil
}
