package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityQuestionType string

const (
	SecurityQuestionType_Custom     SecurityQuestionType = "custom"
	SecurityQuestionType_Predefined SecurityQuestionType = "predefined"
)

func PossibleValuesForSecurityQuestionType() []string {
	return []string{
		string(SecurityQuestionType_Custom),
		string(SecurityQuestionType_Predefined),
	}
}

func (s *SecurityQuestionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityQuestionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityQuestionType(input string) (*SecurityQuestionType, error) {
	vals := map[string]SecurityQuestionType{
		"custom":     SecurityQuestionType_Custom,
		"predefined": SecurityQuestionType_Predefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityQuestionType(input)
	return &out, nil
}
