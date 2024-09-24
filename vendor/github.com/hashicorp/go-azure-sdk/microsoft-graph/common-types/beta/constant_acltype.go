package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AclType string

const (
	AclType_Everyone             AclType = "everyone"
	AclType_EveryoneExceptGuests AclType = "everyoneExceptGuests"
	AclType_ExternalGroup        AclType = "externalGroup"
	AclType_Group                AclType = "group"
	AclType_User                 AclType = "user"
)

func PossibleValuesForAclType() []string {
	return []string{
		string(AclType_Everyone),
		string(AclType_EveryoneExceptGuests),
		string(AclType_ExternalGroup),
		string(AclType_Group),
		string(AclType_User),
	}
}

func (s *AclType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAclType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAclType(input string) (*AclType, error) {
	vals := map[string]AclType{
		"everyone":             AclType_Everyone,
		"everyoneexceptguests": AclType_EveryoneExceptGuests,
		"externalgroup":        AclType_ExternalGroup,
		"group":                AclType_Group,
		"user":                 AclType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AclType(input)
	return &out, nil
}
