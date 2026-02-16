package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerCreationSourceKind string

const (
	PlannerCreationSourceKind_External    PlannerCreationSourceKind = "external"
	PlannerCreationSourceKind_None        PlannerCreationSourceKind = "none"
	PlannerCreationSourceKind_Publication PlannerCreationSourceKind = "publication"
)

func PossibleValuesForPlannerCreationSourceKind() []string {
	return []string{
		string(PlannerCreationSourceKind_External),
		string(PlannerCreationSourceKind_None),
		string(PlannerCreationSourceKind_Publication),
	}
}

func (s *PlannerCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerCreationSourceKind(input string) (*PlannerCreationSourceKind, error) {
	vals := map[string]PlannerCreationSourceKind{
		"external":    PlannerCreationSourceKind_External,
		"none":        PlannerCreationSourceKind_None,
		"publication": PlannerCreationSourceKind_Publication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerCreationSourceKind(input)
	return &out, nil
}
