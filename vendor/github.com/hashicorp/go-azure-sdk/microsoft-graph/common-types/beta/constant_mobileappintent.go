package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntent string

const (
	MobileAppIntent_Available                         MobileAppIntent = "available"
	MobileAppIntent_AvailableInstallWithoutEnrollment MobileAppIntent = "availableInstallWithoutEnrollment"
	MobileAppIntent_Exclude                           MobileAppIntent = "exclude"
	MobileAppIntent_NotAvailable                      MobileAppIntent = "notAvailable"
	MobileAppIntent_RequiredAndAvailableInstall       MobileAppIntent = "requiredAndAvailableInstall"
	MobileAppIntent_RequiredInstall                   MobileAppIntent = "requiredInstall"
	MobileAppIntent_RequiredUninstall                 MobileAppIntent = "requiredUninstall"
)

func PossibleValuesForMobileAppIntent() []string {
	return []string{
		string(MobileAppIntent_Available),
		string(MobileAppIntent_AvailableInstallWithoutEnrollment),
		string(MobileAppIntent_Exclude),
		string(MobileAppIntent_NotAvailable),
		string(MobileAppIntent_RequiredAndAvailableInstall),
		string(MobileAppIntent_RequiredInstall),
		string(MobileAppIntent_RequiredUninstall),
	}
}

func (s *MobileAppIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppIntent(input string) (*MobileAppIntent, error) {
	vals := map[string]MobileAppIntent{
		"available":                         MobileAppIntent_Available,
		"availableinstallwithoutenrollment": MobileAppIntent_AvailableInstallWithoutEnrollment,
		"exclude":                           MobileAppIntent_Exclude,
		"notavailable":                      MobileAppIntent_NotAvailable,
		"requiredandavailableinstall":       MobileAppIntent_RequiredAndAvailableInstall,
		"requiredinstall":                   MobileAppIntent_RequiredInstall,
		"requireduninstall":                 MobileAppIntent_RequiredUninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppIntent(input)
	return &out, nil
}
