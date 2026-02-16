package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FrontlineCloudPCAccessState string

const (
	FrontlineCloudPCAccessState_Activating          FrontlineCloudPCAccessState = "activating"
	FrontlineCloudPCAccessState_ActivationFailed    FrontlineCloudPCAccessState = "activationFailed"
	FrontlineCloudPCAccessState_Active              FrontlineCloudPCAccessState = "active"
	FrontlineCloudPCAccessState_NoLicensesAvailable FrontlineCloudPCAccessState = "noLicensesAvailable"
	FrontlineCloudPCAccessState_StandbyMode         FrontlineCloudPCAccessState = "standbyMode"
	FrontlineCloudPCAccessState_Unassigned          FrontlineCloudPCAccessState = "unassigned"
)

func PossibleValuesForFrontlineCloudPCAccessState() []string {
	return []string{
		string(FrontlineCloudPCAccessState_Activating),
		string(FrontlineCloudPCAccessState_ActivationFailed),
		string(FrontlineCloudPCAccessState_Active),
		string(FrontlineCloudPCAccessState_NoLicensesAvailable),
		string(FrontlineCloudPCAccessState_StandbyMode),
		string(FrontlineCloudPCAccessState_Unassigned),
	}
}

func (s *FrontlineCloudPCAccessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFrontlineCloudPCAccessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFrontlineCloudPCAccessState(input string) (*FrontlineCloudPCAccessState, error) {
	vals := map[string]FrontlineCloudPCAccessState{
		"activating":          FrontlineCloudPCAccessState_Activating,
		"activationfailed":    FrontlineCloudPCAccessState_ActivationFailed,
		"active":              FrontlineCloudPCAccessState_Active,
		"nolicensesavailable": FrontlineCloudPCAccessState_NoLicensesAvailable,
		"standbymode":         FrontlineCloudPCAccessState_StandbyMode,
		"unassigned":          FrontlineCloudPCAccessState_Unassigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FrontlineCloudPCAccessState(input)
	return &out, nil
}
