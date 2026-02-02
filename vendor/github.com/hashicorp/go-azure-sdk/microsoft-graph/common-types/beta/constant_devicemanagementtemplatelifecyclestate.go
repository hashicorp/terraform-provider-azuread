package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplateLifecycleState string

const (
	DeviceManagementTemplateLifecycleState_Active     DeviceManagementTemplateLifecycleState = "active"
	DeviceManagementTemplateLifecycleState_Deprecated DeviceManagementTemplateLifecycleState = "deprecated"
	DeviceManagementTemplateLifecycleState_Draft      DeviceManagementTemplateLifecycleState = "draft"
	DeviceManagementTemplateLifecycleState_Invalid    DeviceManagementTemplateLifecycleState = "invalid"
	DeviceManagementTemplateLifecycleState_Retired    DeviceManagementTemplateLifecycleState = "retired"
	DeviceManagementTemplateLifecycleState_Superseded DeviceManagementTemplateLifecycleState = "superseded"
)

func PossibleValuesForDeviceManagementTemplateLifecycleState() []string {
	return []string{
		string(DeviceManagementTemplateLifecycleState_Active),
		string(DeviceManagementTemplateLifecycleState_Deprecated),
		string(DeviceManagementTemplateLifecycleState_Draft),
		string(DeviceManagementTemplateLifecycleState_Invalid),
		string(DeviceManagementTemplateLifecycleState_Retired),
		string(DeviceManagementTemplateLifecycleState_Superseded),
	}
}

func (s *DeviceManagementTemplateLifecycleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementTemplateLifecycleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementTemplateLifecycleState(input string) (*DeviceManagementTemplateLifecycleState, error) {
	vals := map[string]DeviceManagementTemplateLifecycleState{
		"active":     DeviceManagementTemplateLifecycleState_Active,
		"deprecated": DeviceManagementTemplateLifecycleState_Deprecated,
		"draft":      DeviceManagementTemplateLifecycleState_Draft,
		"invalid":    DeviceManagementTemplateLifecycleState_Invalid,
		"retired":    DeviceManagementTemplateLifecycleState_Retired,
		"superseded": DeviceManagementTemplateLifecycleState_Superseded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplateLifecycleState(input)
	return &out, nil
}
