package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityLeafCertificateStatus string

const (
	CloudCertificationAuthorityLeafCertificateStatus_Active  CloudCertificationAuthorityLeafCertificateStatus = "active"
	CloudCertificationAuthorityLeafCertificateStatus_Expired CloudCertificationAuthorityLeafCertificateStatus = "expired"
	CloudCertificationAuthorityLeafCertificateStatus_Revoked CloudCertificationAuthorityLeafCertificateStatus = "revoked"
	CloudCertificationAuthorityLeafCertificateStatus_Unknown CloudCertificationAuthorityLeafCertificateStatus = "unknown"
)

func PossibleValuesForCloudCertificationAuthorityLeafCertificateStatus() []string {
	return []string{
		string(CloudCertificationAuthorityLeafCertificateStatus_Active),
		string(CloudCertificationAuthorityLeafCertificateStatus_Expired),
		string(CloudCertificationAuthorityLeafCertificateStatus_Revoked),
		string(CloudCertificationAuthorityLeafCertificateStatus_Unknown),
	}
}

func (s *CloudCertificationAuthorityLeafCertificateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudCertificationAuthorityLeafCertificateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudCertificationAuthorityLeafCertificateStatus(input string) (*CloudCertificationAuthorityLeafCertificateStatus, error) {
	vals := map[string]CloudCertificationAuthorityLeafCertificateStatus{
		"active":  CloudCertificationAuthorityLeafCertificateStatus_Active,
		"expired": CloudCertificationAuthorityLeafCertificateStatus_Expired,
		"revoked": CloudCertificationAuthorityLeafCertificateStatus_Revoked,
		"unknown": CloudCertificationAuthorityLeafCertificateStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudCertificationAuthorityLeafCertificateStatus(input)
	return &out, nil
}
