package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleScheduleRequestActions string

const (
	UnifiedRoleScheduleRequestActions_AdminAssign    UnifiedRoleScheduleRequestActions = "adminAssign"
	UnifiedRoleScheduleRequestActions_AdminExtend    UnifiedRoleScheduleRequestActions = "adminExtend"
	UnifiedRoleScheduleRequestActions_AdminRemove    UnifiedRoleScheduleRequestActions = "adminRemove"
	UnifiedRoleScheduleRequestActions_AdminRenew     UnifiedRoleScheduleRequestActions = "adminRenew"
	UnifiedRoleScheduleRequestActions_AdminUpdate    UnifiedRoleScheduleRequestActions = "adminUpdate"
	UnifiedRoleScheduleRequestActions_SelfActivate   UnifiedRoleScheduleRequestActions = "selfActivate"
	UnifiedRoleScheduleRequestActions_SelfDeactivate UnifiedRoleScheduleRequestActions = "selfDeactivate"
	UnifiedRoleScheduleRequestActions_SelfExtend     UnifiedRoleScheduleRequestActions = "selfExtend"
	UnifiedRoleScheduleRequestActions_SelfRenew      UnifiedRoleScheduleRequestActions = "selfRenew"
)

func PossibleValuesForUnifiedRoleScheduleRequestActions() []string {
	return []string{
		string(UnifiedRoleScheduleRequestActions_AdminAssign),
		string(UnifiedRoleScheduleRequestActions_AdminExtend),
		string(UnifiedRoleScheduleRequestActions_AdminRemove),
		string(UnifiedRoleScheduleRequestActions_AdminRenew),
		string(UnifiedRoleScheduleRequestActions_AdminUpdate),
		string(UnifiedRoleScheduleRequestActions_SelfActivate),
		string(UnifiedRoleScheduleRequestActions_SelfDeactivate),
		string(UnifiedRoleScheduleRequestActions_SelfExtend),
		string(UnifiedRoleScheduleRequestActions_SelfRenew),
	}
}

func (s *UnifiedRoleScheduleRequestActions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnifiedRoleScheduleRequestActions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnifiedRoleScheduleRequestActions(input string) (*UnifiedRoleScheduleRequestActions, error) {
	vals := map[string]UnifiedRoleScheduleRequestActions{
		"adminassign":    UnifiedRoleScheduleRequestActions_AdminAssign,
		"adminextend":    UnifiedRoleScheduleRequestActions_AdminExtend,
		"adminremove":    UnifiedRoleScheduleRequestActions_AdminRemove,
		"adminrenew":     UnifiedRoleScheduleRequestActions_AdminRenew,
		"adminupdate":    UnifiedRoleScheduleRequestActions_AdminUpdate,
		"selfactivate":   UnifiedRoleScheduleRequestActions_SelfActivate,
		"selfdeactivate": UnifiedRoleScheduleRequestActions_SelfDeactivate,
		"selfextend":     UnifiedRoleScheduleRequestActions_SelfExtend,
		"selfrenew":      UnifiedRoleScheduleRequestActions_SelfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleScheduleRequestActions(input)
	return &out, nil
}
