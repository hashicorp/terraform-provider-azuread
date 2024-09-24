package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsRoleType string

const (
	AwsRoleType_Custom AwsRoleType = "custom"
	AwsRoleType_System AwsRoleType = "system"
)

func PossibleValuesForAwsRoleType() []string {
	return []string{
		string(AwsRoleType_Custom),
		string(AwsRoleType_System),
	}
}

func (s *AwsRoleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsRoleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsRoleType(input string) (*AwsRoleType, error) {
	vals := map[string]AwsRoleType{
		"custom": AwsRoleType_Custom,
		"system": AwsRoleType_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsRoleType(input)
	return &out, nil
}
