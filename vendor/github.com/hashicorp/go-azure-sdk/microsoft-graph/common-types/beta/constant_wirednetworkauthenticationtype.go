package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WiredNetworkAuthenticationType string

const (
	WiredNetworkAuthenticationType_Guest         WiredNetworkAuthenticationType = "guest"
	WiredNetworkAuthenticationType_Machine       WiredNetworkAuthenticationType = "machine"
	WiredNetworkAuthenticationType_MachineOrUser WiredNetworkAuthenticationType = "machineOrUser"
	WiredNetworkAuthenticationType_None          WiredNetworkAuthenticationType = "none"
	WiredNetworkAuthenticationType_User          WiredNetworkAuthenticationType = "user"
)

func PossibleValuesForWiredNetworkAuthenticationType() []string {
	return []string{
		string(WiredNetworkAuthenticationType_Guest),
		string(WiredNetworkAuthenticationType_Machine),
		string(WiredNetworkAuthenticationType_MachineOrUser),
		string(WiredNetworkAuthenticationType_None),
		string(WiredNetworkAuthenticationType_User),
	}
}

func (s *WiredNetworkAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWiredNetworkAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWiredNetworkAuthenticationType(input string) (*WiredNetworkAuthenticationType, error) {
	vals := map[string]WiredNetworkAuthenticationType{
		"guest":         WiredNetworkAuthenticationType_Guest,
		"machine":       WiredNetworkAuthenticationType_Machine,
		"machineoruser": WiredNetworkAuthenticationType_MachineOrUser,
		"none":          WiredNetworkAuthenticationType_None,
		"user":          WiredNetworkAuthenticationType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WiredNetworkAuthenticationType(input)
	return &out, nil
}
