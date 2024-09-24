package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppSafetyNetAppsVerificationType string

const (
	AndroidManagedAppSafetyNetAppsVerificationType_Enabled AndroidManagedAppSafetyNetAppsVerificationType = "enabled"
	AndroidManagedAppSafetyNetAppsVerificationType_None    AndroidManagedAppSafetyNetAppsVerificationType = "none"
)

func PossibleValuesForAndroidManagedAppSafetyNetAppsVerificationType() []string {
	return []string{
		string(AndroidManagedAppSafetyNetAppsVerificationType_Enabled),
		string(AndroidManagedAppSafetyNetAppsVerificationType_None),
	}
}

func (s *AndroidManagedAppSafetyNetAppsVerificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppSafetyNetAppsVerificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppSafetyNetAppsVerificationType(input string) (*AndroidManagedAppSafetyNetAppsVerificationType, error) {
	vals := map[string]AndroidManagedAppSafetyNetAppsVerificationType{
		"enabled": AndroidManagedAppSafetyNetAppsVerificationType_Enabled,
		"none":    AndroidManagedAppSafetyNetAppsVerificationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppSafetyNetAppsVerificationType(input)
	return &out, nil
}
