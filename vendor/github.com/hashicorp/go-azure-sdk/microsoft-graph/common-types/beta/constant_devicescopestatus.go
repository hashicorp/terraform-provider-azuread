package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeStatus string

const (
	DeviceScopeStatus_Completed        DeviceScopeStatus = "completed"
	DeviceScopeStatus_Computing        DeviceScopeStatus = "computing"
	DeviceScopeStatus_InsufficientData DeviceScopeStatus = "insufficientData"
	DeviceScopeStatus_None             DeviceScopeStatus = "none"
)

func PossibleValuesForDeviceScopeStatus() []string {
	return []string{
		string(DeviceScopeStatus_Completed),
		string(DeviceScopeStatus_Computing),
		string(DeviceScopeStatus_InsufficientData),
		string(DeviceScopeStatus_None),
	}
}

func (s *DeviceScopeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceScopeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceScopeStatus(input string) (*DeviceScopeStatus, error) {
	vals := map[string]DeviceScopeStatus{
		"completed":        DeviceScopeStatus_Completed,
		"computing":        DeviceScopeStatus_Computing,
		"insufficientdata": DeviceScopeStatus_InsufficientData,
		"none":             DeviceScopeStatus_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceScopeStatus(input)
	return &out, nil
}
