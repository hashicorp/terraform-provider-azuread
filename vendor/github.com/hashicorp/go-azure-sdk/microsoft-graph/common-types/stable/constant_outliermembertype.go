package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlierMemberType string

const (
	OutlierMemberType_User OutlierMemberType = "user"
)

func PossibleValuesForOutlierMemberType() []string {
	return []string{
		string(OutlierMemberType_User),
	}
}

func (s *OutlierMemberType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlierMemberType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlierMemberType(input string) (*OutlierMemberType, error) {
	vals := map[string]OutlierMemberType{
		"user": OutlierMemberType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlierMemberType(input)
	return &out, nil
}
