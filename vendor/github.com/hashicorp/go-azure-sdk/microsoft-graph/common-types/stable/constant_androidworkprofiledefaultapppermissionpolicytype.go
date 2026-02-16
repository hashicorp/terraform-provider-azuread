package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileDefaultAppPermissionPolicyType string

const (
	AndroidWorkProfileDefaultAppPermissionPolicyType_AutoDeny      AndroidWorkProfileDefaultAppPermissionPolicyType = "autoDeny"
	AndroidWorkProfileDefaultAppPermissionPolicyType_AutoGrant     AndroidWorkProfileDefaultAppPermissionPolicyType = "autoGrant"
	AndroidWorkProfileDefaultAppPermissionPolicyType_DeviceDefault AndroidWorkProfileDefaultAppPermissionPolicyType = "deviceDefault"
	AndroidWorkProfileDefaultAppPermissionPolicyType_Prompt        AndroidWorkProfileDefaultAppPermissionPolicyType = "prompt"
)

func PossibleValuesForAndroidWorkProfileDefaultAppPermissionPolicyType() []string {
	return []string{
		string(AndroidWorkProfileDefaultAppPermissionPolicyType_AutoDeny),
		string(AndroidWorkProfileDefaultAppPermissionPolicyType_AutoGrant),
		string(AndroidWorkProfileDefaultAppPermissionPolicyType_DeviceDefault),
		string(AndroidWorkProfileDefaultAppPermissionPolicyType_Prompt),
	}
}

func (s *AndroidWorkProfileDefaultAppPermissionPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileDefaultAppPermissionPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileDefaultAppPermissionPolicyType(input string) (*AndroidWorkProfileDefaultAppPermissionPolicyType, error) {
	vals := map[string]AndroidWorkProfileDefaultAppPermissionPolicyType{
		"autodeny":      AndroidWorkProfileDefaultAppPermissionPolicyType_AutoDeny,
		"autogrant":     AndroidWorkProfileDefaultAppPermissionPolicyType_AutoGrant,
		"devicedefault": AndroidWorkProfileDefaultAppPermissionPolicyType_DeviceDefault,
		"prompt":        AndroidWorkProfileDefaultAppPermissionPolicyType_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileDefaultAppPermissionPolicyType(input)
	return &out, nil
}
