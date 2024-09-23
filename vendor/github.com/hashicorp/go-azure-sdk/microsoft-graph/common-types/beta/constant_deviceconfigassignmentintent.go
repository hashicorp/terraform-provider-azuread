package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigAssignmentIntent string

const (
	DeviceConfigAssignmentIntent_Apply  DeviceConfigAssignmentIntent = "apply"
	DeviceConfigAssignmentIntent_Remove DeviceConfigAssignmentIntent = "remove"
)

func PossibleValuesForDeviceConfigAssignmentIntent() []string {
	return []string{
		string(DeviceConfigAssignmentIntent_Apply),
		string(DeviceConfigAssignmentIntent_Remove),
	}
}

func (s *DeviceConfigAssignmentIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigAssignmentIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigAssignmentIntent(input string) (*DeviceConfigAssignmentIntent, error) {
	vals := map[string]DeviceConfigAssignmentIntent{
		"apply":  DeviceConfigAssignmentIntent_Apply,
		"remove": DeviceConfigAssignmentIntent_Remove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigAssignmentIntent(input)
	return &out, nil
}
