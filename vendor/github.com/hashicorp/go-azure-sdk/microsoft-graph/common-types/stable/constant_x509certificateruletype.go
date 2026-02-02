package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateRuleType string

const (
	X509CertificateRuleType_IssuerSubject             X509CertificateRuleType = "issuerSubject"
	X509CertificateRuleType_IssuerSubjectAndPolicyOID X509CertificateRuleType = "issuerSubjectAndPolicyOID"
	X509CertificateRuleType_PolicyOID                 X509CertificateRuleType = "policyOID"
)

func PossibleValuesForX509CertificateRuleType() []string {
	return []string{
		string(X509CertificateRuleType_IssuerSubject),
		string(X509CertificateRuleType_IssuerSubjectAndPolicyOID),
		string(X509CertificateRuleType_PolicyOID),
	}
}

func (s *X509CertificateRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateRuleType(input string) (*X509CertificateRuleType, error) {
	vals := map[string]X509CertificateRuleType{
		"issuersubject":             X509CertificateRuleType_IssuerSubject,
		"issuersubjectandpolicyoid": X509CertificateRuleType_IssuerSubjectAndPolicyOID,
		"policyoid":                 X509CertificateRuleType_PolicyOID,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateRuleType(input)
	return &out, nil
}
