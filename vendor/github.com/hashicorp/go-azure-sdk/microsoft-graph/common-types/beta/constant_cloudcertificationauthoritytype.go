package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityType string

const (
	CloudCertificationAuthorityType_IssuingCertificationAuthority                 CloudCertificationAuthorityType = "issuingCertificationAuthority"
	CloudCertificationAuthorityType_IssuingCertificationAuthorityWithExternalRoot CloudCertificationAuthorityType = "issuingCertificationAuthorityWithExternalRoot"
	CloudCertificationAuthorityType_RootCertificationAuthority                    CloudCertificationAuthorityType = "rootCertificationAuthority"
	CloudCertificationAuthorityType_Unknown                                       CloudCertificationAuthorityType = "unknown"
)

func PossibleValuesForCloudCertificationAuthorityType() []string {
	return []string{
		string(CloudCertificationAuthorityType_IssuingCertificationAuthority),
		string(CloudCertificationAuthorityType_IssuingCertificationAuthorityWithExternalRoot),
		string(CloudCertificationAuthorityType_RootCertificationAuthority),
		string(CloudCertificationAuthorityType_Unknown),
	}
}

func (s *CloudCertificationAuthorityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudCertificationAuthorityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudCertificationAuthorityType(input string) (*CloudCertificationAuthorityType, error) {
	vals := map[string]CloudCertificationAuthorityType{
		"issuingcertificationauthority":                 CloudCertificationAuthorityType_IssuingCertificationAuthority,
		"issuingcertificationauthoritywithexternalroot": CloudCertificationAuthorityType_IssuingCertificationAuthorityWithExternalRoot,
		"rootcertificationauthority":                    CloudCertificationAuthorityType_RootCertificationAuthority,
		"unknown":                                       CloudCertificationAuthorityType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudCertificationAuthorityType(input)
	return &out, nil
}
