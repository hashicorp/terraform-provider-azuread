package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationType string

const (
	PrivilegeManagementElevationType_SupportApprovedElevation PrivilegeManagementElevationType = "supportApprovedElevation"
	PrivilegeManagementElevationType_Undetermined             PrivilegeManagementElevationType = "undetermined"
	PrivilegeManagementElevationType_UnmanagedElevation       PrivilegeManagementElevationType = "unmanagedElevation"
	PrivilegeManagementElevationType_UserConfirmedElevation   PrivilegeManagementElevationType = "userConfirmedElevation"
	PrivilegeManagementElevationType_ZeroTouchElevation       PrivilegeManagementElevationType = "zeroTouchElevation"
)

func PossibleValuesForPrivilegeManagementElevationType() []string {
	return []string{
		string(PrivilegeManagementElevationType_SupportApprovedElevation),
		string(PrivilegeManagementElevationType_Undetermined),
		string(PrivilegeManagementElevationType_UnmanagedElevation),
		string(PrivilegeManagementElevationType_UserConfirmedElevation),
		string(PrivilegeManagementElevationType_ZeroTouchElevation),
	}
}

func (s *PrivilegeManagementElevationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegeManagementElevationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegeManagementElevationType(input string) (*PrivilegeManagementElevationType, error) {
	vals := map[string]PrivilegeManagementElevationType{
		"supportapprovedelevation": PrivilegeManagementElevationType_SupportApprovedElevation,
		"undetermined":             PrivilegeManagementElevationType_Undetermined,
		"unmanagedelevation":       PrivilegeManagementElevationType_UnmanagedElevation,
		"userconfirmedelevation":   PrivilegeManagementElevationType_UserConfirmedElevation,
		"zerotouchelevation":       PrivilegeManagementElevationType_ZeroTouchElevation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementElevationType(input)
	return &out, nil
}
