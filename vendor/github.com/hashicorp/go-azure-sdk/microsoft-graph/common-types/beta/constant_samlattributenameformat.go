package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlAttributeNameFormat string

const (
	SamlAttributeNameFormat_Basic       SamlAttributeNameFormat = "basic"
	SamlAttributeNameFormat_Unspecified SamlAttributeNameFormat = "unspecified"
	SamlAttributeNameFormat_Uri         SamlAttributeNameFormat = "uri"
)

func PossibleValuesForSamlAttributeNameFormat() []string {
	return []string{
		string(SamlAttributeNameFormat_Basic),
		string(SamlAttributeNameFormat_Unspecified),
		string(SamlAttributeNameFormat_Uri),
	}
}

func (s *SamlAttributeNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSamlAttributeNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSamlAttributeNameFormat(input string) (*SamlAttributeNameFormat, error) {
	vals := map[string]SamlAttributeNameFormat{
		"basic":       SamlAttributeNameFormat_Basic,
		"unspecified": SamlAttributeNameFormat_Unspecified,
		"uri":         SamlAttributeNameFormat_Uri,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamlAttributeNameFormat(input)
	return &out, nil
}
