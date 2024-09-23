package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingPriceType string

const (
	BookingPriceType_CallUs      BookingPriceType = "callUs"
	BookingPriceType_FixedPrice  BookingPriceType = "fixedPrice"
	BookingPriceType_Free        BookingPriceType = "free"
	BookingPriceType_Hourly      BookingPriceType = "hourly"
	BookingPriceType_NotSet      BookingPriceType = "notSet"
	BookingPriceType_PriceVaries BookingPriceType = "priceVaries"
	BookingPriceType_StartingAt  BookingPriceType = "startingAt"
	BookingPriceType_Undefined   BookingPriceType = "undefined"
)

func PossibleValuesForBookingPriceType() []string {
	return []string{
		string(BookingPriceType_CallUs),
		string(BookingPriceType_FixedPrice),
		string(BookingPriceType_Free),
		string(BookingPriceType_Hourly),
		string(BookingPriceType_NotSet),
		string(BookingPriceType_PriceVaries),
		string(BookingPriceType_StartingAt),
		string(BookingPriceType_Undefined),
	}
}

func (s *BookingPriceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingPriceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingPriceType(input string) (*BookingPriceType, error) {
	vals := map[string]BookingPriceType{
		"callus":      BookingPriceType_CallUs,
		"fixedprice":  BookingPriceType_FixedPrice,
		"free":        BookingPriceType_Free,
		"hourly":      BookingPriceType_Hourly,
		"notset":      BookingPriceType_NotSet,
		"pricevaries": BookingPriceType_PriceVaries,
		"startingat":  BookingPriceType_StartingAt,
		"undefined":   BookingPriceType_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingPriceType(input)
	return &out, nil
}
