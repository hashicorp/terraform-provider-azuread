package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateNetworkProtocol string

const (
	PrivateNetworkProtocol_Tcp PrivateNetworkProtocol = "tcp"
	PrivateNetworkProtocol_Udp PrivateNetworkProtocol = "udp"
)

func PossibleValuesForPrivateNetworkProtocol() []string {
	return []string{
		string(PrivateNetworkProtocol_Tcp),
		string(PrivateNetworkProtocol_Udp),
	}
}

func (s *PrivateNetworkProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateNetworkProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateNetworkProtocol(input string) (*PrivateNetworkProtocol, error) {
	vals := map[string]PrivateNetworkProtocol{
		"tcp": PrivateNetworkProtocol_Tcp,
		"udp": PrivateNetworkProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateNetworkProtocol(input)
	return &out, nil
}
