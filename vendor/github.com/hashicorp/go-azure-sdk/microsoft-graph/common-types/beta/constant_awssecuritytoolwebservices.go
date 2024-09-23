package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsSecurityToolWebServices string

const (
	AwsSecurityToolWebServices_CloudTrail  AwsSecurityToolWebServices = "cloudTrail"
	AwsSecurityToolWebServices_Detective   AwsSecurityToolWebServices = "detective"
	AwsSecurityToolWebServices_GuardDuty   AwsSecurityToolWebServices = "guardDuty"
	AwsSecurityToolWebServices_Inspector   AwsSecurityToolWebServices = "inspector"
	AwsSecurityToolWebServices_Macie       AwsSecurityToolWebServices = "macie"
	AwsSecurityToolWebServices_SecurityHub AwsSecurityToolWebServices = "securityHub"
	AwsSecurityToolWebServices_WafShield   AwsSecurityToolWebServices = "wafShield"
)

func PossibleValuesForAwsSecurityToolWebServices() []string {
	return []string{
		string(AwsSecurityToolWebServices_CloudTrail),
		string(AwsSecurityToolWebServices_Detective),
		string(AwsSecurityToolWebServices_GuardDuty),
		string(AwsSecurityToolWebServices_Inspector),
		string(AwsSecurityToolWebServices_Macie),
		string(AwsSecurityToolWebServices_SecurityHub),
		string(AwsSecurityToolWebServices_WafShield),
	}
}

func (s *AwsSecurityToolWebServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsSecurityToolWebServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsSecurityToolWebServices(input string) (*AwsSecurityToolWebServices, error) {
	vals := map[string]AwsSecurityToolWebServices{
		"cloudtrail":  AwsSecurityToolWebServices_CloudTrail,
		"detective":   AwsSecurityToolWebServices_Detective,
		"guardduty":   AwsSecurityToolWebServices_GuardDuty,
		"inspector":   AwsSecurityToolWebServices_Inspector,
		"macie":       AwsSecurityToolWebServices_Macie,
		"securityhub": AwsSecurityToolWebServices_SecurityHub,
		"wafshield":   AwsSecurityToolWebServices_WafShield,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsSecurityToolWebServices(input)
	return &out, nil
}
