package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsIdentityType string

const (
	ExternalConnectorsIdentityType_ExternalGroup ExternalConnectorsIdentityType = "externalGroup"
	ExternalConnectorsIdentityType_Group         ExternalConnectorsIdentityType = "group"
	ExternalConnectorsIdentityType_User          ExternalConnectorsIdentityType = "user"
)

func PossibleValuesForExternalConnectorsIdentityType() []string {
	return []string{
		string(ExternalConnectorsIdentityType_ExternalGroup),
		string(ExternalConnectorsIdentityType_Group),
		string(ExternalConnectorsIdentityType_User),
	}
}

func (s *ExternalConnectorsIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsIdentityType(input string) (*ExternalConnectorsIdentityType, error) {
	vals := map[string]ExternalConnectorsIdentityType{
		"externalgroup": ExternalConnectorsIdentityType_ExternalGroup,
		"group":         ExternalConnectorsIdentityType_Group,
		"user":          ExternalConnectorsIdentityType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsIdentityType(input)
	return &out, nil
}
