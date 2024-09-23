package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeType string

const (
	ChangeType_Created ChangeType = "created"
	ChangeType_Deleted ChangeType = "deleted"
	ChangeType_Updated ChangeType = "updated"
)

func PossibleValuesForChangeType() []string {
	return []string{
		string(ChangeType_Created),
		string(ChangeType_Deleted),
		string(ChangeType_Updated),
	}
}

func (s *ChangeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeType(input string) (*ChangeType, error) {
	vals := map[string]ChangeType{
		"created": ChangeType_Created,
		"deleted": ChangeType_Deleted,
		"updated": ChangeType_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeType(input)
	return &out, nil
}
