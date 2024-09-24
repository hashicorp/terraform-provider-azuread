package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAccessType string

const (
	GroupAccessType_None    GroupAccessType = "none"
	GroupAccessType_Private GroupAccessType = "private"
	GroupAccessType_Public  GroupAccessType = "public"
	GroupAccessType_Secret  GroupAccessType = "secret"
)

func PossibleValuesForGroupAccessType() []string {
	return []string{
		string(GroupAccessType_None),
		string(GroupAccessType_Private),
		string(GroupAccessType_Public),
		string(GroupAccessType_Secret),
	}
}

func (s *GroupAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupAccessType(input string) (*GroupAccessType, error) {
	vals := map[string]GroupAccessType{
		"none":    GroupAccessType_None,
		"private": GroupAccessType_Private,
		"public":  GroupAccessType_Public,
		"secret":  GroupAccessType_Secret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupAccessType(input)
	return &out, nil
}
