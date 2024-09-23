package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSingleSignOnType string

const (
	NetworkSingleSignOnType_Disabled  NetworkSingleSignOnType = "disabled"
	NetworkSingleSignOnType_Postlogon NetworkSingleSignOnType = "postlogon"
	NetworkSingleSignOnType_Prelogon  NetworkSingleSignOnType = "prelogon"
)

func PossibleValuesForNetworkSingleSignOnType() []string {
	return []string{
		string(NetworkSingleSignOnType_Disabled),
		string(NetworkSingleSignOnType_Postlogon),
		string(NetworkSingleSignOnType_Prelogon),
	}
}

func (s *NetworkSingleSignOnType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSingleSignOnType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSingleSignOnType(input string) (*NetworkSingleSignOnType, error) {
	vals := map[string]NetworkSingleSignOnType{
		"disabled":  NetworkSingleSignOnType_Disabled,
		"postlogon": NetworkSingleSignOnType_Postlogon,
		"prelogon":  NetworkSingleSignOnType_Prelogon,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSingleSignOnType(input)
	return &out, nil
}
