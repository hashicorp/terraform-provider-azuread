package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkDefaultAppPermissionPolicyType string

const (
	AndroidForWorkDefaultAppPermissionPolicyType_AutoDeny      AndroidForWorkDefaultAppPermissionPolicyType = "autoDeny"
	AndroidForWorkDefaultAppPermissionPolicyType_AutoGrant     AndroidForWorkDefaultAppPermissionPolicyType = "autoGrant"
	AndroidForWorkDefaultAppPermissionPolicyType_DeviceDefault AndroidForWorkDefaultAppPermissionPolicyType = "deviceDefault"
	AndroidForWorkDefaultAppPermissionPolicyType_Prompt        AndroidForWorkDefaultAppPermissionPolicyType = "prompt"
)

func PossibleValuesForAndroidForWorkDefaultAppPermissionPolicyType() []string {
	return []string{
		string(AndroidForWorkDefaultAppPermissionPolicyType_AutoDeny),
		string(AndroidForWorkDefaultAppPermissionPolicyType_AutoGrant),
		string(AndroidForWorkDefaultAppPermissionPolicyType_DeviceDefault),
		string(AndroidForWorkDefaultAppPermissionPolicyType_Prompt),
	}
}

func (s *AndroidForWorkDefaultAppPermissionPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkDefaultAppPermissionPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkDefaultAppPermissionPolicyType(input string) (*AndroidForWorkDefaultAppPermissionPolicyType, error) {
	vals := map[string]AndroidForWorkDefaultAppPermissionPolicyType{
		"autodeny":      AndroidForWorkDefaultAppPermissionPolicyType_AutoDeny,
		"autogrant":     AndroidForWorkDefaultAppPermissionPolicyType_AutoGrant,
		"devicedefault": AndroidForWorkDefaultAppPermissionPolicyType_DeviceDefault,
		"prompt":        AndroidForWorkDefaultAppPermissionPolicyType_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkDefaultAppPermissionPolicyType(input)
	return &out, nil
}
