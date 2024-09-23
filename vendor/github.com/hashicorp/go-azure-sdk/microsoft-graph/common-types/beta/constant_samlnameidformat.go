package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlNameIDFormat string

const (
	SamlNameIDFormat_Default                    SamlNameIDFormat = "default"
	SamlNameIDFormat_EmailAddress               SamlNameIDFormat = "emailAddress"
	SamlNameIDFormat_Persistent                 SamlNameIDFormat = "persistent"
	SamlNameIDFormat_Unspecified                SamlNameIDFormat = "unspecified"
	SamlNameIDFormat_WindowsDomainQualifiedName SamlNameIDFormat = "windowsDomainQualifiedName"
)

func PossibleValuesForSamlNameIDFormat() []string {
	return []string{
		string(SamlNameIDFormat_Default),
		string(SamlNameIDFormat_EmailAddress),
		string(SamlNameIDFormat_Persistent),
		string(SamlNameIDFormat_Unspecified),
		string(SamlNameIDFormat_WindowsDomainQualifiedName),
	}
}

func (s *SamlNameIDFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSamlNameIDFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSamlNameIDFormat(input string) (*SamlNameIDFormat, error) {
	vals := map[string]SamlNameIDFormat{
		"default":                    SamlNameIDFormat_Default,
		"emailaddress":               SamlNameIDFormat_EmailAddress,
		"persistent":                 SamlNameIDFormat_Persistent,
		"unspecified":                SamlNameIDFormat_Unspecified,
		"windowsdomainqualifiedname": SamlNameIDFormat_WindowsDomainQualifiedName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamlNameIDFormat(input)
	return &out, nil
}
