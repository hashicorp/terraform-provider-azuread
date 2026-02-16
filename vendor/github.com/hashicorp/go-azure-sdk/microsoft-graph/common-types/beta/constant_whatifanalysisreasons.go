package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WhatIfAnalysisReasons string

const (
	WhatIfAnalysisReasons_Application           WhatIfAnalysisReasons = "application"
	WhatIfAnalysisReasons_AuthenticationContext WhatIfAnalysisReasons = "authenticationContext"
	WhatIfAnalysisReasons_AuthenticationFlow    WhatIfAnalysisReasons = "authenticationFlow"
	WhatIfAnalysisReasons_ClientApps            WhatIfAnalysisReasons = "clientApps"
	WhatIfAnalysisReasons_DevicePlatform        WhatIfAnalysisReasons = "devicePlatform"
	WhatIfAnalysisReasons_Devices               WhatIfAnalysisReasons = "devices"
	WhatIfAnalysisReasons_EmptyPolicy           WhatIfAnalysisReasons = "emptyPolicy"
	WhatIfAnalysisReasons_InsiderRisk           WhatIfAnalysisReasons = "insiderRisk"
	WhatIfAnalysisReasons_InvalidCondition      WhatIfAnalysisReasons = "invalidCondition"
	WhatIfAnalysisReasons_InvalidPolicy         WhatIfAnalysisReasons = "invalidPolicy"
	WhatIfAnalysisReasons_Location              WhatIfAnalysisReasons = "location"
	WhatIfAnalysisReasons_NotEnoughInformation  WhatIfAnalysisReasons = "notEnoughInformation"
	WhatIfAnalysisReasons_NotSet                WhatIfAnalysisReasons = "notSet"
	WhatIfAnalysisReasons_PolicyNotEnabled      WhatIfAnalysisReasons = "policyNotEnabled"
	WhatIfAnalysisReasons_SignInRisk            WhatIfAnalysisReasons = "signInRisk"
	WhatIfAnalysisReasons_Time                  WhatIfAnalysisReasons = "time"
	WhatIfAnalysisReasons_UserActions           WhatIfAnalysisReasons = "userActions"
	WhatIfAnalysisReasons_UserRisk              WhatIfAnalysisReasons = "userRisk"
	WhatIfAnalysisReasons_Users                 WhatIfAnalysisReasons = "users"
	WhatIfAnalysisReasons_WorkloadIdentities    WhatIfAnalysisReasons = "workloadIdentities"
)

func PossibleValuesForWhatIfAnalysisReasons() []string {
	return []string{
		string(WhatIfAnalysisReasons_Application),
		string(WhatIfAnalysisReasons_AuthenticationContext),
		string(WhatIfAnalysisReasons_AuthenticationFlow),
		string(WhatIfAnalysisReasons_ClientApps),
		string(WhatIfAnalysisReasons_DevicePlatform),
		string(WhatIfAnalysisReasons_Devices),
		string(WhatIfAnalysisReasons_EmptyPolicy),
		string(WhatIfAnalysisReasons_InsiderRisk),
		string(WhatIfAnalysisReasons_InvalidCondition),
		string(WhatIfAnalysisReasons_InvalidPolicy),
		string(WhatIfAnalysisReasons_Location),
		string(WhatIfAnalysisReasons_NotEnoughInformation),
		string(WhatIfAnalysisReasons_NotSet),
		string(WhatIfAnalysisReasons_PolicyNotEnabled),
		string(WhatIfAnalysisReasons_SignInRisk),
		string(WhatIfAnalysisReasons_Time),
		string(WhatIfAnalysisReasons_UserActions),
		string(WhatIfAnalysisReasons_UserRisk),
		string(WhatIfAnalysisReasons_Users),
		string(WhatIfAnalysisReasons_WorkloadIdentities),
	}
}

func (s *WhatIfAnalysisReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWhatIfAnalysisReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWhatIfAnalysisReasons(input string) (*WhatIfAnalysisReasons, error) {
	vals := map[string]WhatIfAnalysisReasons{
		"application":           WhatIfAnalysisReasons_Application,
		"authenticationcontext": WhatIfAnalysisReasons_AuthenticationContext,
		"authenticationflow":    WhatIfAnalysisReasons_AuthenticationFlow,
		"clientapps":            WhatIfAnalysisReasons_ClientApps,
		"deviceplatform":        WhatIfAnalysisReasons_DevicePlatform,
		"devices":               WhatIfAnalysisReasons_Devices,
		"emptypolicy":           WhatIfAnalysisReasons_EmptyPolicy,
		"insiderrisk":           WhatIfAnalysisReasons_InsiderRisk,
		"invalidcondition":      WhatIfAnalysisReasons_InvalidCondition,
		"invalidpolicy":         WhatIfAnalysisReasons_InvalidPolicy,
		"location":              WhatIfAnalysisReasons_Location,
		"notenoughinformation":  WhatIfAnalysisReasons_NotEnoughInformation,
		"notset":                WhatIfAnalysisReasons_NotSet,
		"policynotenabled":      WhatIfAnalysisReasons_PolicyNotEnabled,
		"signinrisk":            WhatIfAnalysisReasons_SignInRisk,
		"time":                  WhatIfAnalysisReasons_Time,
		"useractions":           WhatIfAnalysisReasons_UserActions,
		"userrisk":              WhatIfAnalysisReasons_UserRisk,
		"users":                 WhatIfAnalysisReasons_Users,
		"workloadidentities":    WhatIfAnalysisReasons_WorkloadIdentities,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WhatIfAnalysisReasons(input)
	return &out, nil
}
