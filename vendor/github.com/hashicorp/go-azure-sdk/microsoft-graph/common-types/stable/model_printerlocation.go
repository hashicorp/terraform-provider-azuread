package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterLocation struct {
	// The altitude, in meters, that the printer is located at.
	AltitudeInMeters nullable.Type[int64] `json:"altitudeInMeters,omitempty"`

	// The building that the printer is located in.
	Building nullable.Type[string] `json:"building,omitempty"`

	// The city that the printer is located in.
	City nullable.Type[string] `json:"city,omitempty"`

	// The country or region that the printer is located in.
	CountryOrRegion nullable.Type[string] `json:"countryOrRegion,omitempty"`

	// The floor that the printer is located on. Only numerical values are supported right now.
	Floor nullable.Type[string] `json:"floor,omitempty"`

	// The description of the floor that the printer is located on.
	FloorDescription nullable.Type[string] `json:"floorDescription,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
	Organization *[]string `json:"organization,omitempty"`

	// The postal code that the printer is located in.
	PostalCode nullable.Type[string] `json:"postalCode,omitempty"`

	// The description of the room that the printer is located in.
	RoomDescription nullable.Type[string] `json:"roomDescription,omitempty"`

	// The room that the printer is located in. Only numerical values are supported right now.
	RoomName nullable.Type[string] `json:"roomName,omitempty"`

	// The site that the printer is located in.
	Site nullable.Type[string] `json:"site,omitempty"`

	// The state or province that the printer is located in.
	StateOrProvince nullable.Type[string] `json:"stateOrProvince,omitempty"`

	// The street address where the printer is located.
	StreetAddress nullable.Type[string] `json:"streetAddress,omitempty"`

	// The subdivision that the printer is located in. The elements should be in hierarchical order.
	Subdivision *[]string `json:"subdivision,omitempty"`

	Subunit *[]string `json:"subunit,omitempty"`
}
