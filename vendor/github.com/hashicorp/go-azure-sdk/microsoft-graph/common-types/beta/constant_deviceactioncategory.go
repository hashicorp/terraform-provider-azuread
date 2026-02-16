package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceActionCategory string

const (
	DeviceActionCategory_Bulk   DeviceActionCategory = "bulk"
	DeviceActionCategory_Single DeviceActionCategory = "single"
)

func PossibleValuesForDeviceActionCategory() []string {
	return []string{
		string(DeviceActionCategory_Bulk),
		string(DeviceActionCategory_Single),
	}
}

func (s *DeviceActionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceActionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceActionCategory(input string) (*DeviceActionCategory, error) {
	vals := map[string]DeviceActionCategory{
		"bulk":   DeviceActionCategory_Bulk,
		"single": DeviceActionCategory_Single,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceActionCategory(input)
	return &out, nil
}
