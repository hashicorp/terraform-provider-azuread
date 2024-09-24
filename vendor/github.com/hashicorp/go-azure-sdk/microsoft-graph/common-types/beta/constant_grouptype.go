package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupType string

const (
	GroupType_AzureAD       GroupType = "azureAD"
	GroupType_UnifiedGroups GroupType = "unifiedGroups"
)

func PossibleValuesForGroupType() []string {
	return []string{
		string(GroupType_AzureAD),
		string(GroupType_UnifiedGroups),
	}
}

func (s *GroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupType(input string) (*GroupType, error) {
	vals := map[string]GroupType{
		"azuread":       GroupType_AzureAD,
		"unifiedgroups": GroupType_UnifiedGroups,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupType(input)
	return &out, nil
}
