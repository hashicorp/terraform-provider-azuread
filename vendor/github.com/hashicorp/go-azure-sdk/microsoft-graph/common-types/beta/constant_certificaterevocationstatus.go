package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateRevocationStatus string

const (
	CertificateRevocationStatus_Failed  CertificateRevocationStatus = "failed"
	CertificateRevocationStatus_Issued  CertificateRevocationStatus = "issued"
	CertificateRevocationStatus_None    CertificateRevocationStatus = "none"
	CertificateRevocationStatus_Pending CertificateRevocationStatus = "pending"
	CertificateRevocationStatus_Revoked CertificateRevocationStatus = "revoked"
)

func PossibleValuesForCertificateRevocationStatus() []string {
	return []string{
		string(CertificateRevocationStatus_Failed),
		string(CertificateRevocationStatus_Issued),
		string(CertificateRevocationStatus_None),
		string(CertificateRevocationStatus_Pending),
		string(CertificateRevocationStatus_Revoked),
	}
}

func (s *CertificateRevocationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateRevocationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateRevocationStatus(input string) (*CertificateRevocationStatus, error) {
	vals := map[string]CertificateRevocationStatus{
		"failed":  CertificateRevocationStatus_Failed,
		"issued":  CertificateRevocationStatus_Issued,
		"none":    CertificateRevocationStatus_None,
		"pending": CertificateRevocationStatus_Pending,
		"revoked": CertificateRevocationStatus_Revoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateRevocationStatus(input)
	return &out, nil
}
