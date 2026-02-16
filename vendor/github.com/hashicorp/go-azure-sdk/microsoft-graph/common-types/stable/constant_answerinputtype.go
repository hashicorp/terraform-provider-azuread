package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnswerInputType string

const (
	AnswerInputType_RadioButton AnswerInputType = "radioButton"
	AnswerInputType_Text        AnswerInputType = "text"
)

func PossibleValuesForAnswerInputType() []string {
	return []string{
		string(AnswerInputType_RadioButton),
		string(AnswerInputType_Text),
	}
}

func (s *AnswerInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAnswerInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAnswerInputType(input string) (*AnswerInputType, error) {
	vals := map[string]AnswerInputType{
		"radiobutton": AnswerInputType_RadioButton,
		"text":        AnswerInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AnswerInputType(input)
	return &out, nil
}
