package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskEventType string

const (
	RiskEventType_AdminConfirmedUserCompromised                RiskEventType = "adminConfirmedUserCompromised"
	RiskEventType_AnonymizedIPAddress                          RiskEventType = "anonymizedIPAddress"
	RiskEventType_Generic                                      RiskEventType = "generic"
	RiskEventType_InvestigationsThreatIntelligence             RiskEventType = "investigationsThreatIntelligence"
	RiskEventType_InvestigationsThreatIntelligenceSigninLinked RiskEventType = "investigationsThreatIntelligenceSigninLinked"
	RiskEventType_LeakedCredentials                            RiskEventType = "leakedCredentials"
	RiskEventType_MaliciousIPAddress                           RiskEventType = "maliciousIPAddress"
	RiskEventType_MaliciousIPAddressValidCredentialsBlockedIP  RiskEventType = "maliciousIPAddressValidCredentialsBlockedIP"
	RiskEventType_MalwareInfectedIPAddress                     RiskEventType = "malwareInfectedIPAddress"
	RiskEventType_McasImpossibleTravel                         RiskEventType = "mcasImpossibleTravel"
	RiskEventType_McasSuspiciousInboxManipulationRules         RiskEventType = "mcasSuspiciousInboxManipulationRules"
	RiskEventType_SuspiciousIPAddress                          RiskEventType = "suspiciousIPAddress"
	RiskEventType_UnfamiliarFeatures                           RiskEventType = "unfamiliarFeatures"
	RiskEventType_UnlikelyTravel                               RiskEventType = "unlikelyTravel"
)

func PossibleValuesForRiskEventType() []string {
	return []string{
		string(RiskEventType_AdminConfirmedUserCompromised),
		string(RiskEventType_AnonymizedIPAddress),
		string(RiskEventType_Generic),
		string(RiskEventType_InvestigationsThreatIntelligence),
		string(RiskEventType_InvestigationsThreatIntelligenceSigninLinked),
		string(RiskEventType_LeakedCredentials),
		string(RiskEventType_MaliciousIPAddress),
		string(RiskEventType_MaliciousIPAddressValidCredentialsBlockedIP),
		string(RiskEventType_MalwareInfectedIPAddress),
		string(RiskEventType_McasImpossibleTravel),
		string(RiskEventType_McasSuspiciousInboxManipulationRules),
		string(RiskEventType_SuspiciousIPAddress),
		string(RiskEventType_UnfamiliarFeatures),
		string(RiskEventType_UnlikelyTravel),
	}
}

func (s *RiskEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskEventType(input string) (*RiskEventType, error) {
	vals := map[string]RiskEventType{
		"adminconfirmedusercompromised":                RiskEventType_AdminConfirmedUserCompromised,
		"anonymizedipaddress":                          RiskEventType_AnonymizedIPAddress,
		"generic":                                      RiskEventType_Generic,
		"investigationsthreatintelligence":             RiskEventType_InvestigationsThreatIntelligence,
		"investigationsthreatintelligencesigninlinked": RiskEventType_InvestigationsThreatIntelligenceSigninLinked,
		"leakedcredentials":                            RiskEventType_LeakedCredentials,
		"maliciousipaddress":                           RiskEventType_MaliciousIPAddress,
		"maliciousipaddressvalidcredentialsblockedip":  RiskEventType_MaliciousIPAddressValidCredentialsBlockedIP,
		"malwareinfectedipaddress":                     RiskEventType_MalwareInfectedIPAddress,
		"mcasimpossibletravel":                         RiskEventType_McasImpossibleTravel,
		"mcassuspiciousinboxmanipulationrules":         RiskEventType_McasSuspiciousInboxManipulationRules,
		"suspiciousipaddress":                          RiskEventType_SuspiciousIPAddress,
		"unfamiliarfeatures":                           RiskEventType_UnfamiliarFeatures,
		"unlikelytravel":                               RiskEventType_UnlikelyTravel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskEventType(input)
	return &out, nil
}
