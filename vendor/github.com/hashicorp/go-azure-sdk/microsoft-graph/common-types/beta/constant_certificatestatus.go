package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateStatus string

const (
	CertificateStatus_NotProvisioned CertificateStatus = "notProvisioned"
	CertificateStatus_Provisioned    CertificateStatus = "provisioned"
)

func PossibleValuesForCertificateStatus() []string {
	return []string{
		string(CertificateStatus_NotProvisioned),
		string(CertificateStatus_Provisioned),
	}
}

func (s *CertificateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateStatus(input string) (*CertificateStatus, error) {
	vals := map[string]CertificateStatus{
		"notprovisioned": CertificateStatus_NotProvisioned,
		"provisioned":    CertificateStatus_Provisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateStatus(input)
	return &out, nil
}
