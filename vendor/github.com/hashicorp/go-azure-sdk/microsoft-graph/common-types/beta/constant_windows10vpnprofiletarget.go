package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnProfileTarget string

const (
	Windows10VpnProfileTarget_AutoPilotDevice Windows10VpnProfileTarget = "autoPilotDevice"
	Windows10VpnProfileTarget_Device          Windows10VpnProfileTarget = "device"
	Windows10VpnProfileTarget_User            Windows10VpnProfileTarget = "user"
)

func PossibleValuesForWindows10VpnProfileTarget() []string {
	return []string{
		string(Windows10VpnProfileTarget_AutoPilotDevice),
		string(Windows10VpnProfileTarget_Device),
		string(Windows10VpnProfileTarget_User),
	}
}

func (s *Windows10VpnProfileTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10VpnProfileTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10VpnProfileTarget(input string) (*Windows10VpnProfileTarget, error) {
	vals := map[string]Windows10VpnProfileTarget{
		"autopilotdevice": Windows10VpnProfileTarget_AutoPilotDevice,
		"device":          Windows10VpnProfileTarget_Device,
		"user":            Windows10VpnProfileTarget_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnProfileTarget(input)
	return &out, nil
}
