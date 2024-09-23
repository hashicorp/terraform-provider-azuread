package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionCategory string

const (
	SecuritySubmissionCategory_Malware  SecuritySubmissionCategory = "malware"
	SecuritySubmissionCategory_NotJunk  SecuritySubmissionCategory = "notJunk"
	SecuritySubmissionCategory_Phishing SecuritySubmissionCategory = "phishing"
	SecuritySubmissionCategory_Spam     SecuritySubmissionCategory = "spam"
)

func PossibleValuesForSecuritySubmissionCategory() []string {
	return []string{
		string(SecuritySubmissionCategory_Malware),
		string(SecuritySubmissionCategory_NotJunk),
		string(SecuritySubmissionCategory_Phishing),
		string(SecuritySubmissionCategory_Spam),
	}
}

func (s *SecuritySubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionCategory(input string) (*SecuritySubmissionCategory, error) {
	vals := map[string]SecuritySubmissionCategory{
		"malware":  SecuritySubmissionCategory_Malware,
		"notjunk":  SecuritySubmissionCategory_NotJunk,
		"phishing": SecuritySubmissionCategory_Phishing,
		"spam":     SecuritySubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionCategory(input)
	return &out, nil
}
