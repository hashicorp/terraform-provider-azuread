package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAntispamDirectionality string

const (
	SecurityAntispamDirectionality_Inbound  SecurityAntispamDirectionality = "inbound"
	SecurityAntispamDirectionality_IntraOrg SecurityAntispamDirectionality = "intraOrg"
	SecurityAntispamDirectionality_Outbound SecurityAntispamDirectionality = "outbound"
	SecurityAntispamDirectionality_Unknown  SecurityAntispamDirectionality = "unknown"
)

func PossibleValuesForSecurityAntispamDirectionality() []string {
	return []string{
		string(SecurityAntispamDirectionality_Inbound),
		string(SecurityAntispamDirectionality_IntraOrg),
		string(SecurityAntispamDirectionality_Outbound),
		string(SecurityAntispamDirectionality_Unknown),
	}
}

func (s *SecurityAntispamDirectionality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAntispamDirectionality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAntispamDirectionality(input string) (*SecurityAntispamDirectionality, error) {
	vals := map[string]SecurityAntispamDirectionality{
		"inbound":  SecurityAntispamDirectionality_Inbound,
		"intraorg": SecurityAntispamDirectionality_IntraOrg,
		"outbound": SecurityAntispamDirectionality_Outbound,
		"unknown":  SecurityAntispamDirectionality_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAntispamDirectionality(input)
	return &out, nil
}
