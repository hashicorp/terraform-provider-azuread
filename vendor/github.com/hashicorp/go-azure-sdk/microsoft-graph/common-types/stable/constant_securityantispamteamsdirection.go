package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAntispamTeamsDirection string

const (
	SecurityAntispamTeamsDirection_Inbound  SecurityAntispamTeamsDirection = "inbound"
	SecurityAntispamTeamsDirection_Intraorg SecurityAntispamTeamsDirection = "intraorg"
	SecurityAntispamTeamsDirection_Outbound SecurityAntispamTeamsDirection = "outbound"
	SecurityAntispamTeamsDirection_Unknown  SecurityAntispamTeamsDirection = "unknown"
)

func PossibleValuesForSecurityAntispamTeamsDirection() []string {
	return []string{
		string(SecurityAntispamTeamsDirection_Inbound),
		string(SecurityAntispamTeamsDirection_Intraorg),
		string(SecurityAntispamTeamsDirection_Outbound),
		string(SecurityAntispamTeamsDirection_Unknown),
	}
}

func (s *SecurityAntispamTeamsDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAntispamTeamsDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAntispamTeamsDirection(input string) (*SecurityAntispamTeamsDirection, error) {
	vals := map[string]SecurityAntispamTeamsDirection{
		"inbound":  SecurityAntispamTeamsDirection_Inbound,
		"intraorg": SecurityAntispamTeamsDirection_Intraorg,
		"outbound": SecurityAntispamTeamsDirection_Outbound,
		"unknown":  SecurityAntispamTeamsDirection_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAntispamTeamsDirection(input)
	return &out, nil
}
