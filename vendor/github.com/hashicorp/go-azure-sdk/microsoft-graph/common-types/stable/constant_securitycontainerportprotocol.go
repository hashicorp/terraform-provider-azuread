package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerPortProtocol string

const (
	SecurityContainerPortProtocol_Sctp SecurityContainerPortProtocol = "sctp"
	SecurityContainerPortProtocol_Tcp  SecurityContainerPortProtocol = "tcp"
	SecurityContainerPortProtocol_Udp  SecurityContainerPortProtocol = "udp"
)

func PossibleValuesForSecurityContainerPortProtocol() []string {
	return []string{
		string(SecurityContainerPortProtocol_Sctp),
		string(SecurityContainerPortProtocol_Tcp),
		string(SecurityContainerPortProtocol_Udp),
	}
}

func (s *SecurityContainerPortProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerPortProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerPortProtocol(input string) (*SecurityContainerPortProtocol, error) {
	vals := map[string]SecurityContainerPortProtocol{
		"sctp": SecurityContainerPortProtocol_Sctp,
		"tcp":  SecurityContainerPortProtocol_Tcp,
		"udp":  SecurityContainerPortProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerPortProtocol(input)
	return &out, nil
}
