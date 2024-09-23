package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationType string

const (
	LocationType_BusinessAddress LocationType = "businessAddress"
	LocationType_ConferenceRoom  LocationType = "conferenceRoom"
	LocationType_Default         LocationType = "default"
	LocationType_GeoCoordinates  LocationType = "geoCoordinates"
	LocationType_HomeAddress     LocationType = "homeAddress"
	LocationType_Hotel           LocationType = "hotel"
	LocationType_LocalBusiness   LocationType = "localBusiness"
	LocationType_PostalAddress   LocationType = "postalAddress"
	LocationType_Restaurant      LocationType = "restaurant"
	LocationType_StreetAddress   LocationType = "streetAddress"
)

func PossibleValuesForLocationType() []string {
	return []string{
		string(LocationType_BusinessAddress),
		string(LocationType_ConferenceRoom),
		string(LocationType_Default),
		string(LocationType_GeoCoordinates),
		string(LocationType_HomeAddress),
		string(LocationType_Hotel),
		string(LocationType_LocalBusiness),
		string(LocationType_PostalAddress),
		string(LocationType_Restaurant),
		string(LocationType_StreetAddress),
	}
}

func (s *LocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationType(input string) (*LocationType, error) {
	vals := map[string]LocationType{
		"businessaddress": LocationType_BusinessAddress,
		"conferenceroom":  LocationType_ConferenceRoom,
		"default":         LocationType_Default,
		"geocoordinates":  LocationType_GeoCoordinates,
		"homeaddress":     LocationType_HomeAddress,
		"hotel":           LocationType_Hotel,
		"localbusiness":   LocationType_LocalBusiness,
		"postaladdress":   LocationType_PostalAddress,
		"restaurant":      LocationType_Restaurant,
		"streetaddress":   LocationType_StreetAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationType(input)
	return &out, nil
}
