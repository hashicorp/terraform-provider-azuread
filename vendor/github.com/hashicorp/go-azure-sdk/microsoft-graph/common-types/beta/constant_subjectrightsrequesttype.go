package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestType string

const (
	SubjectRightsRequestType_Access       SubjectRightsRequestType = "access"
	SubjectRightsRequestType_Delete       SubjectRightsRequestType = "delete"
	SubjectRightsRequestType_Export       SubjectRightsRequestType = "export"
	SubjectRightsRequestType_TagForAction SubjectRightsRequestType = "tagForAction"
)

func PossibleValuesForSubjectRightsRequestType() []string {
	return []string{
		string(SubjectRightsRequestType_Access),
		string(SubjectRightsRequestType_Delete),
		string(SubjectRightsRequestType_Export),
		string(SubjectRightsRequestType_TagForAction),
	}
}

func (s *SubjectRightsRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestType(input string) (*SubjectRightsRequestType, error) {
	vals := map[string]SubjectRightsRequestType{
		"access":       SubjectRightsRequestType_Access,
		"delete":       SubjectRightsRequestType_Delete,
		"export":       SubjectRightsRequestType_Export,
		"tagforaction": SubjectRightsRequestType_TagForAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestType(input)
	return &out, nil
}
