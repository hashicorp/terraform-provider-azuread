package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAutoUpdateMode string

const (
	AndroidManagedStoreAutoUpdateMode_Default   AndroidManagedStoreAutoUpdateMode = "default"
	AndroidManagedStoreAutoUpdateMode_Postponed AndroidManagedStoreAutoUpdateMode = "postponed"
	AndroidManagedStoreAutoUpdateMode_Priority  AndroidManagedStoreAutoUpdateMode = "priority"
)

func PossibleValuesForAndroidManagedStoreAutoUpdateMode() []string {
	return []string{
		string(AndroidManagedStoreAutoUpdateMode_Default),
		string(AndroidManagedStoreAutoUpdateMode_Postponed),
		string(AndroidManagedStoreAutoUpdateMode_Priority),
	}
}

func (s *AndroidManagedStoreAutoUpdateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAutoUpdateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAutoUpdateMode(input string) (*AndroidManagedStoreAutoUpdateMode, error) {
	vals := map[string]AndroidManagedStoreAutoUpdateMode{
		"default":   AndroidManagedStoreAutoUpdateMode_Default,
		"postponed": AndroidManagedStoreAutoUpdateMode_Postponed,
		"priority":  AndroidManagedStoreAutoUpdateMode_Priority,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAutoUpdateMode(input)
	return &out, nil
}
