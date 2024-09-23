package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnrollmentTarget string

const (
	AndroidForWorkEnrollmentTarget_All                              AndroidForWorkEnrollmentTarget = "all"
	AndroidForWorkEnrollmentTarget_None                             AndroidForWorkEnrollmentTarget = "none"
	AndroidForWorkEnrollmentTarget_Targeted                         AndroidForWorkEnrollmentTarget = "targeted"
	AndroidForWorkEnrollmentTarget_TargetedAsEnrollmentRestrictions AndroidForWorkEnrollmentTarget = "targetedAsEnrollmentRestrictions"
)

func PossibleValuesForAndroidForWorkEnrollmentTarget() []string {
	return []string{
		string(AndroidForWorkEnrollmentTarget_All),
		string(AndroidForWorkEnrollmentTarget_None),
		string(AndroidForWorkEnrollmentTarget_Targeted),
		string(AndroidForWorkEnrollmentTarget_TargetedAsEnrollmentRestrictions),
	}
}

func (s *AndroidForWorkEnrollmentTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEnrollmentTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEnrollmentTarget(input string) (*AndroidForWorkEnrollmentTarget, error) {
	vals := map[string]AndroidForWorkEnrollmentTarget{
		"all":                              AndroidForWorkEnrollmentTarget_All,
		"none":                             AndroidForWorkEnrollmentTarget_None,
		"targeted":                         AndroidForWorkEnrollmentTarget_Targeted,
		"targetedasenrollmentrestrictions": AndroidForWorkEnrollmentTarget_TargetedAsEnrollmentRestrictions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnrollmentTarget(input)
	return &out, nil
}
