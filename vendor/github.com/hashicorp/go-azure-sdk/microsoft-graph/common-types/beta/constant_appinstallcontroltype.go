package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppInstallControlType string

const (
	AppInstallControlType_Anywhere        AppInstallControlType = "anywhere"
	AppInstallControlType_NotConfigured   AppInstallControlType = "notConfigured"
	AppInstallControlType_PreferStore     AppInstallControlType = "preferStore"
	AppInstallControlType_Recommendations AppInstallControlType = "recommendations"
	AppInstallControlType_StoreOnly       AppInstallControlType = "storeOnly"
)

func PossibleValuesForAppInstallControlType() []string {
	return []string{
		string(AppInstallControlType_Anywhere),
		string(AppInstallControlType_NotConfigured),
		string(AppInstallControlType_PreferStore),
		string(AppInstallControlType_Recommendations),
		string(AppInstallControlType_StoreOnly),
	}
}

func (s *AppInstallControlType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppInstallControlType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppInstallControlType(input string) (*AppInstallControlType, error) {
	vals := map[string]AppInstallControlType{
		"anywhere":        AppInstallControlType_Anywhere,
		"notconfigured":   AppInstallControlType_NotConfigured,
		"preferstore":     AppInstallControlType_PreferStore,
		"recommendations": AppInstallControlType_Recommendations,
		"storeonly":       AppInstallControlType_StoreOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppInstallControlType(input)
	return &out, nil
}
