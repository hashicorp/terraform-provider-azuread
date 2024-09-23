package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthStatus string

const (
	ServiceHealthStatus_Confirmed                   ServiceHealthStatus = "confirmed"
	ServiceHealthStatus_ExtendedRecovery            ServiceHealthStatus = "extendedRecovery"
	ServiceHealthStatus_FalsePositive               ServiceHealthStatus = "falsePositive"
	ServiceHealthStatus_Investigating               ServiceHealthStatus = "investigating"
	ServiceHealthStatus_InvestigationSuspended      ServiceHealthStatus = "investigationSuspended"
	ServiceHealthStatus_Mitigated                   ServiceHealthStatus = "mitigated"
	ServiceHealthStatus_MitigatedExternal           ServiceHealthStatus = "mitigatedExternal"
	ServiceHealthStatus_PostIncidentReviewPublished ServiceHealthStatus = "postIncidentReviewPublished"
	ServiceHealthStatus_Reported                    ServiceHealthStatus = "reported"
	ServiceHealthStatus_Resolved                    ServiceHealthStatus = "resolved"
	ServiceHealthStatus_ResolvedExternal            ServiceHealthStatus = "resolvedExternal"
	ServiceHealthStatus_RestoringService            ServiceHealthStatus = "restoringService"
	ServiceHealthStatus_ServiceDegradation          ServiceHealthStatus = "serviceDegradation"
	ServiceHealthStatus_ServiceInterruption         ServiceHealthStatus = "serviceInterruption"
	ServiceHealthStatus_ServiceOperational          ServiceHealthStatus = "serviceOperational"
	ServiceHealthStatus_ServiceRestored             ServiceHealthStatus = "serviceRestored"
	ServiceHealthStatus_VerifyingService            ServiceHealthStatus = "verifyingService"
)

func PossibleValuesForServiceHealthStatus() []string {
	return []string{
		string(ServiceHealthStatus_Confirmed),
		string(ServiceHealthStatus_ExtendedRecovery),
		string(ServiceHealthStatus_FalsePositive),
		string(ServiceHealthStatus_Investigating),
		string(ServiceHealthStatus_InvestigationSuspended),
		string(ServiceHealthStatus_Mitigated),
		string(ServiceHealthStatus_MitigatedExternal),
		string(ServiceHealthStatus_PostIncidentReviewPublished),
		string(ServiceHealthStatus_Reported),
		string(ServiceHealthStatus_Resolved),
		string(ServiceHealthStatus_ResolvedExternal),
		string(ServiceHealthStatus_RestoringService),
		string(ServiceHealthStatus_ServiceDegradation),
		string(ServiceHealthStatus_ServiceInterruption),
		string(ServiceHealthStatus_ServiceOperational),
		string(ServiceHealthStatus_ServiceRestored),
		string(ServiceHealthStatus_VerifyingService),
	}
}

func (s *ServiceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceHealthStatus(input string) (*ServiceHealthStatus, error) {
	vals := map[string]ServiceHealthStatus{
		"confirmed":                   ServiceHealthStatus_Confirmed,
		"extendedrecovery":            ServiceHealthStatus_ExtendedRecovery,
		"falsepositive":               ServiceHealthStatus_FalsePositive,
		"investigating":               ServiceHealthStatus_Investigating,
		"investigationsuspended":      ServiceHealthStatus_InvestigationSuspended,
		"mitigated":                   ServiceHealthStatus_Mitigated,
		"mitigatedexternal":           ServiceHealthStatus_MitigatedExternal,
		"postincidentreviewpublished": ServiceHealthStatus_PostIncidentReviewPublished,
		"reported":                    ServiceHealthStatus_Reported,
		"resolved":                    ServiceHealthStatus_Resolved,
		"resolvedexternal":            ServiceHealthStatus_ResolvedExternal,
		"restoringservice":            ServiceHealthStatus_RestoringService,
		"servicedegradation":          ServiceHealthStatus_ServiceDegradation,
		"serviceinterruption":         ServiceHealthStatus_ServiceInterruption,
		"serviceoperational":          ServiceHealthStatus_ServiceOperational,
		"servicerestored":             ServiceHealthStatus_ServiceRestored,
		"verifyingservice":            ServiceHealthStatus_VerifyingService,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthStatus(input)
	return &out, nil
}
