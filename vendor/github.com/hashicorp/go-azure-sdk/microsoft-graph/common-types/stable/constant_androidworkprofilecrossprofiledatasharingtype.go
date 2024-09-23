package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCrossProfileDataSharingType string

const (
	AndroidWorkProfileCrossProfileDataSharingType_AllowPersonalToWork AndroidWorkProfileCrossProfileDataSharingType = "allowPersonalToWork"
	AndroidWorkProfileCrossProfileDataSharingType_DeviceDefault       AndroidWorkProfileCrossProfileDataSharingType = "deviceDefault"
	AndroidWorkProfileCrossProfileDataSharingType_NoRestrictions      AndroidWorkProfileCrossProfileDataSharingType = "noRestrictions"
	AndroidWorkProfileCrossProfileDataSharingType_PreventAny          AndroidWorkProfileCrossProfileDataSharingType = "preventAny"
)

func PossibleValuesForAndroidWorkProfileCrossProfileDataSharingType() []string {
	return []string{
		string(AndroidWorkProfileCrossProfileDataSharingType_AllowPersonalToWork),
		string(AndroidWorkProfileCrossProfileDataSharingType_DeviceDefault),
		string(AndroidWorkProfileCrossProfileDataSharingType_NoRestrictions),
		string(AndroidWorkProfileCrossProfileDataSharingType_PreventAny),
	}
}

func (s *AndroidWorkProfileCrossProfileDataSharingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCrossProfileDataSharingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCrossProfileDataSharingType(input string) (*AndroidWorkProfileCrossProfileDataSharingType, error) {
	vals := map[string]AndroidWorkProfileCrossProfileDataSharingType{
		"allowpersonaltowork": AndroidWorkProfileCrossProfileDataSharingType_AllowPersonalToWork,
		"devicedefault":       AndroidWorkProfileCrossProfileDataSharingType_DeviceDefault,
		"norestrictions":      AndroidWorkProfileCrossProfileDataSharingType_NoRestrictions,
		"preventany":          AndroidWorkProfileCrossProfileDataSharingType_PreventAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCrossProfileDataSharingType(input)
	return &out, nil
}
