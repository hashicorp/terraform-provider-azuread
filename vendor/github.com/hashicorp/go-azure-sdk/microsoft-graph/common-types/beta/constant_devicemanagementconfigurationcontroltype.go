package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationControlType string

const (
	DeviceManagementConfigurationControlType_ContextPane     DeviceManagementConfigurationControlType = "contextPane"
	DeviceManagementConfigurationControlType_Default         DeviceManagementConfigurationControlType = "default"
	DeviceManagementConfigurationControlType_Dropdown        DeviceManagementConfigurationControlType = "dropdown"
	DeviceManagementConfigurationControlType_LargeTextBox    DeviceManagementConfigurationControlType = "largeTextBox"
	DeviceManagementConfigurationControlType_MultiheaderGrid DeviceManagementConfigurationControlType = "multiheaderGrid"
	DeviceManagementConfigurationControlType_SmallTextBox    DeviceManagementConfigurationControlType = "smallTextBox"
	DeviceManagementConfigurationControlType_Toggle          DeviceManagementConfigurationControlType = "toggle"
)

func PossibleValuesForDeviceManagementConfigurationControlType() []string {
	return []string{
		string(DeviceManagementConfigurationControlType_ContextPane),
		string(DeviceManagementConfigurationControlType_Default),
		string(DeviceManagementConfigurationControlType_Dropdown),
		string(DeviceManagementConfigurationControlType_LargeTextBox),
		string(DeviceManagementConfigurationControlType_MultiheaderGrid),
		string(DeviceManagementConfigurationControlType_SmallTextBox),
		string(DeviceManagementConfigurationControlType_Toggle),
	}
}

func (s *DeviceManagementConfigurationControlType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationControlType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationControlType(input string) (*DeviceManagementConfigurationControlType, error) {
	vals := map[string]DeviceManagementConfigurationControlType{
		"contextpane":     DeviceManagementConfigurationControlType_ContextPane,
		"default":         DeviceManagementConfigurationControlType_Default,
		"dropdown":        DeviceManagementConfigurationControlType_Dropdown,
		"largetextbox":    DeviceManagementConfigurationControlType_LargeTextBox,
		"multiheadergrid": DeviceManagementConfigurationControlType_MultiheaderGrid,
		"smalltextbox":    DeviceManagementConfigurationControlType_SmallTextBox,
		"toggle":          DeviceManagementConfigurationControlType_Toggle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationControlType(input)
	return &out, nil
}
