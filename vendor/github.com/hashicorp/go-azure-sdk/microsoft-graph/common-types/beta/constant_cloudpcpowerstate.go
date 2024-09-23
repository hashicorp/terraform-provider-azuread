package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPowerState string

const (
	CloudPCPowerState_PoweredOff CloudPCPowerState = "poweredOff"
	CloudPCPowerState_Running    CloudPCPowerState = "running"
)

func PossibleValuesForCloudPCPowerState() []string {
	return []string{
		string(CloudPCPowerState_PoweredOff),
		string(CloudPCPowerState_Running),
	}
}

func (s *CloudPCPowerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPowerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPowerState(input string) (*CloudPCPowerState, error) {
	vals := map[string]CloudPCPowerState{
		"poweredoff": CloudPCPowerState_PoweredOff,
		"running":    CloudPCPowerState_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPowerState(input)
	return &out, nil
}
