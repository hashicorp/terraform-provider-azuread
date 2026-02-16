package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRuleKind string

const (
	PlannerRuleKind_BucketRule PlannerRuleKind = "bucketRule"
	PlannerRuleKind_PlanRule   PlannerRuleKind = "planRule"
	PlannerRuleKind_TaskRule   PlannerRuleKind = "taskRule"
)

func PossibleValuesForPlannerRuleKind() []string {
	return []string{
		string(PlannerRuleKind_BucketRule),
		string(PlannerRuleKind_PlanRule),
		string(PlannerRuleKind_TaskRule),
	}
}

func (s *PlannerRuleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerRuleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerRuleKind(input string) (*PlannerRuleKind, error) {
	vals := map[string]PlannerRuleKind{
		"bucketrule": PlannerRuleKind_BucketRule,
		"planrule":   PlannerRuleKind_PlanRule,
		"taskrule":   PlannerRuleKind_TaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerRuleKind(input)
	return &out, nil
}
