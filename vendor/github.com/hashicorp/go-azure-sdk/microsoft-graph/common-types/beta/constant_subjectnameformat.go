package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectNameFormat string

const (
	SubjectNameFormat_CommonName                  SubjectNameFormat = "commonName"
	SubjectNameFormat_CommonNameAsAadDeviceId     SubjectNameFormat = "commonNameAsAadDeviceId"
	SubjectNameFormat_CommonNameAsDurableDeviceId SubjectNameFormat = "commonNameAsDurableDeviceId"
	SubjectNameFormat_CommonNameAsEmail           SubjectNameFormat = "commonNameAsEmail"
	SubjectNameFormat_CommonNameAsIMEI            SubjectNameFormat = "commonNameAsIMEI"
	SubjectNameFormat_CommonNameAsIntuneDeviceId  SubjectNameFormat = "commonNameAsIntuneDeviceId"
	SubjectNameFormat_CommonNameAsSerialNumber    SubjectNameFormat = "commonNameAsSerialNumber"
	SubjectNameFormat_CommonNameIncludingEmail    SubjectNameFormat = "commonNameIncludingEmail"
	SubjectNameFormat_Custom                      SubjectNameFormat = "custom"
)

func PossibleValuesForSubjectNameFormat() []string {
	return []string{
		string(SubjectNameFormat_CommonName),
		string(SubjectNameFormat_CommonNameAsAadDeviceId),
		string(SubjectNameFormat_CommonNameAsDurableDeviceId),
		string(SubjectNameFormat_CommonNameAsEmail),
		string(SubjectNameFormat_CommonNameAsIMEI),
		string(SubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(SubjectNameFormat_CommonNameAsSerialNumber),
		string(SubjectNameFormat_CommonNameIncludingEmail),
		string(SubjectNameFormat_Custom),
	}
}

func (s *SubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectNameFormat(input string) (*SubjectNameFormat, error) {
	vals := map[string]SubjectNameFormat{
		"commonname":                  SubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     SubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": SubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           SubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            SubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  SubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    SubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    SubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      SubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectNameFormat(input)
	return &out, nil
}
