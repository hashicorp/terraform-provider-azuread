package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityPrivacy string

const (
	CommunityPrivacy_Private CommunityPrivacy = "private"
	CommunityPrivacy_Public  CommunityPrivacy = "public"
)

func PossibleValuesForCommunityPrivacy() []string {
	return []string{
		string(CommunityPrivacy_Private),
		string(CommunityPrivacy_Public),
	}
}

func (s *CommunityPrivacy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommunityPrivacy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommunityPrivacy(input string) (*CommunityPrivacy, error) {
	vals := map[string]CommunityPrivacy{
		"private": CommunityPrivacy_Private,
		"public":  CommunityPrivacy_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommunityPrivacy(input)
	return &out, nil
}
