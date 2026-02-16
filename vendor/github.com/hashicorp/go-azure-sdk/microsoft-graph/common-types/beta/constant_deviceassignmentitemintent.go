package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAssignmentItemIntent string

const (
	DeviceAssignmentItemIntent_Remove  DeviceAssignmentItemIntent = "remove"
	DeviceAssignmentItemIntent_Restore DeviceAssignmentItemIntent = "restore"
)

func PossibleValuesForDeviceAssignmentItemIntent() []string {
	return []string{
		string(DeviceAssignmentItemIntent_Remove),
		string(DeviceAssignmentItemIntent_Restore),
	}
}

func (s *DeviceAssignmentItemIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAssignmentItemIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAssignmentItemIntent(input string) (*DeviceAssignmentItemIntent, error) {
	vals := map[string]DeviceAssignmentItemIntent{
		"remove":  DeviceAssignmentItemIntent_Remove,
		"restore": DeviceAssignmentItemIntent_Restore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAssignmentItemIntent(input)
	return &out, nil
}
