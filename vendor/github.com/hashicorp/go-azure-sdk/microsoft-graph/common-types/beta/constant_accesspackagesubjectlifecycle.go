package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubjectLifecycle string

const (
	AccessPackageSubjectLifecycle_Governed    AccessPackageSubjectLifecycle = "governed"
	AccessPackageSubjectLifecycle_NotDefined  AccessPackageSubjectLifecycle = "notDefined"
	AccessPackageSubjectLifecycle_NotGoverned AccessPackageSubjectLifecycle = "notGoverned"
)

func PossibleValuesForAccessPackageSubjectLifecycle() []string {
	return []string{
		string(AccessPackageSubjectLifecycle_Governed),
		string(AccessPackageSubjectLifecycle_NotDefined),
		string(AccessPackageSubjectLifecycle_NotGoverned),
	}
}

func (s *AccessPackageSubjectLifecycle) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageSubjectLifecycle(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageSubjectLifecycle(input string) (*AccessPackageSubjectLifecycle, error) {
	vals := map[string]AccessPackageSubjectLifecycle{
		"governed":    AccessPackageSubjectLifecycle_Governed,
		"notdefined":  AccessPackageSubjectLifecycle_NotDefined,
		"notgoverned": AccessPackageSubjectLifecycle_NotGoverned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageSubjectLifecycle(input)
	return &out, nil
}
