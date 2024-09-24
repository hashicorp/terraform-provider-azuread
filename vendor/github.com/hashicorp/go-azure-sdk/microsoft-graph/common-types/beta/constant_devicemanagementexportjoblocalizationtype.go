package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJobLocalizationType string

const (
	DeviceManagementExportJobLocalizationType_LocalizedValuesAsAdditionalColumn DeviceManagementExportJobLocalizationType = "localizedValuesAsAdditionalColumn"
	DeviceManagementExportJobLocalizationType_ReplaceLocalizableValues          DeviceManagementExportJobLocalizationType = "replaceLocalizableValues"
)

func PossibleValuesForDeviceManagementExportJobLocalizationType() []string {
	return []string{
		string(DeviceManagementExportJobLocalizationType_LocalizedValuesAsAdditionalColumn),
		string(DeviceManagementExportJobLocalizationType_ReplaceLocalizableValues),
	}
}

func (s *DeviceManagementExportJobLocalizationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExportJobLocalizationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExportJobLocalizationType(input string) (*DeviceManagementExportJobLocalizationType, error) {
	vals := map[string]DeviceManagementExportJobLocalizationType{
		"localizedvaluesasadditionalcolumn": DeviceManagementExportJobLocalizationType_LocalizedValuesAsAdditionalColumn,
		"replacelocalizablevalues":          DeviceManagementExportJobLocalizationType_ReplaceLocalizableValues,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExportJobLocalizationType(input)
	return &out, nil
}
