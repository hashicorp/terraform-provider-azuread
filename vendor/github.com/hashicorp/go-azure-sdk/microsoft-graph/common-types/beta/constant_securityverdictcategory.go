package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVerdictCategory string

const (
	SecurityVerdictCategory_DecryptionFailed     SecurityVerdictCategory = "decryptionFailed"
	SecurityVerdictCategory_Malware              SecurityVerdictCategory = "malware"
	SecurityVerdictCategory_None                 SecurityVerdictCategory = "none"
	SecurityVerdictCategory_Phish                SecurityVerdictCategory = "phish"
	SecurityVerdictCategory_SiteUnavailable      SecurityVerdictCategory = "siteUnavailable"
	SecurityVerdictCategory_Spam                 SecurityVerdictCategory = "spam"
	SecurityVerdictCategory_Undefined            SecurityVerdictCategory = "undefined"
	SecurityVerdictCategory_UnsupportedFileType  SecurityVerdictCategory = "unsupportedFileType"
	SecurityVerdictCategory_UnsupportedUriScheme SecurityVerdictCategory = "unsupportedUriScheme"
)

func PossibleValuesForSecurityVerdictCategory() []string {
	return []string{
		string(SecurityVerdictCategory_DecryptionFailed),
		string(SecurityVerdictCategory_Malware),
		string(SecurityVerdictCategory_None),
		string(SecurityVerdictCategory_Phish),
		string(SecurityVerdictCategory_SiteUnavailable),
		string(SecurityVerdictCategory_Spam),
		string(SecurityVerdictCategory_Undefined),
		string(SecurityVerdictCategory_UnsupportedFileType),
		string(SecurityVerdictCategory_UnsupportedUriScheme),
	}
}

func (s *SecurityVerdictCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityVerdictCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityVerdictCategory(input string) (*SecurityVerdictCategory, error) {
	vals := map[string]SecurityVerdictCategory{
		"decryptionfailed":     SecurityVerdictCategory_DecryptionFailed,
		"malware":              SecurityVerdictCategory_Malware,
		"none":                 SecurityVerdictCategory_None,
		"phish":                SecurityVerdictCategory_Phish,
		"siteunavailable":      SecurityVerdictCategory_SiteUnavailable,
		"spam":                 SecurityVerdictCategory_Spam,
		"undefined":            SecurityVerdictCategory_Undefined,
		"unsupportedfiletype":  SecurityVerdictCategory_UnsupportedFileType,
		"unsupportedurischeme": SecurityVerdictCategory_UnsupportedUriScheme,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityVerdictCategory(input)
	return &out, nil
}
