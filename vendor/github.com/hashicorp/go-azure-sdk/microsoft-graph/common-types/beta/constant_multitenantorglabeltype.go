package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrgLabelType string

const (
	MultiTenantOrgLabelType_CustomName MultiTenantOrgLabelType = "customName"
	MultiTenantOrgLabelType_GroupName  MultiTenantOrgLabelType = "groupName"
	MultiTenantOrgLabelType_None       MultiTenantOrgLabelType = "none"
)

func PossibleValuesForMultiTenantOrgLabelType() []string {
	return []string{
		string(MultiTenantOrgLabelType_CustomName),
		string(MultiTenantOrgLabelType_GroupName),
		string(MultiTenantOrgLabelType_None),
	}
}

func (s *MultiTenantOrgLabelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrgLabelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrgLabelType(input string) (*MultiTenantOrgLabelType, error) {
	vals := map[string]MultiTenantOrgLabelType{
		"customname": MultiTenantOrgLabelType_CustomName,
		"groupname":  MultiTenantOrgLabelType_GroupName,
		"none":       MultiTenantOrgLabelType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrgLabelType(input)
	return &out, nil
}
