package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticUpdateMode string

const (
	AutomaticUpdateMode_AutoInstallAndRebootAtMaintenanceTime     AutomaticUpdateMode = "autoInstallAndRebootAtMaintenanceTime"
	AutomaticUpdateMode_AutoInstallAndRebootAtScheduledTime       AutomaticUpdateMode = "autoInstallAndRebootAtScheduledTime"
	AutomaticUpdateMode_AutoInstallAndRebootWithoutEndUserControl AutomaticUpdateMode = "autoInstallAndRebootWithoutEndUserControl"
	AutomaticUpdateMode_AutoInstallAtMaintenanceTime              AutomaticUpdateMode = "autoInstallAtMaintenanceTime"
	AutomaticUpdateMode_NotifyDownload                            AutomaticUpdateMode = "notifyDownload"
	AutomaticUpdateMode_UserDefined                               AutomaticUpdateMode = "userDefined"
)

func PossibleValuesForAutomaticUpdateMode() []string {
	return []string{
		string(AutomaticUpdateMode_AutoInstallAndRebootAtMaintenanceTime),
		string(AutomaticUpdateMode_AutoInstallAndRebootAtScheduledTime),
		string(AutomaticUpdateMode_AutoInstallAndRebootWithoutEndUserControl),
		string(AutomaticUpdateMode_AutoInstallAtMaintenanceTime),
		string(AutomaticUpdateMode_NotifyDownload),
		string(AutomaticUpdateMode_UserDefined),
	}
}

func (s *AutomaticUpdateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticUpdateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticUpdateMode(input string) (*AutomaticUpdateMode, error) {
	vals := map[string]AutomaticUpdateMode{
		"autoinstallandrebootatmaintenancetime":     AutomaticUpdateMode_AutoInstallAndRebootAtMaintenanceTime,
		"autoinstallandrebootatscheduledtime":       AutomaticUpdateMode_AutoInstallAndRebootAtScheduledTime,
		"autoinstallandrebootwithoutendusercontrol": AutomaticUpdateMode_AutoInstallAndRebootWithoutEndUserControl,
		"autoinstallatmaintenancetime":              AutomaticUpdateMode_AutoInstallAtMaintenanceTime,
		"notifydownload":                            AutomaticUpdateMode_NotifyDownload,
		"userdefined":                               AutomaticUpdateMode_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticUpdateMode(input)
	return &out, nil
}
