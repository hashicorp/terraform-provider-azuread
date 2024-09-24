package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleSubjectNameFormat string

const (
	AppleSubjectNameFormat_CommonName               AppleSubjectNameFormat = "commonName"
	AppleSubjectNameFormat_CommonNameAsEmail        AppleSubjectNameFormat = "commonNameAsEmail"
	AppleSubjectNameFormat_CommonNameAsIMEI         AppleSubjectNameFormat = "commonNameAsIMEI"
	AppleSubjectNameFormat_CommonNameAsSerialNumber AppleSubjectNameFormat = "commonNameAsSerialNumber"
	AppleSubjectNameFormat_CommonNameIncludingEmail AppleSubjectNameFormat = "commonNameIncludingEmail"
	AppleSubjectNameFormat_Custom                   AppleSubjectNameFormat = "custom"
)

func PossibleValuesForAppleSubjectNameFormat() []string {
	return []string{
		string(AppleSubjectNameFormat_CommonName),
		string(AppleSubjectNameFormat_CommonNameAsEmail),
		string(AppleSubjectNameFormat_CommonNameAsIMEI),
		string(AppleSubjectNameFormat_CommonNameAsSerialNumber),
		string(AppleSubjectNameFormat_CommonNameIncludingEmail),
		string(AppleSubjectNameFormat_Custom),
	}
}

func (s *AppleSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleSubjectNameFormat(input string) (*AppleSubjectNameFormat, error) {
	vals := map[string]AppleSubjectNameFormat{
		"commonname":               AppleSubjectNameFormat_CommonName,
		"commonnameasemail":        AppleSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         AppleSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": AppleSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": AppleSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   AppleSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleSubjectNameFormat(input)
	return &out, nil
}
