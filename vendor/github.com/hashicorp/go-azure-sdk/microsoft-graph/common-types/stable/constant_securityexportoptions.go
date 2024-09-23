package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportOptions string

const (
	SecurityExportOptions_OriginalFiles  SecurityExportOptions = "originalFiles"
	SecurityExportOptions_PdfReplacement SecurityExportOptions = "pdfReplacement"
	SecurityExportOptions_Tags           SecurityExportOptions = "tags"
	SecurityExportOptions_Text           SecurityExportOptions = "text"
)

func PossibleValuesForSecurityExportOptions() []string {
	return []string{
		string(SecurityExportOptions_OriginalFiles),
		string(SecurityExportOptions_PdfReplacement),
		string(SecurityExportOptions_Tags),
		string(SecurityExportOptions_Text),
	}
}

func (s *SecurityExportOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityExportOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityExportOptions(input string) (*SecurityExportOptions, error) {
	vals := map[string]SecurityExportOptions{
		"originalfiles":  SecurityExportOptions_OriginalFiles,
		"pdfreplacement": SecurityExportOptions_PdfReplacement,
		"tags":           SecurityExportOptions_Tags,
		"text":           SecurityExportOptions_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityExportOptions(input)
	return &out, nil
}
