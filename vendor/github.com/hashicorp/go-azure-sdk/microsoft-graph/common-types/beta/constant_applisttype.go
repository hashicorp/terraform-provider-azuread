package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppListType string

const (
	AppListType_AppsInListCompliant    AppListType = "appsInListCompliant"
	AppListType_AppsNotInListCompliant AppListType = "appsNotInListCompliant"
	AppListType_None                   AppListType = "none"
)

func PossibleValuesForAppListType() []string {
	return []string{
		string(AppListType_AppsInListCompliant),
		string(AppListType_AppsNotInListCompliant),
		string(AppListType_None),
	}
}

func (s *AppListType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppListType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppListType(input string) (*AppListType, error) {
	vals := map[string]AppListType{
		"appsinlistcompliant":    AppListType_AppsInListCompliant,
		"appsnotinlistcompliant": AppListType_AppsNotInListCompliant,
		"none":                   AppListType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppListType(input)
	return &out, nil
}
