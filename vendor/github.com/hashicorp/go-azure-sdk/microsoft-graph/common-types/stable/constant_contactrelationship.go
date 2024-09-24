package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactRelationship string

const (
	ContactRelationship_Aide     ContactRelationship = "aide"
	ContactRelationship_Child    ContactRelationship = "child"
	ContactRelationship_Doctor   ContactRelationship = "doctor"
	ContactRelationship_Guardian ContactRelationship = "guardian"
	ContactRelationship_Other    ContactRelationship = "other"
	ContactRelationship_Parent   ContactRelationship = "parent"
	ContactRelationship_Relative ContactRelationship = "relative"
)

func PossibleValuesForContactRelationship() []string {
	return []string{
		string(ContactRelationship_Aide),
		string(ContactRelationship_Child),
		string(ContactRelationship_Doctor),
		string(ContactRelationship_Guardian),
		string(ContactRelationship_Other),
		string(ContactRelationship_Parent),
		string(ContactRelationship_Relative),
	}
}

func (s *ContactRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContactRelationship(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContactRelationship(input string) (*ContactRelationship, error) {
	vals := map[string]ContactRelationship{
		"aide":     ContactRelationship_Aide,
		"child":    ContactRelationship_Child,
		"doctor":   ContactRelationship_Doctor,
		"guardian": ContactRelationship_Guardian,
		"other":    ContactRelationship_Other,
		"parent":   ContactRelationship_Parent,
		"relative": ContactRelationship_Relative,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContactRelationship(input)
	return &out, nil
}
