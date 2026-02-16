package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentContentType string

const (
	ThreatAssessmentContentType_File ThreatAssessmentContentType = "file"
	ThreatAssessmentContentType_Mail ThreatAssessmentContentType = "mail"
	ThreatAssessmentContentType_Url  ThreatAssessmentContentType = "url"
)

func PossibleValuesForThreatAssessmentContentType() []string {
	return []string{
		string(ThreatAssessmentContentType_File),
		string(ThreatAssessmentContentType_Mail),
		string(ThreatAssessmentContentType_Url),
	}
}

func (s *ThreatAssessmentContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentContentType(input string) (*ThreatAssessmentContentType, error) {
	vals := map[string]ThreatAssessmentContentType{
		"file": ThreatAssessmentContentType_File,
		"mail": ThreatAssessmentContentType_Mail,
		"url":  ThreatAssessmentContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentContentType(input)
	return &out, nil
}
