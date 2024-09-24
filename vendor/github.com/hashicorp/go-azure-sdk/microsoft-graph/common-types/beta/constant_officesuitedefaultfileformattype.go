package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteDefaultFileFormatType string

const (
	OfficeSuiteDefaultFileFormatType_NotConfigured            OfficeSuiteDefaultFileFormatType = "notConfigured"
	OfficeSuiteDefaultFileFormatType_OfficeOpenDocumentFormat OfficeSuiteDefaultFileFormatType = "officeOpenDocumentFormat"
	OfficeSuiteDefaultFileFormatType_OfficeOpenXMLFormat      OfficeSuiteDefaultFileFormatType = "officeOpenXMLFormat"
)

func PossibleValuesForOfficeSuiteDefaultFileFormatType() []string {
	return []string{
		string(OfficeSuiteDefaultFileFormatType_NotConfigured),
		string(OfficeSuiteDefaultFileFormatType_OfficeOpenDocumentFormat),
		string(OfficeSuiteDefaultFileFormatType_OfficeOpenXMLFormat),
	}
}

func (s *OfficeSuiteDefaultFileFormatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteDefaultFileFormatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteDefaultFileFormatType(input string) (*OfficeSuiteDefaultFileFormatType, error) {
	vals := map[string]OfficeSuiteDefaultFileFormatType{
		"notconfigured":            OfficeSuiteDefaultFileFormatType_NotConfigured,
		"officeopendocumentformat": OfficeSuiteDefaultFileFormatType_OfficeOpenDocumentFormat,
		"officeopenxmlformat":      OfficeSuiteDefaultFileFormatType_OfficeOpenXMLFormat,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteDefaultFileFormatType(input)
	return &out, nil
}
