package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkType string

const (
	NetworkType_Extranet             NetworkType = "extranet"
	NetworkType_Intranet             NetworkType = "intranet"
	NetworkType_NamedNetwork         NetworkType = "namedNetwork"
	NetworkType_Trusted              NetworkType = "trusted"
	NetworkType_TrustedNamedLocation NetworkType = "trustedNamedLocation"
)

func PossibleValuesForNetworkType() []string {
	return []string{
		string(NetworkType_Extranet),
		string(NetworkType_Intranet),
		string(NetworkType_NamedNetwork),
		string(NetworkType_Trusted),
		string(NetworkType_TrustedNamedLocation),
	}
}

func (s *NetworkType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkType(input string) (*NetworkType, error) {
	vals := map[string]NetworkType{
		"extranet":             NetworkType_Extranet,
		"intranet":             NetworkType_Intranet,
		"namednetwork":         NetworkType_NamedNetwork,
		"trusted":              NetworkType_Trusted,
		"trustednamedlocation": NetworkType_TrustedNamedLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkType(input)
	return &out, nil
}
