package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLockerApplicationControlType string

const (
	AppLockerApplicationControlType_AuditComponentsAndStoreApps              AppLockerApplicationControlType = "auditComponentsAndStoreApps"
	AppLockerApplicationControlType_AuditComponentsStoreAppsAndSmartlocker   AppLockerApplicationControlType = "auditComponentsStoreAppsAndSmartlocker"
	AppLockerApplicationControlType_EnforceComponentsAndStoreApps            AppLockerApplicationControlType = "enforceComponentsAndStoreApps"
	AppLockerApplicationControlType_EnforceComponentsStoreAppsAndSmartlocker AppLockerApplicationControlType = "enforceComponentsStoreAppsAndSmartlocker"
	AppLockerApplicationControlType_NotConfigured                            AppLockerApplicationControlType = "notConfigured"
)

func PossibleValuesForAppLockerApplicationControlType() []string {
	return []string{
		string(AppLockerApplicationControlType_AuditComponentsAndStoreApps),
		string(AppLockerApplicationControlType_AuditComponentsStoreAppsAndSmartlocker),
		string(AppLockerApplicationControlType_EnforceComponentsAndStoreApps),
		string(AppLockerApplicationControlType_EnforceComponentsStoreAppsAndSmartlocker),
		string(AppLockerApplicationControlType_NotConfigured),
	}
}

func (s *AppLockerApplicationControlType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppLockerApplicationControlType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppLockerApplicationControlType(input string) (*AppLockerApplicationControlType, error) {
	vals := map[string]AppLockerApplicationControlType{
		"auditcomponentsandstoreapps":              AppLockerApplicationControlType_AuditComponentsAndStoreApps,
		"auditcomponentsstoreappsandsmartlocker":   AppLockerApplicationControlType_AuditComponentsStoreAppsAndSmartlocker,
		"enforcecomponentsandstoreapps":            AppLockerApplicationControlType_EnforceComponentsAndStoreApps,
		"enforcecomponentsstoreappsandsmartlocker": AppLockerApplicationControlType_EnforceComponentsStoreAppsAndSmartlocker,
		"notconfigured":                            AppLockerApplicationControlType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLockerApplicationControlType(input)
	return &out, nil
}
