package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIntelligenceProfileKind string

const (
	SecurityIntelligenceProfileKind_Actor SecurityIntelligenceProfileKind = "actor"
	SecurityIntelligenceProfileKind_Tool  SecurityIntelligenceProfileKind = "tool"
)

func PossibleValuesForSecurityIntelligenceProfileKind() []string {
	return []string{
		string(SecurityIntelligenceProfileKind_Actor),
		string(SecurityIntelligenceProfileKind_Tool),
	}
}

func (s *SecurityIntelligenceProfileKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIntelligenceProfileKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIntelligenceProfileKind(input string) (*SecurityIntelligenceProfileKind, error) {
	vals := map[string]SecurityIntelligenceProfileKind{
		"actor": SecurityIntelligenceProfileKind_Actor,
		"tool":  SecurityIntelligenceProfileKind_Tool,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIntelligenceProfileKind(input)
	return &out, nil
}
