package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityStatus string

const (
	CloudCertificationAuthorityStatus_Active         CloudCertificationAuthorityStatus = "active"
	CloudCertificationAuthorityStatus_Paused         CloudCertificationAuthorityStatus = "paused"
	CloudCertificationAuthorityStatus_Revoked        CloudCertificationAuthorityStatus = "revoked"
	CloudCertificationAuthorityStatus_SigningPending CloudCertificationAuthorityStatus = "signingPending"
	CloudCertificationAuthorityStatus_Unknown        CloudCertificationAuthorityStatus = "unknown"
)

func PossibleValuesForCloudCertificationAuthorityStatus() []string {
	return []string{
		string(CloudCertificationAuthorityStatus_Active),
		string(CloudCertificationAuthorityStatus_Paused),
		string(CloudCertificationAuthorityStatus_Revoked),
		string(CloudCertificationAuthorityStatus_SigningPending),
		string(CloudCertificationAuthorityStatus_Unknown),
	}
}

func (s *CloudCertificationAuthorityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudCertificationAuthorityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudCertificationAuthorityStatus(input string) (*CloudCertificationAuthorityStatus, error) {
	vals := map[string]CloudCertificationAuthorityStatus{
		"active":         CloudCertificationAuthorityStatus_Active,
		"paused":         CloudCertificationAuthorityStatus_Paused,
		"revoked":        CloudCertificationAuthorityStatus_Revoked,
		"signingpending": CloudCertificationAuthorityStatus_SigningPending,
		"unknown":        CloudCertificationAuthorityStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudCertificationAuthorityStatus(input)
	return &out, nil
}
