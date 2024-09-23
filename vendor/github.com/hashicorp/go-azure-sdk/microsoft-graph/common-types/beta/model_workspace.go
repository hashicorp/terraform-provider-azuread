package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Place = Workspace{}

type Workspace struct {
	// Specifies the building name or building number that the workspace is in.
	Building nullable.Type[string] `json:"building,omitempty"`

	// Specifies the capacity of the workspace.
	Capacity nullable.Type[int64] `json:"capacity,omitempty"`

	// Email address of the workspace.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Specifies a descriptive label for the floor, for example, P.
	FloorLabel nullable.Type[string] `json:"floorLabel,omitempty"`

	// Specifies the floor number that the workspace is on.
	FloorNumber nullable.Type[int64] `json:"floorNumber,omitempty"`

	// Specifies whether the workspace is wheelchair accessible.
	IsWheelChairAccessible nullable.Type[bool] `json:"isWheelChairAccessible,omitempty"`

	// Specifies a descriptive label for the workspace, for example, a number or name.
	Label nullable.Type[string] `json:"label,omitempty"`

	// Specifies a nickname for the workspace, for example, 'quiet workspace'.
	Nickname *string `json:"nickname,omitempty"`

	// Specifies other features of the workspace; for example, the type of view or furniture type.
	Tags *[]string `json:"tags,omitempty"`

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

func (s Workspace) Place() BasePlaceImpl {
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

func (s Workspace) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Workspace{}

func (s Workspace) MarshalJSON() ([]byte, error) {
	type wrapper Workspace
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Workspace: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Workspace: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workspace"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Workspace: %+v", err)
	}

	return encoded, nil
}
