package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CountryLookupMethodType string

const (
	CountryLookupMethodType_AuthenticatorAppGps CountryLookupMethodType = "authenticatorAppGps"
	CountryLookupMethodType_ClientIPAddress     CountryLookupMethodType = "clientIpAddress"
)

func PossibleValuesForCountryLookupMethodType() []string {
	return []string{
		string(CountryLookupMethodType_AuthenticatorAppGps),
		string(CountryLookupMethodType_ClientIPAddress),
	}
}

func (s *CountryLookupMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCountryLookupMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCountryLookupMethodType(input string) (*CountryLookupMethodType, error) {
	vals := map[string]CountryLookupMethodType{
		"authenticatorappgps": CountryLookupMethodType_AuthenticatorAppGps,
		"clientipaddress":     CountryLookupMethodType_ClientIPAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CountryLookupMethodType(input)
	return &out, nil
}
