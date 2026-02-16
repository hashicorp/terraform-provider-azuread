package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppActionType string

const (
	MobileAppActionType_InstallCommandSent   MobileAppActionType = "installCommandSent"
	MobileAppActionType_Installed            MobileAppActionType = "installed"
	MobileAppActionType_Uninstalled          MobileAppActionType = "uninstalled"
	MobileAppActionType_Unknown              MobileAppActionType = "unknown"
	MobileAppActionType_UserRequestedInstall MobileAppActionType = "userRequestedInstall"
)

func PossibleValuesForMobileAppActionType() []string {
	return []string{
		string(MobileAppActionType_InstallCommandSent),
		string(MobileAppActionType_Installed),
		string(MobileAppActionType_Uninstalled),
		string(MobileAppActionType_Unknown),
		string(MobileAppActionType_UserRequestedInstall),
	}
}

func (s *MobileAppActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppActionType(input string) (*MobileAppActionType, error) {
	vals := map[string]MobileAppActionType{
		"installcommandsent":   MobileAppActionType_InstallCommandSent,
		"installed":            MobileAppActionType_Installed,
		"uninstalled":          MobileAppActionType_Uninstalled,
		"unknown":              MobileAppActionType_Unknown,
		"userrequestedinstall": MobileAppActionType_UserRequestedInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppActionType(input)
	return &out, nil
}
