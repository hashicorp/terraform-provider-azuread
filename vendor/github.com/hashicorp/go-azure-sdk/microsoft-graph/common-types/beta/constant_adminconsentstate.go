package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentState string

const (
	AdminConsentState_Granted       AdminConsentState = "granted"
	AdminConsentState_NotConfigured AdminConsentState = "notConfigured"
	AdminConsentState_NotGranted    AdminConsentState = "notGranted"
)

func PossibleValuesForAdminConsentState() []string {
	return []string{
		string(AdminConsentState_Granted),
		string(AdminConsentState_NotConfigured),
		string(AdminConsentState_NotGranted),
	}
}

func (s *AdminConsentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdminConsentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdminConsentState(input string) (*AdminConsentState, error) {
	vals := map[string]AdminConsentState{
		"granted":       AdminConsentState_Granted,
		"notconfigured": AdminConsentState_NotConfigured,
		"notgranted":    AdminConsentState_NotGranted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdminConsentState(input)
	return &out, nil
}
