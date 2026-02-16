package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateIssuerHintsState string

const (
	X509CertificateIssuerHintsState_Disabled X509CertificateIssuerHintsState = "disabled"
	X509CertificateIssuerHintsState_Enabled  X509CertificateIssuerHintsState = "enabled"
)

func PossibleValuesForX509CertificateIssuerHintsState() []string {
	return []string{
		string(X509CertificateIssuerHintsState_Disabled),
		string(X509CertificateIssuerHintsState_Enabled),
	}
}

func (s *X509CertificateIssuerHintsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateIssuerHintsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateIssuerHintsState(input string) (*X509CertificateIssuerHintsState, error) {
	vals := map[string]X509CertificateIssuerHintsState{
		"disabled": X509CertificateIssuerHintsState_Disabled,
		"enabled":  X509CertificateIssuerHintsState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateIssuerHintsState(input)
	return &out, nil
}
