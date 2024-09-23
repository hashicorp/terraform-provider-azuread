package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeActionStatus string

const (
	DeviceScopeActionStatus_Failed    DeviceScopeActionStatus = "failed"
	DeviceScopeActionStatus_Succeeded DeviceScopeActionStatus = "succeeded"
)

func PossibleValuesForDeviceScopeActionStatus() []string {
	return []string{
		string(DeviceScopeActionStatus_Failed),
		string(DeviceScopeActionStatus_Succeeded),
	}
}

func (s *DeviceScopeActionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceScopeActionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceScopeActionStatus(input string) (*DeviceScopeActionStatus, error) {
	vals := map[string]DeviceScopeActionStatus{
		"failed":    DeviceScopeActionStatus_Failed,
		"succeeded": DeviceScopeActionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceScopeActionStatus(input)
	return &out, nil
}
