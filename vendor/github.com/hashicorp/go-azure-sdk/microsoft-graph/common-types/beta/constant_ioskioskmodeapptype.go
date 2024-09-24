package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosKioskModeAppType string

const (
	IosKioskModeAppType_AppStoreApp   IosKioskModeAppType = "appStoreApp"
	IosKioskModeAppType_BuiltInApp    IosKioskModeAppType = "builtInApp"
	IosKioskModeAppType_ManagedApp    IosKioskModeAppType = "managedApp"
	IosKioskModeAppType_NotConfigured IosKioskModeAppType = "notConfigured"
)

func PossibleValuesForIosKioskModeAppType() []string {
	return []string{
		string(IosKioskModeAppType_AppStoreApp),
		string(IosKioskModeAppType_BuiltInApp),
		string(IosKioskModeAppType_ManagedApp),
		string(IosKioskModeAppType_NotConfigured),
	}
}

func (s *IosKioskModeAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosKioskModeAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosKioskModeAppType(input string) (*IosKioskModeAppType, error) {
	vals := map[string]IosKioskModeAppType{
		"appstoreapp":   IosKioskModeAppType_AppStoreApp,
		"builtinapp":    IosKioskModeAppType_BuiltInApp,
		"managedapp":    IosKioskModeAppType_ManagedApp,
		"notconfigured": IosKioskModeAppType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosKioskModeAppType(input)
	return &out, nil
}
