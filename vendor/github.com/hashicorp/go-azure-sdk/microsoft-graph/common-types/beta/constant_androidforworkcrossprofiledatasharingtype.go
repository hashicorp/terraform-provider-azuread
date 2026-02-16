package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCrossProfileDataSharingType string

const (
	AndroidForWorkCrossProfileDataSharingType_AllowPersonalToWork AndroidForWorkCrossProfileDataSharingType = "allowPersonalToWork"
	AndroidForWorkCrossProfileDataSharingType_DeviceDefault       AndroidForWorkCrossProfileDataSharingType = "deviceDefault"
	AndroidForWorkCrossProfileDataSharingType_NoRestrictions      AndroidForWorkCrossProfileDataSharingType = "noRestrictions"
	AndroidForWorkCrossProfileDataSharingType_PreventAny          AndroidForWorkCrossProfileDataSharingType = "preventAny"
)

func PossibleValuesForAndroidForWorkCrossProfileDataSharingType() []string {
	return []string{
		string(AndroidForWorkCrossProfileDataSharingType_AllowPersonalToWork),
		string(AndroidForWorkCrossProfileDataSharingType_DeviceDefault),
		string(AndroidForWorkCrossProfileDataSharingType_NoRestrictions),
		string(AndroidForWorkCrossProfileDataSharingType_PreventAny),
	}
}

func (s *AndroidForWorkCrossProfileDataSharingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCrossProfileDataSharingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCrossProfileDataSharingType(input string) (*AndroidForWorkCrossProfileDataSharingType, error) {
	vals := map[string]AndroidForWorkCrossProfileDataSharingType{
		"allowpersonaltowork": AndroidForWorkCrossProfileDataSharingType_AllowPersonalToWork,
		"devicedefault":       AndroidForWorkCrossProfileDataSharingType_DeviceDefault,
		"norestrictions":      AndroidForWorkCrossProfileDataSharingType_NoRestrictions,
		"preventany":          AndroidForWorkCrossProfileDataSharingType_PreventAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCrossProfileDataSharingType(input)
	return &out, nil
}
