package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSubscriptions string

const (
	DeviceManagementSubscriptions_Intune        DeviceManagementSubscriptions = "intune"
	DeviceManagementSubscriptions_IntuneEDU     DeviceManagementSubscriptions = "intune_EDU"
	DeviceManagementSubscriptions_IntunePremium DeviceManagementSubscriptions = "intunePremium"
	DeviceManagementSubscriptions_IntuneSMB     DeviceManagementSubscriptions = "intune_SMB"
	DeviceManagementSubscriptions_None          DeviceManagementSubscriptions = "none"
	DeviceManagementSubscriptions_Office365     DeviceManagementSubscriptions = "office365"
)

func PossibleValuesForDeviceManagementSubscriptions() []string {
	return []string{
		string(DeviceManagementSubscriptions_Intune),
		string(DeviceManagementSubscriptions_IntuneEDU),
		string(DeviceManagementSubscriptions_IntunePremium),
		string(DeviceManagementSubscriptions_IntuneSMB),
		string(DeviceManagementSubscriptions_None),
		string(DeviceManagementSubscriptions_Office365),
	}
}

func (s *DeviceManagementSubscriptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementSubscriptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementSubscriptions(input string) (*DeviceManagementSubscriptions, error) {
	vals := map[string]DeviceManagementSubscriptions{
		"intune":        DeviceManagementSubscriptions_Intune,
		"intune_edu":    DeviceManagementSubscriptions_IntuneEDU,
		"intunepremium": DeviceManagementSubscriptions_IntunePremium,
		"intune_smb":    DeviceManagementSubscriptions_IntuneSMB,
		"none":          DeviceManagementSubscriptions_None,
		"office365":     DeviceManagementSubscriptions_Office365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSubscriptions(input)
	return &out, nil
}
