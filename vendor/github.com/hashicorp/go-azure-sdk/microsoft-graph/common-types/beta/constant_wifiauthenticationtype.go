package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WifiAuthenticationType string

const (
	WifiAuthenticationType_Guest         WifiAuthenticationType = "guest"
	WifiAuthenticationType_Machine       WifiAuthenticationType = "machine"
	WifiAuthenticationType_MachineOrUser WifiAuthenticationType = "machineOrUser"
	WifiAuthenticationType_None          WifiAuthenticationType = "none"
	WifiAuthenticationType_User          WifiAuthenticationType = "user"
)

func PossibleValuesForWifiAuthenticationType() []string {
	return []string{
		string(WifiAuthenticationType_Guest),
		string(WifiAuthenticationType_Machine),
		string(WifiAuthenticationType_MachineOrUser),
		string(WifiAuthenticationType_None),
		string(WifiAuthenticationType_User),
	}
}

func (s *WifiAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWifiAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWifiAuthenticationType(input string) (*WifiAuthenticationType, error) {
	vals := map[string]WifiAuthenticationType{
		"guest":         WifiAuthenticationType_Guest,
		"machine":       WifiAuthenticationType_Machine,
		"machineoruser": WifiAuthenticationType_MachineOrUser,
		"none":          WifiAuthenticationType_None,
		"user":          WifiAuthenticationType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WifiAuthenticationType(input)
	return &out, nil
}
