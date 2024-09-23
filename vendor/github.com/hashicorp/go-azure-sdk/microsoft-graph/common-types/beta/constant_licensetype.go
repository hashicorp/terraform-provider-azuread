package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseType string

const (
	LicenseType_NotPaid LicenseType = "notPaid"
	LicenseType_Paid    LicenseType = "paid"
	LicenseType_Trial   LicenseType = "trial"
)

func PossibleValuesForLicenseType() []string {
	return []string{
		string(LicenseType_NotPaid),
		string(LicenseType_Paid),
		string(LicenseType_Trial),
	}
}

func (s *LicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLicenseType(input string) (*LicenseType, error) {
	vals := map[string]LicenseType{
		"notpaid": LicenseType_NotPaid,
		"paid":    LicenseType_Paid,
		"trial":   LicenseType_Trial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LicenseType(input)
	return &out, nil
}
