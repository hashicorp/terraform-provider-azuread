package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppRelationshipType string

const (
	MobileAppRelationshipType_Child  MobileAppRelationshipType = "child"
	MobileAppRelationshipType_Parent MobileAppRelationshipType = "parent"
)

func PossibleValuesForMobileAppRelationshipType() []string {
	return []string{
		string(MobileAppRelationshipType_Child),
		string(MobileAppRelationshipType_Parent),
	}
}

func (s *MobileAppRelationshipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppRelationshipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppRelationshipType(input string) (*MobileAppRelationshipType, error) {
	vals := map[string]MobileAppRelationshipType{
		"child":  MobileAppRelationshipType_Child,
		"parent": MobileAppRelationshipType_Parent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppRelationshipType(input)
	return &out, nil
}
