package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisableReason string

const (
	DisableReason_ControllerServiceAppDeleted DisableReason = "controllerServiceAppDeleted"
	DisableReason_InvalidBillingProfile       DisableReason = "invalidBillingProfile"
	DisableReason_None                        DisableReason = "none"
	DisableReason_UserRequested               DisableReason = "userRequested"
)

func PossibleValuesForDisableReason() []string {
	return []string{
		string(DisableReason_ControllerServiceAppDeleted),
		string(DisableReason_InvalidBillingProfile),
		string(DisableReason_None),
		string(DisableReason_UserRequested),
	}
}

func (s *DisableReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDisableReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDisableReason(input string) (*DisableReason, error) {
	vals := map[string]DisableReason{
		"controllerserviceappdeleted": DisableReason_ControllerServiceAppDeleted,
		"invalidbillingprofile":       DisableReason_InvalidBillingProfile,
		"none":                        DisableReason_None,
		"userrequested":               DisableReason_UserRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DisableReason(input)
	return &out, nil
}
