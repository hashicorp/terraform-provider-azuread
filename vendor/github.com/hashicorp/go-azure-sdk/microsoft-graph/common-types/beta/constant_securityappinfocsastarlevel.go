package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoCsaStarLevel string

const (
	SecurityAppInfoCsaStarLevel_Attestation          SecurityAppInfoCsaStarLevel = "attestation"
	SecurityAppInfoCsaStarLevel_CStarAssessment      SecurityAppInfoCsaStarLevel = "cStarAssessment"
	SecurityAppInfoCsaStarLevel_Certification        SecurityAppInfoCsaStarLevel = "certification"
	SecurityAppInfoCsaStarLevel_ContinuousMonitoring SecurityAppInfoCsaStarLevel = "continuousMonitoring"
	SecurityAppInfoCsaStarLevel_SelfAssessment       SecurityAppInfoCsaStarLevel = "selfAssessment"
	SecurityAppInfoCsaStarLevel_Unknown              SecurityAppInfoCsaStarLevel = "unknown"
)

func PossibleValuesForSecurityAppInfoCsaStarLevel() []string {
	return []string{
		string(SecurityAppInfoCsaStarLevel_Attestation),
		string(SecurityAppInfoCsaStarLevel_CStarAssessment),
		string(SecurityAppInfoCsaStarLevel_Certification),
		string(SecurityAppInfoCsaStarLevel_ContinuousMonitoring),
		string(SecurityAppInfoCsaStarLevel_SelfAssessment),
		string(SecurityAppInfoCsaStarLevel_Unknown),
	}
}

func (s *SecurityAppInfoCsaStarLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoCsaStarLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoCsaStarLevel(input string) (*SecurityAppInfoCsaStarLevel, error) {
	vals := map[string]SecurityAppInfoCsaStarLevel{
		"attestation":          SecurityAppInfoCsaStarLevel_Attestation,
		"cstarassessment":      SecurityAppInfoCsaStarLevel_CStarAssessment,
		"certification":        SecurityAppInfoCsaStarLevel_Certification,
		"continuousmonitoring": SecurityAppInfoCsaStarLevel_ContinuousMonitoring,
		"selfassessment":       SecurityAppInfoCsaStarLevel_SelfAssessment,
		"unknown":              SecurityAppInfoCsaStarLevel_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoCsaStarLevel(input)
	return &out, nil
}
