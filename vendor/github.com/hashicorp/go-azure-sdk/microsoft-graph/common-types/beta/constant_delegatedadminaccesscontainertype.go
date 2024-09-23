package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminAccessContainerType string

const (
	DelegatedAdminAccessContainerType_SecurityGroup DelegatedAdminAccessContainerType = "securityGroup"
)

func PossibleValuesForDelegatedAdminAccessContainerType() []string {
	return []string{
		string(DelegatedAdminAccessContainerType_SecurityGroup),
	}
}

func (s *DelegatedAdminAccessContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminAccessContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminAccessContainerType(input string) (*DelegatedAdminAccessContainerType, error) {
	vals := map[string]DelegatedAdminAccessContainerType{
		"securitygroup": DelegatedAdminAccessContainerType_SecurityGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminAccessContainerType(input)
	return &out, nil
}
