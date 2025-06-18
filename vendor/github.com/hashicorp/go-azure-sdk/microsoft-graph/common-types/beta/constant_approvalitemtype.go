package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalItemType string

const (
	ApprovalItemType_Basic          ApprovalItemType = "basic"
	ApprovalItemType_BasicAwaitAll  ApprovalItemType = "basicAwaitAll"
	ApprovalItemType_Custom         ApprovalItemType = "custom"
	ApprovalItemType_CustomAwaitAll ApprovalItemType = "customAwaitAll"
)

func PossibleValuesForApprovalItemType() []string {
	return []string{
		string(ApprovalItemType_Basic),
		string(ApprovalItemType_BasicAwaitAll),
		string(ApprovalItemType_Custom),
		string(ApprovalItemType_CustomAwaitAll),
	}
}

func (s *ApprovalItemType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApprovalItemType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApprovalItemType(input string) (*ApprovalItemType, error) {
	vals := map[string]ApprovalItemType{
		"basic":          ApprovalItemType_Basic,
		"basicawaitall":  ApprovalItemType_BasicAwaitAll,
		"custom":         ApprovalItemType_Custom,
		"customawaitall": ApprovalItemType_CustomAwaitAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApprovalItemType(input)
	return &out, nil
}
