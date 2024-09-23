package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationSupportedEntities string

const (
	WorkforceIntegrationSupportedEntities_None                 WorkforceIntegrationSupportedEntities = "none"
	WorkforceIntegrationSupportedEntities_OfferShiftRequest    WorkforceIntegrationSupportedEntities = "offerShiftRequest"
	WorkforceIntegrationSupportedEntities_OpenShift            WorkforceIntegrationSupportedEntities = "openShift"
	WorkforceIntegrationSupportedEntities_OpenShiftRequest     WorkforceIntegrationSupportedEntities = "openShiftRequest"
	WorkforceIntegrationSupportedEntities_Shift                WorkforceIntegrationSupportedEntities = "shift"
	WorkforceIntegrationSupportedEntities_SwapRequest          WorkforceIntegrationSupportedEntities = "swapRequest"
	WorkforceIntegrationSupportedEntities_TimeCard             WorkforceIntegrationSupportedEntities = "timeCard"
	WorkforceIntegrationSupportedEntities_TimeOff              WorkforceIntegrationSupportedEntities = "timeOff"
	WorkforceIntegrationSupportedEntities_TimeOffReason        WorkforceIntegrationSupportedEntities = "timeOffReason"
	WorkforceIntegrationSupportedEntities_TimeOffRequest       WorkforceIntegrationSupportedEntities = "timeOffRequest"
	WorkforceIntegrationSupportedEntities_UserShiftPreferences WorkforceIntegrationSupportedEntities = "userShiftPreferences"
)

func PossibleValuesForWorkforceIntegrationSupportedEntities() []string {
	return []string{
		string(WorkforceIntegrationSupportedEntities_None),
		string(WorkforceIntegrationSupportedEntities_OfferShiftRequest),
		string(WorkforceIntegrationSupportedEntities_OpenShift),
		string(WorkforceIntegrationSupportedEntities_OpenShiftRequest),
		string(WorkforceIntegrationSupportedEntities_Shift),
		string(WorkforceIntegrationSupportedEntities_SwapRequest),
		string(WorkforceIntegrationSupportedEntities_TimeCard),
		string(WorkforceIntegrationSupportedEntities_TimeOff),
		string(WorkforceIntegrationSupportedEntities_TimeOffReason),
		string(WorkforceIntegrationSupportedEntities_TimeOffRequest),
		string(WorkforceIntegrationSupportedEntities_UserShiftPreferences),
	}
}

func (s *WorkforceIntegrationSupportedEntities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkforceIntegrationSupportedEntities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkforceIntegrationSupportedEntities(input string) (*WorkforceIntegrationSupportedEntities, error) {
	vals := map[string]WorkforceIntegrationSupportedEntities{
		"none":                 WorkforceIntegrationSupportedEntities_None,
		"offershiftrequest":    WorkforceIntegrationSupportedEntities_OfferShiftRequest,
		"openshift":            WorkforceIntegrationSupportedEntities_OpenShift,
		"openshiftrequest":     WorkforceIntegrationSupportedEntities_OpenShiftRequest,
		"shift":                WorkforceIntegrationSupportedEntities_Shift,
		"swaprequest":          WorkforceIntegrationSupportedEntities_SwapRequest,
		"timecard":             WorkforceIntegrationSupportedEntities_TimeCard,
		"timeoff":              WorkforceIntegrationSupportedEntities_TimeOff,
		"timeoffreason":        WorkforceIntegrationSupportedEntities_TimeOffReason,
		"timeoffrequest":       WorkforceIntegrationSupportedEntities_TimeOffRequest,
		"usershiftpreferences": WorkforceIntegrationSupportedEntities_UserShiftPreferences,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationSupportedEntities(input)
	return &out, nil
}
