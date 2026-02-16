package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityType string

const (
	CertificateAuthorityType_Intermediate CertificateAuthorityType = "intermediate"
	CertificateAuthorityType_Root         CertificateAuthorityType = "root"
)

func PossibleValuesForCertificateAuthorityType() []string {
	return []string{
		string(CertificateAuthorityType_Intermediate),
		string(CertificateAuthorityType_Root),
	}
}

func (s *CertificateAuthorityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateAuthorityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateAuthorityType(input string) (*CertificateAuthorityType, error) {
	vals := map[string]CertificateAuthorityType{
		"intermediate": CertificateAuthorityType_Intermediate,
		"root":         CertificateAuthorityType_Root,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateAuthorityType(input)
	return &out, nil
}
