package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceAccessProfileIntent string

const (
	DeviceManagementResourceAccessProfileIntent_Apply  DeviceManagementResourceAccessProfileIntent = "apply"
	DeviceManagementResourceAccessProfileIntent_Remove DeviceManagementResourceAccessProfileIntent = "remove"
)

func PossibleValuesForDeviceManagementResourceAccessProfileIntent() []string {
	return []string{
		string(DeviceManagementResourceAccessProfileIntent_Apply),
		string(DeviceManagementResourceAccessProfileIntent_Remove),
	}
}

func (s *DeviceManagementResourceAccessProfileIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementResourceAccessProfileIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementResourceAccessProfileIntent(input string) (*DeviceManagementResourceAccessProfileIntent, error) {
	vals := map[string]DeviceManagementResourceAccessProfileIntent{
		"apply":  DeviceManagementResourceAccessProfileIntent_Apply,
		"remove": DeviceManagementResourceAccessProfileIntent_Remove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementResourceAccessProfileIntent(input)
	return &out, nil
}
