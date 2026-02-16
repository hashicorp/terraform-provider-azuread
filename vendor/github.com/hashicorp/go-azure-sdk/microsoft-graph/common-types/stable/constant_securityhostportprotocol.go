package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPortProtocol string

const (
	SecurityHostPortProtocol_Tcp SecurityHostPortProtocol = "tcp"
	SecurityHostPortProtocol_Udp SecurityHostPortProtocol = "udp"
)

func PossibleValuesForSecurityHostPortProtocol() []string {
	return []string{
		string(SecurityHostPortProtocol_Tcp),
		string(SecurityHostPortProtocol_Udp),
	}
}

func (s *SecurityHostPortProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHostPortProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHostPortProtocol(input string) (*SecurityHostPortProtocol, error) {
	vals := map[string]SecurityHostPortProtocol{
		"tcp": SecurityHostPortProtocol_Tcp,
		"udp": SecurityHostPortProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostPortProtocol(input)
	return &out, nil
}
