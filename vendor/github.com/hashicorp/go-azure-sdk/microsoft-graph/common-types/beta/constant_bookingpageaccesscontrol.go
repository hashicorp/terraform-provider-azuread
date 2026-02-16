package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingPageAccessControl string

const (
	BookingPageAccessControl_RestrictedToOrganization BookingPageAccessControl = "restrictedToOrganization"
	BookingPageAccessControl_Unrestricted             BookingPageAccessControl = "unrestricted"
)

func PossibleValuesForBookingPageAccessControl() []string {
	return []string{
		string(BookingPageAccessControl_RestrictedToOrganization),
		string(BookingPageAccessControl_Unrestricted),
	}
}

func (s *BookingPageAccessControl) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingPageAccessControl(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingPageAccessControl(input string) (*BookingPageAccessControl, error) {
	vals := map[string]BookingPageAccessControl{
		"restrictedtoorganization": BookingPageAccessControl_RestrictedToOrganization,
		"unrestricted":             BookingPageAccessControl_Unrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingPageAccessControl(input)
	return &out, nil
}
