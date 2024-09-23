package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Place = RoomList{}

type RoomList struct {
	// The email address of the room list.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	Rooms      *[]Room      `json:"rooms,omitempty"`
	Workspaces *[]Workspace `json:"workspaces,omitempty"`

	// Fields inherited from Place

	// The street address of the place.
	Address *PhysicalAddress `json:"address,omitempty"`

	// The name associated with the place.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies the place location in latitude, longitude, and (optionally) altitude coordinates.
	GeoCoordinates *OutlookGeoCoordinates `json:"geoCoordinates,omitempty"`

	// The phone number of the place.
	Phone nullable.Type[string] `json:"phone,omitempty"`

	// A unique, immutable identifier for the place. Read-only. The value of this identifier is equal to the
	// ExternalDirectoryObjectId returned from the Get-Mailbox cmdlet.
	PlaceId nullable.Type[string] `json:"placeId,omitempty"`

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

func (s RoomList) Place() BasePlaceImpl {
	return BasePlaceImpl{
		Address:        s.Address,
		DisplayName:    s.DisplayName,
		GeoCoordinates: s.GeoCoordinates,
		Phone:          s.Phone,
		PlaceId:        s.PlaceId,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s RoomList) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RoomList{}

func (s RoomList) MarshalJSON() ([]byte, error) {
	type wrapper RoomList
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RoomList: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RoomList: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.roomList"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RoomList: %+v", err)
	}

	return encoded, nil
}
