package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAffinityLevel string

const (
	X509CertificateAffinityLevel_High X509CertificateAffinityLevel = "high"
	X509CertificateAffinityLevel_Low  X509CertificateAffinityLevel = "low"
)

func PossibleValuesForX509CertificateAffinityLevel() []string {
	return []string{
		string(X509CertificateAffinityLevel_High),
		string(X509CertificateAffinityLevel_Low),
	}
}

func (s *X509CertificateAffinityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateAffinityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateAffinityLevel(input string) (*X509CertificateAffinityLevel, error) {
	vals := map[string]X509CertificateAffinityLevel{
		"high": X509CertificateAffinityLevel_High,
		"low":  X509CertificateAffinityLevel_Low,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateAffinityLevel(input)
	return &out, nil
}
