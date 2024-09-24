package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReadiness string

const (
	GroupPolicyMigrationReadiness_Complete      GroupPolicyMigrationReadiness = "complete"
	GroupPolicyMigrationReadiness_Error         GroupPolicyMigrationReadiness = "error"
	GroupPolicyMigrationReadiness_None          GroupPolicyMigrationReadiness = "none"
	GroupPolicyMigrationReadiness_NotApplicable GroupPolicyMigrationReadiness = "notApplicable"
	GroupPolicyMigrationReadiness_Partial       GroupPolicyMigrationReadiness = "partial"
)

func PossibleValuesForGroupPolicyMigrationReadiness() []string {
	return []string{
		string(GroupPolicyMigrationReadiness_Complete),
		string(GroupPolicyMigrationReadiness_Error),
		string(GroupPolicyMigrationReadiness_None),
		string(GroupPolicyMigrationReadiness_NotApplicable),
		string(GroupPolicyMigrationReadiness_Partial),
	}
}

func (s *GroupPolicyMigrationReadiness) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyMigrationReadiness(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyMigrationReadiness(input string) (*GroupPolicyMigrationReadiness, error) {
	vals := map[string]GroupPolicyMigrationReadiness{
		"complete":      GroupPolicyMigrationReadiness_Complete,
		"error":         GroupPolicyMigrationReadiness_Error,
		"none":          GroupPolicyMigrationReadiness_None,
		"notapplicable": GroupPolicyMigrationReadiness_NotApplicable,
		"partial":       GroupPolicyMigrationReadiness_Partial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyMigrationReadiness(input)
	return &out, nil
}
