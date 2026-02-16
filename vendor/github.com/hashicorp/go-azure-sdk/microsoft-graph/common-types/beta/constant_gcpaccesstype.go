package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpAccessType string

const (
	GcpAccessType_Private             GcpAccessType = "private"
	GcpAccessType_Public              GcpAccessType = "public"
	GcpAccessType_SubjectToObjectAcls GcpAccessType = "subjectToObjectAcls"
)

func PossibleValuesForGcpAccessType() []string {
	return []string{
		string(GcpAccessType_Private),
		string(GcpAccessType_Public),
		string(GcpAccessType_SubjectToObjectAcls),
	}
}

func (s *GcpAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGcpAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGcpAccessType(input string) (*GcpAccessType, error) {
	vals := map[string]GcpAccessType{
		"private":             GcpAccessType_Private,
		"public":              GcpAccessType_Public,
		"subjecttoobjectacls": GcpAccessType_SubjectToObjectAcls,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GcpAccessType(input)
	return &out, nil
}
