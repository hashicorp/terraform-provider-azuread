package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeFlowBehavior string

const (
	AttributeFlowBehavior_FlowAlways      AttributeFlowBehavior = "FlowAlways"
	AttributeFlowBehavior_FlowWhenChanged AttributeFlowBehavior = "FlowWhenChanged"
)

func PossibleValuesForAttributeFlowBehavior() []string {
	return []string{
		string(AttributeFlowBehavior_FlowAlways),
		string(AttributeFlowBehavior_FlowWhenChanged),
	}
}

func (s *AttributeFlowBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeFlowBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeFlowBehavior(input string) (*AttributeFlowBehavior, error) {
	vals := map[string]AttributeFlowBehavior{
		"flowalways":      AttributeFlowBehavior_FlowAlways,
		"flowwhenchanged": AttributeFlowBehavior_FlowWhenChanged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeFlowBehavior(input)
	return &out, nil
}
