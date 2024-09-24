package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionContentType string

const (
	SecuritySubmissionContentType_App   SecuritySubmissionContentType = "app"
	SecuritySubmissionContentType_Email SecuritySubmissionContentType = "email"
	SecuritySubmissionContentType_File  SecuritySubmissionContentType = "file"
	SecuritySubmissionContentType_Url   SecuritySubmissionContentType = "url"
)

func PossibleValuesForSecuritySubmissionContentType() []string {
	return []string{
		string(SecuritySubmissionContentType_App),
		string(SecuritySubmissionContentType_Email),
		string(SecuritySubmissionContentType_File),
		string(SecuritySubmissionContentType_Url),
	}
}

func (s *SecuritySubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionContentType(input string) (*SecuritySubmissionContentType, error) {
	vals := map[string]SecuritySubmissionContentType{
		"app":   SecuritySubmissionContentType_App,
		"email": SecuritySubmissionContentType_Email,
		"file":  SecuritySubmissionContentType_File,
		"url":   SecuritySubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionContentType(input)
	return &out, nil
}
