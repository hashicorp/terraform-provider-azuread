package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubjectType string

const (
	AccessPackageSubjectType_NotSpecified     AccessPackageSubjectType = "notSpecified"
	AccessPackageSubjectType_ServicePrincipal AccessPackageSubjectType = "servicePrincipal"
	AccessPackageSubjectType_User             AccessPackageSubjectType = "user"
)

func PossibleValuesForAccessPackageSubjectType() []string {
	return []string{
		string(AccessPackageSubjectType_NotSpecified),
		string(AccessPackageSubjectType_ServicePrincipal),
		string(AccessPackageSubjectType_User),
	}
}

func (s *AccessPackageSubjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageSubjectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageSubjectType(input string) (*AccessPackageSubjectType, error) {
	vals := map[string]AccessPackageSubjectType{
		"notspecified":     AccessPackageSubjectType_NotSpecified,
		"serviceprincipal": AccessPackageSubjectType_ServicePrincipal,
		"user":             AccessPackageSubjectType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageSubjectType(input)
	return &out, nil
}
