package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInLocation struct {
	// Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the
	// sign-in activity.
	City nullable.Type[string] `json:"city,omitempty"`

	// Provides the country code info (two letter code) where the sign-in originated. This is calculated using
	// latitude/longitude information from the sign-in activity.
	CountryOrRegion nullable.Type[string] `json:"countryOrRegion,omitempty"`

	// Provides the latitude, longitude and altitude where the sign-in originated.
	GeoCoordinates *GeoCoordinates `json:"geoCoordinates,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the
	// sign-in activity.
	State nullable.Type[string] `json:"state,omitempty"`
}
