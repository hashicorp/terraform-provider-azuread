package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentStatus string

const (
	AgentStatus_Active   AgentStatus = "active"
	AgentStatus_Inactive AgentStatus = "inactive"
)

func PossibleValuesForAgentStatus() []string {
	return []string{
		string(AgentStatus_Active),
		string(AgentStatus_Inactive),
	}
}

func (s *AgentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentStatus(input string) (*AgentStatus, error) {
	vals := map[string]AgentStatus{
		"active":   AgentStatus_Active,
		"inactive": AgentStatus_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentStatus(input)
	return &out, nil
}
