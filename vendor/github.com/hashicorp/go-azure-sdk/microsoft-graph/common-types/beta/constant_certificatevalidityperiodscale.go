package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateValidityPeriodScale string

const (
	CertificateValidityPeriodScale_Days   CertificateValidityPeriodScale = "days"
	CertificateValidityPeriodScale_Months CertificateValidityPeriodScale = "months"
	CertificateValidityPeriodScale_Years  CertificateValidityPeriodScale = "years"
)

func PossibleValuesForCertificateValidityPeriodScale() []string {
	return []string{
		string(CertificateValidityPeriodScale_Days),
		string(CertificateValidityPeriodScale_Months),
		string(CertificateValidityPeriodScale_Years),
	}
}

func (s *CertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateValidityPeriodScale(input string) (*CertificateValidityPeriodScale, error) {
	vals := map[string]CertificateValidityPeriodScale{
		"days":   CertificateValidityPeriodScale_Days,
		"months": CertificateValidityPeriodScale_Months,
		"years":  CertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateValidityPeriodScale(input)
	return &out, nil
}
