package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessLevel string

const (
	AccessLevel_Everyone                   AccessLevel = "everyone"
	AccessLevel_Invited                    AccessLevel = "invited"
	AccessLevel_Locked                     AccessLevel = "locked"
	AccessLevel_SameEnterprise             AccessLevel = "sameEnterprise"
	AccessLevel_SameEnterpriseAndFederated AccessLevel = "sameEnterpriseAndFederated"
)

func PossibleValuesForAccessLevel() []string {
	return []string{
		string(AccessLevel_Everyone),
		string(AccessLevel_Invited),
		string(AccessLevel_Locked),
		string(AccessLevel_SameEnterprise),
		string(AccessLevel_SameEnterpriseAndFederated),
	}
}

func (s *AccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessLevel(input string) (*AccessLevel, error) {
	vals := map[string]AccessLevel{
		"everyone":                   AccessLevel_Everyone,
		"invited":                    AccessLevel_Invited,
		"locked":                     AccessLevel_Locked,
		"sameenterprise":             AccessLevel_SameEnterprise,
		"sameenterpriseandfederated": AccessLevel_SameEnterpriseAndFederated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessLevel(input)
	return &out, nil
}
