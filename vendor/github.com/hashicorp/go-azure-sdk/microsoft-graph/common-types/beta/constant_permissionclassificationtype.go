package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionClassificationType string

const (
	PermissionClassificationType_High   PermissionClassificationType = "high"
	PermissionClassificationType_Low    PermissionClassificationType = "low"
	PermissionClassificationType_Medium PermissionClassificationType = "medium"
)

func PossibleValuesForPermissionClassificationType() []string {
	return []string{
		string(PermissionClassificationType_High),
		string(PermissionClassificationType_Low),
		string(PermissionClassificationType_Medium),
	}
}

func (s *PermissionClassificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionClassificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionClassificationType(input string) (*PermissionClassificationType, error) {
	vals := map[string]PermissionClassificationType{
		"high":   PermissionClassificationType_High,
		"low":    PermissionClassificationType_Low,
		"medium": PermissionClassificationType_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionClassificationType(input)
	return &out, nil
}
