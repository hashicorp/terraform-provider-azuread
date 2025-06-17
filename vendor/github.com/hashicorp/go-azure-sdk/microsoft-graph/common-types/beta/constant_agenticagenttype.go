package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgenticAgentType string

const (
	AgenticAgentType_AgenticApp         AgenticAgentType = "agenticApp"
	AgenticAgentType_AgenticAppBuilder  AgenticAgentType = "agenticAppBuilder"
	AgenticAgentType_AgenticAppInstance AgenticAgentType = "agenticAppInstance"
	AgenticAgentType_NotAgentic         AgenticAgentType = "notAgentic"
)

func PossibleValuesForAgenticAgentType() []string {
	return []string{
		string(AgenticAgentType_AgenticApp),
		string(AgenticAgentType_AgenticAppBuilder),
		string(AgenticAgentType_AgenticAppInstance),
		string(AgenticAgentType_NotAgentic),
	}
}

func (s *AgenticAgentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgenticAgentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgenticAgentType(input string) (*AgenticAgentType, error) {
	vals := map[string]AgenticAgentType{
		"agenticapp":         AgenticAgentType_AgenticApp,
		"agenticappbuilder":  AgenticAgentType_AgenticAppBuilder,
		"agenticappinstance": AgenticAgentType_AgenticAppInstance,
		"notagentic":         AgenticAgentType_NotAgentic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgenticAgentType(input)
	return &out, nil
}
