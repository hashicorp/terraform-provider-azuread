package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentFormat string

const (
	SecurityContentFormat_Html     SecurityContentFormat = "html"
	SecurityContentFormat_Markdown SecurityContentFormat = "markdown"
	SecurityContentFormat_Text     SecurityContentFormat = "text"
)

func PossibleValuesForSecurityContentFormat() []string {
	return []string{
		string(SecurityContentFormat_Html),
		string(SecurityContentFormat_Markdown),
		string(SecurityContentFormat_Text),
	}
}

func (s *SecurityContentFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContentFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContentFormat(input string) (*SecurityContentFormat, error) {
	vals := map[string]SecurityContentFormat{
		"html":     SecurityContentFormat_Html,
		"markdown": SecurityContentFormat_Markdown,
		"text":     SecurityContentFormat_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContentFormat(input)
	return &out, nil
}
