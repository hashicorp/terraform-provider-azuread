package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPhoneNumberRedirectLevel string

const (
	ManagedAppPhoneNumberRedirectLevel_AllApps     ManagedAppPhoneNumberRedirectLevel = "allApps"
	ManagedAppPhoneNumberRedirectLevel_Blocked     ManagedAppPhoneNumberRedirectLevel = "blocked"
	ManagedAppPhoneNumberRedirectLevel_CustomApp   ManagedAppPhoneNumberRedirectLevel = "customApp"
	ManagedAppPhoneNumberRedirectLevel_ManagedApps ManagedAppPhoneNumberRedirectLevel = "managedApps"
)

func PossibleValuesForManagedAppPhoneNumberRedirectLevel() []string {
	return []string{
		string(ManagedAppPhoneNumberRedirectLevel_AllApps),
		string(ManagedAppPhoneNumberRedirectLevel_Blocked),
		string(ManagedAppPhoneNumberRedirectLevel_CustomApp),
		string(ManagedAppPhoneNumberRedirectLevel_ManagedApps),
	}
}

func (s *ManagedAppPhoneNumberRedirectLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppPhoneNumberRedirectLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppPhoneNumberRedirectLevel(input string) (*ManagedAppPhoneNumberRedirectLevel, error) {
	vals := map[string]ManagedAppPhoneNumberRedirectLevel{
		"allapps":     ManagedAppPhoneNumberRedirectLevel_AllApps,
		"blocked":     ManagedAppPhoneNumberRedirectLevel_Blocked,
		"customapp":   ManagedAppPhoneNumberRedirectLevel_CustomApp,
		"managedapps": ManagedAppPhoneNumberRedirectLevel_ManagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppPhoneNumberRedirectLevel(input)
	return &out, nil
}
