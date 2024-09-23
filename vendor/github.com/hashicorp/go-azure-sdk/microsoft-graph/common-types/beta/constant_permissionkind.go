package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionKind string

const (
	PermissionKind_All                         PermissionKind = "all"
	PermissionKind_AllPermissionsOnResourceApp PermissionKind = "allPermissionsOnResourceApp"
	PermissionKind_Enumerated                  PermissionKind = "enumerated"
)

func PossibleValuesForPermissionKind() []string {
	return []string{
		string(PermissionKind_All),
		string(PermissionKind_AllPermissionsOnResourceApp),
		string(PermissionKind_Enumerated),
	}
}

func (s *PermissionKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionKind(input string) (*PermissionKind, error) {
	vals := map[string]PermissionKind{
		"all":                         PermissionKind_All,
		"allpermissionsonresourceapp": PermissionKind_AllPermissionsOnResourceApp,
		"enumerated":                  PermissionKind_Enumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionKind(input)
	return &out, nil
}
