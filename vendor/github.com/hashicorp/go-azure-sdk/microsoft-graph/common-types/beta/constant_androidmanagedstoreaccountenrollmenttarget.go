package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnrollmentTarget string

const (
	AndroidManagedStoreAccountEnrollmentTarget_All                              AndroidManagedStoreAccountEnrollmentTarget = "all"
	AndroidManagedStoreAccountEnrollmentTarget_None                             AndroidManagedStoreAccountEnrollmentTarget = "none"
	AndroidManagedStoreAccountEnrollmentTarget_Targeted                         AndroidManagedStoreAccountEnrollmentTarget = "targeted"
	AndroidManagedStoreAccountEnrollmentTarget_TargetedAsEnrollmentRestrictions AndroidManagedStoreAccountEnrollmentTarget = "targetedAsEnrollmentRestrictions"
)

func PossibleValuesForAndroidManagedStoreAccountEnrollmentTarget() []string {
	return []string{
		string(AndroidManagedStoreAccountEnrollmentTarget_All),
		string(AndroidManagedStoreAccountEnrollmentTarget_None),
		string(AndroidManagedStoreAccountEnrollmentTarget_Targeted),
		string(AndroidManagedStoreAccountEnrollmentTarget_TargetedAsEnrollmentRestrictions),
	}
}

func (s *AndroidManagedStoreAccountEnrollmentTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAccountEnrollmentTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAccountEnrollmentTarget(input string) (*AndroidManagedStoreAccountEnrollmentTarget, error) {
	vals := map[string]AndroidManagedStoreAccountEnrollmentTarget{
		"all":                              AndroidManagedStoreAccountEnrollmentTarget_All,
		"none":                             AndroidManagedStoreAccountEnrollmentTarget_None,
		"targeted":                         AndroidManagedStoreAccountEnrollmentTarget_Targeted,
		"targetedasenrollmentrestrictions": AndroidManagedStoreAccountEnrollmentTarget_TargetedAsEnrollmentRestrictions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountEnrollmentTarget(input)
	return &out, nil
}
