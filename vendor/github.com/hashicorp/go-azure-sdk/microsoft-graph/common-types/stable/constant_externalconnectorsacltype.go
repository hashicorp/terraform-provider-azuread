package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAclType string

const (
	ExternalConnectorsAclType_Everyone             ExternalConnectorsAclType = "everyone"
	ExternalConnectorsAclType_EveryoneExceptGuests ExternalConnectorsAclType = "everyoneExceptGuests"
	ExternalConnectorsAclType_ExternalGroup        ExternalConnectorsAclType = "externalGroup"
	ExternalConnectorsAclType_Group                ExternalConnectorsAclType = "group"
	ExternalConnectorsAclType_User                 ExternalConnectorsAclType = "user"
)

func PossibleValuesForExternalConnectorsAclType() []string {
	return []string{
		string(ExternalConnectorsAclType_Everyone),
		string(ExternalConnectorsAclType_EveryoneExceptGuests),
		string(ExternalConnectorsAclType_ExternalGroup),
		string(ExternalConnectorsAclType_Group),
		string(ExternalConnectorsAclType_User),
	}
}

func (s *ExternalConnectorsAclType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsAclType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsAclType(input string) (*ExternalConnectorsAclType, error) {
	vals := map[string]ExternalConnectorsAclType{
		"everyone":             ExternalConnectorsAclType_Everyone,
		"everyoneexceptguests": ExternalConnectorsAclType_EveryoneExceptGuests,
		"externalgroup":        ExternalConnectorsAclType_ExternalGroup,
		"group":                ExternalConnectorsAclType_Group,
		"user":                 ExternalConnectorsAclType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsAclType(input)
	return &out, nil
}
