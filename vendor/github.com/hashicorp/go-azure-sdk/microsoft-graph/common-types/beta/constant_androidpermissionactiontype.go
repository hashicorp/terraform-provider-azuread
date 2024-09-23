package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPermissionActionType string

const (
	AndroidPermissionActionType_AutoDeny  AndroidPermissionActionType = "autoDeny"
	AndroidPermissionActionType_AutoGrant AndroidPermissionActionType = "autoGrant"
	AndroidPermissionActionType_Prompt    AndroidPermissionActionType = "prompt"
)

func PossibleValuesForAndroidPermissionActionType() []string {
	return []string{
		string(AndroidPermissionActionType_AutoDeny),
		string(AndroidPermissionActionType_AutoGrant),
		string(AndroidPermissionActionType_Prompt),
	}
}

func (s *AndroidPermissionActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidPermissionActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidPermissionActionType(input string) (*AndroidPermissionActionType, error) {
	vals := map[string]AndroidPermissionActionType{
		"autodeny":  AndroidPermissionActionType_AutoDeny,
		"autogrant": AndroidPermissionActionType_AutoGrant,
		"prompt":    AndroidPermissionActionType_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPermissionActionType(input)
	return &out, nil
}
