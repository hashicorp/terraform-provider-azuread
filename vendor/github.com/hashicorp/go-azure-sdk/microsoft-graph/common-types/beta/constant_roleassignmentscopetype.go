package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScopeType string

const (
	RoleAssignmentScopeType_AllDevices                 RoleAssignmentScopeType = "allDevices"
	RoleAssignmentScopeType_AllDevicesAndLicensedUsers RoleAssignmentScopeType = "allDevicesAndLicensedUsers"
	RoleAssignmentScopeType_AllLicensedUsers           RoleAssignmentScopeType = "allLicensedUsers"
	RoleAssignmentScopeType_ResourceScope              RoleAssignmentScopeType = "resourceScope"
)

func PossibleValuesForRoleAssignmentScopeType() []string {
	return []string{
		string(RoleAssignmentScopeType_AllDevices),
		string(RoleAssignmentScopeType_AllDevicesAndLicensedUsers),
		string(RoleAssignmentScopeType_AllLicensedUsers),
		string(RoleAssignmentScopeType_ResourceScope),
	}
}

func (s *RoleAssignmentScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleAssignmentScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleAssignmentScopeType(input string) (*RoleAssignmentScopeType, error) {
	vals := map[string]RoleAssignmentScopeType{
		"alldevices":                 RoleAssignmentScopeType_AllDevices,
		"alldevicesandlicensedusers": RoleAssignmentScopeType_AllDevicesAndLicensedUsers,
		"alllicensedusers":           RoleAssignmentScopeType_AllLicensedUsers,
		"resourcescope":              RoleAssignmentScopeType_ResourceScope,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleAssignmentScopeType(input)
	return &out, nil
}
