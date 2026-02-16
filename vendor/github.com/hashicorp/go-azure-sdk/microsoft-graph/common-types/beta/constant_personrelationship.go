package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonRelationship string

const (
	PersonRelationship_AlternateContact PersonRelationship = "alternateContact"
	PersonRelationship_Assistant        PersonRelationship = "assistant"
	PersonRelationship_Child            PersonRelationship = "child"
	PersonRelationship_Colleague        PersonRelationship = "colleague"
	PersonRelationship_DirectReport     PersonRelationship = "directReport"
	PersonRelationship_DotLineManager   PersonRelationship = "dotLineManager"
	PersonRelationship_DotLineReport    PersonRelationship = "dotLineReport"
	PersonRelationship_EmergencyContact PersonRelationship = "emergencyContact"
	PersonRelationship_Friend           PersonRelationship = "friend"
	PersonRelationship_Manager          PersonRelationship = "manager"
	PersonRelationship_Other            PersonRelationship = "other"
	PersonRelationship_Parent           PersonRelationship = "parent"
	PersonRelationship_Sibling          PersonRelationship = "sibling"
	PersonRelationship_Sponsor          PersonRelationship = "sponsor"
	PersonRelationship_Spouse           PersonRelationship = "spouse"
)

func PossibleValuesForPersonRelationship() []string {
	return []string{
		string(PersonRelationship_AlternateContact),
		string(PersonRelationship_Assistant),
		string(PersonRelationship_Child),
		string(PersonRelationship_Colleague),
		string(PersonRelationship_DirectReport),
		string(PersonRelationship_DotLineManager),
		string(PersonRelationship_DotLineReport),
		string(PersonRelationship_EmergencyContact),
		string(PersonRelationship_Friend),
		string(PersonRelationship_Manager),
		string(PersonRelationship_Other),
		string(PersonRelationship_Parent),
		string(PersonRelationship_Sibling),
		string(PersonRelationship_Sponsor),
		string(PersonRelationship_Spouse),
	}
}

func (s *PersonRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonRelationship(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonRelationship(input string) (*PersonRelationship, error) {
	vals := map[string]PersonRelationship{
		"alternatecontact": PersonRelationship_AlternateContact,
		"assistant":        PersonRelationship_Assistant,
		"child":            PersonRelationship_Child,
		"colleague":        PersonRelationship_Colleague,
		"directreport":     PersonRelationship_DirectReport,
		"dotlinemanager":   PersonRelationship_DotLineManager,
		"dotlinereport":    PersonRelationship_DotLineReport,
		"emergencycontact": PersonRelationship_EmergencyContact,
		"friend":           PersonRelationship_Friend,
		"manager":          PersonRelationship_Manager,
		"other":            PersonRelationship_Other,
		"parent":           PersonRelationship_Parent,
		"sibling":          PersonRelationship_Sibling,
		"sponsor":          PersonRelationship_Sponsor,
		"spouse":           PersonRelationship_Spouse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonRelationship(input)
	return &out, nil
}
