package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedAudiences string

const (
	AllowedAudiences_Contacts               AllowedAudiences = "contacts"
	AllowedAudiences_Everyone               AllowedAudiences = "everyone"
	AllowedAudiences_Family                 AllowedAudiences = "family"
	AllowedAudiences_FederatedOrganizations AllowedAudiences = "federatedOrganizations"
	AllowedAudiences_GroupMembers           AllowedAudiences = "groupMembers"
	AllowedAudiences_Me                     AllowedAudiences = "me"
	AllowedAudiences_Organization           AllowedAudiences = "organization"
)

func PossibleValuesForAllowedAudiences() []string {
	return []string{
		string(AllowedAudiences_Contacts),
		string(AllowedAudiences_Everyone),
		string(AllowedAudiences_Family),
		string(AllowedAudiences_FederatedOrganizations),
		string(AllowedAudiences_GroupMembers),
		string(AllowedAudiences_Me),
		string(AllowedAudiences_Organization),
	}
}

func (s *AllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedAudiences(input string) (*AllowedAudiences, error) {
	vals := map[string]AllowedAudiences{
		"contacts":               AllowedAudiences_Contacts,
		"everyone":               AllowedAudiences_Everyone,
		"family":                 AllowedAudiences_Family,
		"federatedorganizations": AllowedAudiences_FederatedOrganizations,
		"groupmembers":           AllowedAudiences_GroupMembers,
		"me":                     AllowedAudiences_Me,
		"organization":           AllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedAudiences(input)
	return &out, nil
}
