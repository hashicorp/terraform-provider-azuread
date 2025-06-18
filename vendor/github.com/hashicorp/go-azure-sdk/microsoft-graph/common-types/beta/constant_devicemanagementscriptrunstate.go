package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptRunState string

const (
	DeviceManagementScriptRunState_Fail          DeviceManagementScriptRunState = "fail"
	DeviceManagementScriptRunState_NotApplicable DeviceManagementScriptRunState = "notApplicable"
	DeviceManagementScriptRunState_Pending       DeviceManagementScriptRunState = "pending"
	DeviceManagementScriptRunState_ScriptError   DeviceManagementScriptRunState = "scriptError"
	DeviceManagementScriptRunState_Success       DeviceManagementScriptRunState = "success"
	DeviceManagementScriptRunState_Unknown       DeviceManagementScriptRunState = "unknown"
)

func PossibleValuesForDeviceManagementScriptRunState() []string {
	return []string{
		string(DeviceManagementScriptRunState_Fail),
		string(DeviceManagementScriptRunState_NotApplicable),
		string(DeviceManagementScriptRunState_Pending),
		string(DeviceManagementScriptRunState_ScriptError),
		string(DeviceManagementScriptRunState_Success),
		string(DeviceManagementScriptRunState_Unknown),
	}
}

func (s *DeviceManagementScriptRunState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementScriptRunState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementScriptRunState(input string) (*DeviceManagementScriptRunState, error) {
	vals := map[string]DeviceManagementScriptRunState{
		"fail":          DeviceManagementScriptRunState_Fail,
		"notapplicable": DeviceManagementScriptRunState_NotApplicable,
		"pending":       DeviceManagementScriptRunState_Pending,
		"scripterror":   DeviceManagementScriptRunState_ScriptError,
		"success":       DeviceManagementScriptRunState_Success,
		"unknown":       DeviceManagementScriptRunState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptRunState(input)
	return &out, nil
}
