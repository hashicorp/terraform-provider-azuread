package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NamedLocation = CountryNamedLocation{}

type CountryNamedLocation struct {
	// List of countries and/or regions in two-letter format specified by ISO 3166-2. Required.
	CountriesAndRegions []string `json:"countriesAndRegions"`

	// Determines what method is used to decide which country the user is located in. Possible values are
	// clientIpAddress(default) and authenticatorAppGps. Note: authenticatorAppGps is not yet supported in the Microsoft
	// Cloud for US Government.
	CountryLookupMethod *CountryLookupMethodType `json:"countryLookupMethod,omitempty"`

	// true if IP addresses that don't map to a country or region should be included in the named location. Optional.
	// Default value is false.
	IncludeUnknownCountriesAndRegions *bool `json:"includeUnknownCountriesAndRegions,omitempty"`

	// Fields inherited from NamedLocation

	// The Timestamp type represents creation date and time of the location using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Human-readable name of the location.
	DisplayName *string `json:"displayName,omitempty"`

	// The Timestamp type represents last modified date and time of the location using ISO 8601 format and is always in UTC
	// time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

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

func (s CountryNamedLocation) NamedLocation() BaseNamedLocationImpl {
	return BaseNamedLocationImpl{
		CreatedDateTime:  s.CreatedDateTime,
		DisplayName:      s.DisplayName,
		ModifiedDateTime: s.ModifiedDateTime,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s CountryNamedLocation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CountryNamedLocation{}

func (s CountryNamedLocation) MarshalJSON() ([]byte, error) {
	type wrapper CountryNamedLocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CountryNamedLocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CountryNamedLocation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.countryNamedLocation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CountryNamedLocation: %+v", err)
	}

	return encoded, nil
}
