package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessWhatIfReasons string

const (
	ConditionalAccessWhatIfReasons_Application           ConditionalAccessWhatIfReasons = "application"
	ConditionalAccessWhatIfReasons_AuthenticationContext ConditionalAccessWhatIfReasons = "authenticationContext"
	ConditionalAccessWhatIfReasons_AuthenticationFlow    ConditionalAccessWhatIfReasons = "authenticationFlow"
	ConditionalAccessWhatIfReasons_ClientApps            ConditionalAccessWhatIfReasons = "clientApps"
	ConditionalAccessWhatIfReasons_DevicePlatform        ConditionalAccessWhatIfReasons = "devicePlatform"
	ConditionalAccessWhatIfReasons_Devices               ConditionalAccessWhatIfReasons = "devices"
	ConditionalAccessWhatIfReasons_EmptyPolicy           ConditionalAccessWhatIfReasons = "emptyPolicy"
	ConditionalAccessWhatIfReasons_InsiderRisk           ConditionalAccessWhatIfReasons = "insiderRisk"
	ConditionalAccessWhatIfReasons_InvalidCondition      ConditionalAccessWhatIfReasons = "invalidCondition"
	ConditionalAccessWhatIfReasons_InvalidPolicy         ConditionalAccessWhatIfReasons = "invalidPolicy"
	ConditionalAccessWhatIfReasons_Location              ConditionalAccessWhatIfReasons = "location"
	ConditionalAccessWhatIfReasons_NotEnoughInformation  ConditionalAccessWhatIfReasons = "notEnoughInformation"
	ConditionalAccessWhatIfReasons_NotSet                ConditionalAccessWhatIfReasons = "notSet"
	ConditionalAccessWhatIfReasons_PolicyNotEnabled      ConditionalAccessWhatIfReasons = "policyNotEnabled"
	ConditionalAccessWhatIfReasons_SignInRisk            ConditionalAccessWhatIfReasons = "signInRisk"
	ConditionalAccessWhatIfReasons_Time                  ConditionalAccessWhatIfReasons = "time"
	ConditionalAccessWhatIfReasons_UserActions           ConditionalAccessWhatIfReasons = "userActions"
	ConditionalAccessWhatIfReasons_UserRisk              ConditionalAccessWhatIfReasons = "userRisk"
	ConditionalAccessWhatIfReasons_Users                 ConditionalAccessWhatIfReasons = "users"
	ConditionalAccessWhatIfReasons_WorkloadIdentities    ConditionalAccessWhatIfReasons = "workloadIdentities"
)

func PossibleValuesForConditionalAccessWhatIfReasons() []string {
	return []string{
		string(ConditionalAccessWhatIfReasons_Application),
		string(ConditionalAccessWhatIfReasons_AuthenticationContext),
		string(ConditionalAccessWhatIfReasons_AuthenticationFlow),
		string(ConditionalAccessWhatIfReasons_ClientApps),
		string(ConditionalAccessWhatIfReasons_DevicePlatform),
		string(ConditionalAccessWhatIfReasons_Devices),
		string(ConditionalAccessWhatIfReasons_EmptyPolicy),
		string(ConditionalAccessWhatIfReasons_InsiderRisk),
		string(ConditionalAccessWhatIfReasons_InvalidCondition),
		string(ConditionalAccessWhatIfReasons_InvalidPolicy),
		string(ConditionalAccessWhatIfReasons_Location),
		string(ConditionalAccessWhatIfReasons_NotEnoughInformation),
		string(ConditionalAccessWhatIfReasons_NotSet),
		string(ConditionalAccessWhatIfReasons_PolicyNotEnabled),
		string(ConditionalAccessWhatIfReasons_SignInRisk),
		string(ConditionalAccessWhatIfReasons_Time),
		string(ConditionalAccessWhatIfReasons_UserActions),
		string(ConditionalAccessWhatIfReasons_UserRisk),
		string(ConditionalAccessWhatIfReasons_Users),
		string(ConditionalAccessWhatIfReasons_WorkloadIdentities),
	}
}

func (s *ConditionalAccessWhatIfReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessWhatIfReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessWhatIfReasons(input string) (*ConditionalAccessWhatIfReasons, error) {
	vals := map[string]ConditionalAccessWhatIfReasons{
		"application":           ConditionalAccessWhatIfReasons_Application,
		"authenticationcontext": ConditionalAccessWhatIfReasons_AuthenticationContext,
		"authenticationflow":    ConditionalAccessWhatIfReasons_AuthenticationFlow,
		"clientapps":            ConditionalAccessWhatIfReasons_ClientApps,
		"deviceplatform":        ConditionalAccessWhatIfReasons_DevicePlatform,
		"devices":               ConditionalAccessWhatIfReasons_Devices,
		"emptypolicy":           ConditionalAccessWhatIfReasons_EmptyPolicy,
		"insiderrisk":           ConditionalAccessWhatIfReasons_InsiderRisk,
		"invalidcondition":      ConditionalAccessWhatIfReasons_InvalidCondition,
		"invalidpolicy":         ConditionalAccessWhatIfReasons_InvalidPolicy,
		"location":              ConditionalAccessWhatIfReasons_Location,
		"notenoughinformation":  ConditionalAccessWhatIfReasons_NotEnoughInformation,
		"notset":                ConditionalAccessWhatIfReasons_NotSet,
		"policynotenabled":      ConditionalAccessWhatIfReasons_PolicyNotEnabled,
		"signinrisk":            ConditionalAccessWhatIfReasons_SignInRisk,
		"time":                  ConditionalAccessWhatIfReasons_Time,
		"useractions":           ConditionalAccessWhatIfReasons_UserActions,
		"userrisk":              ConditionalAccessWhatIfReasons_UserRisk,
		"users":                 ConditionalAccessWhatIfReasons_Users,
		"workloadidentities":    ConditionalAccessWhatIfReasons_WorkloadIdentities,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessWhatIfReasons(input)
	return &out, nil
}
