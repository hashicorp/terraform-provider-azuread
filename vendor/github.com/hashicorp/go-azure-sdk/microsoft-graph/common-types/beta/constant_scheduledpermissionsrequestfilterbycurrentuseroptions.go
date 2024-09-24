package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledPermissionsRequestFilterByCurrentUserOptions string

const (
	ScheduledPermissionsRequestFilterByCurrentUserOptions_Approver  ScheduledPermissionsRequestFilterByCurrentUserOptions = "approver"
	ScheduledPermissionsRequestFilterByCurrentUserOptions_CreatedBy ScheduledPermissionsRequestFilterByCurrentUserOptions = "createdBy"
	ScheduledPermissionsRequestFilterByCurrentUserOptions_Principal ScheduledPermissionsRequestFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForScheduledPermissionsRequestFilterByCurrentUserOptions() []string {
	return []string{
		string(ScheduledPermissionsRequestFilterByCurrentUserOptions_Approver),
		string(ScheduledPermissionsRequestFilterByCurrentUserOptions_CreatedBy),
		string(ScheduledPermissionsRequestFilterByCurrentUserOptions_Principal),
	}
}

func (s *ScheduledPermissionsRequestFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduledPermissionsRequestFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduledPermissionsRequestFilterByCurrentUserOptions(input string) (*ScheduledPermissionsRequestFilterByCurrentUserOptions, error) {
	vals := map[string]ScheduledPermissionsRequestFilterByCurrentUserOptions{
		"approver":  ScheduledPermissionsRequestFilterByCurrentUserOptions_Approver,
		"createdby": ScheduledPermissionsRequestFilterByCurrentUserOptions_CreatedBy,
		"principal": ScheduledPermissionsRequestFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduledPermissionsRequestFilterByCurrentUserOptions(input)
	return &out, nil
}
