package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailCertificateType string

const (
	EmailCertificateType_Certificate       EmailCertificateType = "certificate"
	EmailCertificateType_DerivedCredential EmailCertificateType = "derivedCredential"
	EmailCertificateType_None              EmailCertificateType = "none"
)

func PossibleValuesForEmailCertificateType() []string {
	return []string{
		string(EmailCertificateType_Certificate),
		string(EmailCertificateType_DerivedCredential),
		string(EmailCertificateType_None),
	}
}

func (s *EmailCertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailCertificateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailCertificateType(input string) (*EmailCertificateType, error) {
	vals := map[string]EmailCertificateType{
		"certificate":       EmailCertificateType_Certificate,
		"derivedcredential": EmailCertificateType_DerivedCredential,
		"none":              EmailCertificateType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailCertificateType(input)
	return &out, nil
}
