package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Place = Room{}

type Room struct {
	// Specifies the name of the audio device in the room.
	AudioDeviceName nullable.Type[string] `json:"audioDeviceName,omitempty"`

	// Type of room. Possible values are standard, and reserved.
	BookingType *BookingType `json:"bookingType,omitempty"`

	// Specifies the building name or building number that the room is in.
	Building nullable.Type[string] `json:"building,omitempty"`

	// Specifies the capacity of the room.
	Capacity nullable.Type[int64] `json:"capacity,omitempty"`

	// Specifies the name of the display device in the room.
	DisplayDeviceName nullable.Type[string] `json:"displayDeviceName,omitempty"`

	// Email address of the room.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Specifies a descriptive label for the floor, for example, P.
	FloorLabel nullable.Type[string] `json:"floorLabel,omitempty"`

	// Specifies the floor number that the room is on.
	FloorNumber nullable.Type[int64] `json:"floorNumber,omitempty"`

	// Specifies whether the room is wheelchair accessible.
	IsWheelChairAccessible nullable.Type[bool] `json:"isWheelChairAccessible,omitempty"`

	// Specifies a descriptive label for the room, for example, a number or name.
	Label nullable.Type[string] `json:"label,omitempty"`

	// Specifies a nickname for the room, for example, 'conf room'.
	Nickname *string `json:"nickname,omitempty"`

	// Specifies other features of the room, for example, details like the type of view or furniture type.
	Tags *[]string `json:"tags,omitempty"`

	// Specifies the name of the video device in the room.
	VideoDeviceName nullable.Type[string] `json:"videoDeviceName,omitempty"`

	// Fields inherited from Place

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

func (s Room) Place() BasePlaceImpl {
	return BasePlaceImpl{
		Address:        s.Address,
		DisplayName:    s.DisplayName,
		GeoCoordinates: s.GeoCoordinates,
		Phone:          s.Phone,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s Room) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Room{}

func (s Room) MarshalJSON() ([]byte, error) {
	type wrapper Room
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Room: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Room: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.room"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Room: %+v", err)
	}

	return encoded, nil
}
