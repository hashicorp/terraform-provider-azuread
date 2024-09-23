package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidTargetedPlatforms string

const (
	AndroidTargetedPlatforms_AndroidDeviceAdministrator AndroidTargetedPlatforms = "androidDeviceAdministrator"
	AndroidTargetedPlatforms_AndroidOpenSourceProject   AndroidTargetedPlatforms = "androidOpenSourceProject"
)

func PossibleValuesForAndroidTargetedPlatforms() []string {
	return []string{
		string(AndroidTargetedPlatforms_AndroidDeviceAdministrator),
		string(AndroidTargetedPlatforms_AndroidOpenSourceProject),
	}
}

func (s *AndroidTargetedPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidTargetedPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidTargetedPlatforms(input string) (*AndroidTargetedPlatforms, error) {
	vals := map[string]AndroidTargetedPlatforms{
		"androiddeviceadministrator": AndroidTargetedPlatforms_AndroidDeviceAdministrator,
		"androidopensourceproject":   AndroidTargetedPlatforms_AndroidOpenSourceProject,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidTargetedPlatforms(input)
	return &out, nil
}
