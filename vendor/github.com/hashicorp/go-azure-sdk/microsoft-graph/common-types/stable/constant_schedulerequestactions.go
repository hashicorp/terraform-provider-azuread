package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleRequestActions string

const (
	ScheduleRequestActions_AdminAssign    ScheduleRequestActions = "adminAssign"
	ScheduleRequestActions_AdminExtend    ScheduleRequestActions = "adminExtend"
	ScheduleRequestActions_AdminRemove    ScheduleRequestActions = "adminRemove"
	ScheduleRequestActions_AdminRenew     ScheduleRequestActions = "adminRenew"
	ScheduleRequestActions_AdminUpdate    ScheduleRequestActions = "adminUpdate"
	ScheduleRequestActions_SelfActivate   ScheduleRequestActions = "selfActivate"
	ScheduleRequestActions_SelfDeactivate ScheduleRequestActions = "selfDeactivate"
	ScheduleRequestActions_SelfExtend     ScheduleRequestActions = "selfExtend"
	ScheduleRequestActions_SelfRenew      ScheduleRequestActions = "selfRenew"
)

func PossibleValuesForScheduleRequestActions() []string {
	return []string{
		string(ScheduleRequestActions_AdminAssign),
		string(ScheduleRequestActions_AdminExtend),
		string(ScheduleRequestActions_AdminRemove),
		string(ScheduleRequestActions_AdminRenew),
		string(ScheduleRequestActions_AdminUpdate),
		string(ScheduleRequestActions_SelfActivate),
		string(ScheduleRequestActions_SelfDeactivate),
		string(ScheduleRequestActions_SelfExtend),
		string(ScheduleRequestActions_SelfRenew),
	}
}

func (s *ScheduleRequestActions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleRequestActions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleRequestActions(input string) (*ScheduleRequestActions, error) {
	vals := map[string]ScheduleRequestActions{
		"adminassign":    ScheduleRequestActions_AdminAssign,
		"adminextend":    ScheduleRequestActions_AdminExtend,
		"adminremove":    ScheduleRequestActions_AdminRemove,
		"adminrenew":     ScheduleRequestActions_AdminRenew,
		"adminupdate":    ScheduleRequestActions_AdminUpdate,
		"selfactivate":   ScheduleRequestActions_SelfActivate,
		"selfdeactivate": ScheduleRequestActions_SelfDeactivate,
		"selfextend":     ScheduleRequestActions_SelfExtend,
		"selfrenew":      ScheduleRequestActions_SelfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleRequestActions(input)
	return &out, nil
}
