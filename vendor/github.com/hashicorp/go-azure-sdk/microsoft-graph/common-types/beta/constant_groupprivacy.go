package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPrivacy string

const (
	GroupPrivacy_Private     GroupPrivacy = "private"
	GroupPrivacy_Public      GroupPrivacy = "public"
	GroupPrivacy_Unspecified GroupPrivacy = "unspecified"
)

func PossibleValuesForGroupPrivacy() []string {
	return []string{
		string(GroupPrivacy_Private),
		string(GroupPrivacy_Public),
		string(GroupPrivacy_Unspecified),
	}
}

func (s *GroupPrivacy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPrivacy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPrivacy(input string) (*GroupPrivacy, error) {
	vals := map[string]GroupPrivacy{
		"private":     GroupPrivacy_Private,
		"public":      GroupPrivacy_Public,
		"unspecified": GroupPrivacy_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPrivacy(input)
	return &out, nil
}
