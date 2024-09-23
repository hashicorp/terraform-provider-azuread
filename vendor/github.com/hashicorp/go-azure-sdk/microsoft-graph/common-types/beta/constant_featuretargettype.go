package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureTargetType string

const (
	FeatureTargetType_AdministrativeUnit FeatureTargetType = "administrativeUnit"
	FeatureTargetType_Group              FeatureTargetType = "group"
	FeatureTargetType_Role               FeatureTargetType = "role"
)

func PossibleValuesForFeatureTargetType() []string {
	return []string{
		string(FeatureTargetType_AdministrativeUnit),
		string(FeatureTargetType_Group),
		string(FeatureTargetType_Role),
	}
}

func (s *FeatureTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureTargetType(input string) (*FeatureTargetType, error) {
	vals := map[string]FeatureTargetType{
		"administrativeunit": FeatureTargetType_AdministrativeUnit,
		"group":              FeatureTargetType_Group,
		"role":               FeatureTargetType_Role,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureTargetType(input)
	return &out, nil
}
