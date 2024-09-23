package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProtocolType string

const (
	SecurityProtocolType_Tcp SecurityProtocolType = "tcp"
	SecurityProtocolType_Udp SecurityProtocolType = "udp"
)

func PossibleValuesForSecurityProtocolType() []string {
	return []string{
		string(SecurityProtocolType_Tcp),
		string(SecurityProtocolType_Udp),
	}
}

func (s *SecurityProtocolType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityProtocolType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityProtocolType(input string) (*SecurityProtocolType, error) {
	vals := map[string]SecurityProtocolType{
		"tcp": SecurityProtocolType_Tcp,
		"udp": SecurityProtocolType_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProtocolType(input)
	return &out, nil
}
