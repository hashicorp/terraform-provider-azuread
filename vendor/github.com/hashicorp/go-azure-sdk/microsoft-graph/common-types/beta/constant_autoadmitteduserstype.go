package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoAdmittedUsersType string

const (
	AutoAdmittedUsersType_Everyone          AutoAdmittedUsersType = "everyone"
	AutoAdmittedUsersType_EveryoneInCompany AutoAdmittedUsersType = "everyoneInCompany"
)

func PossibleValuesForAutoAdmittedUsersType() []string {
	return []string{
		string(AutoAdmittedUsersType_Everyone),
		string(AutoAdmittedUsersType_EveryoneInCompany),
	}
}

func (s *AutoAdmittedUsersType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoAdmittedUsersType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoAdmittedUsersType(input string) (*AutoAdmittedUsersType, error) {
	vals := map[string]AutoAdmittedUsersType{
		"everyone":          AutoAdmittedUsersType_Everyone,
		"everyoneincompany": AutoAdmittedUsersType_EveryoneInCompany,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoAdmittedUsersType(input)
	return &out, nil
}
