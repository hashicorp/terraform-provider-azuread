package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDefaultAppPermissionPolicyType string

const (
	AndroidDeviceOwnerDefaultAppPermissionPolicyType_AutoDeny      AndroidDeviceOwnerDefaultAppPermissionPolicyType = "autoDeny"
	AndroidDeviceOwnerDefaultAppPermissionPolicyType_AutoGrant     AndroidDeviceOwnerDefaultAppPermissionPolicyType = "autoGrant"
	AndroidDeviceOwnerDefaultAppPermissionPolicyType_DeviceDefault AndroidDeviceOwnerDefaultAppPermissionPolicyType = "deviceDefault"
	AndroidDeviceOwnerDefaultAppPermissionPolicyType_Prompt        AndroidDeviceOwnerDefaultAppPermissionPolicyType = "prompt"
)

func PossibleValuesForAndroidDeviceOwnerDefaultAppPermissionPolicyType() []string {
	return []string{
		string(AndroidDeviceOwnerDefaultAppPermissionPolicyType_AutoDeny),
		string(AndroidDeviceOwnerDefaultAppPermissionPolicyType_AutoGrant),
		string(AndroidDeviceOwnerDefaultAppPermissionPolicyType_DeviceDefault),
		string(AndroidDeviceOwnerDefaultAppPermissionPolicyType_Prompt),
	}
}

func (s *AndroidDeviceOwnerDefaultAppPermissionPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerDefaultAppPermissionPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerDefaultAppPermissionPolicyType(input string) (*AndroidDeviceOwnerDefaultAppPermissionPolicyType, error) {
	vals := map[string]AndroidDeviceOwnerDefaultAppPermissionPolicyType{
		"autodeny":      AndroidDeviceOwnerDefaultAppPermissionPolicyType_AutoDeny,
		"autogrant":     AndroidDeviceOwnerDefaultAppPermissionPolicyType_AutoGrant,
		"devicedefault": AndroidDeviceOwnerDefaultAppPermissionPolicyType_DeviceDefault,
		"prompt":        AndroidDeviceOwnerDefaultAppPermissionPolicyType_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerDefaultAppPermissionPolicyType(input)
	return &out, nil
}
