package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestionAnswerInputType string

const (
	VirtualEventRegistrationQuestionAnswerInputType_Boolean       VirtualEventRegistrationQuestionAnswerInputType = "boolean"
	VirtualEventRegistrationQuestionAnswerInputType_MultiChoice   VirtualEventRegistrationQuestionAnswerInputType = "multiChoice"
	VirtualEventRegistrationQuestionAnswerInputType_MultilineText VirtualEventRegistrationQuestionAnswerInputType = "multilineText"
	VirtualEventRegistrationQuestionAnswerInputType_SingleChoice  VirtualEventRegistrationQuestionAnswerInputType = "singleChoice"
	VirtualEventRegistrationQuestionAnswerInputType_Text          VirtualEventRegistrationQuestionAnswerInputType = "text"
)

func PossibleValuesForVirtualEventRegistrationQuestionAnswerInputType() []string {
	return []string{
		string(VirtualEventRegistrationQuestionAnswerInputType_Boolean),
		string(VirtualEventRegistrationQuestionAnswerInputType_MultiChoice),
		string(VirtualEventRegistrationQuestionAnswerInputType_MultilineText),
		string(VirtualEventRegistrationQuestionAnswerInputType_SingleChoice),
		string(VirtualEventRegistrationQuestionAnswerInputType_Text),
	}
}

func (s *VirtualEventRegistrationQuestionAnswerInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventRegistrationQuestionAnswerInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventRegistrationQuestionAnswerInputType(input string) (*VirtualEventRegistrationQuestionAnswerInputType, error) {
	vals := map[string]VirtualEventRegistrationQuestionAnswerInputType{
		"boolean":       VirtualEventRegistrationQuestionAnswerInputType_Boolean,
		"multichoice":   VirtualEventRegistrationQuestionAnswerInputType_MultiChoice,
		"multilinetext": VirtualEventRegistrationQuestionAnswerInputType_MultilineText,
		"singlechoice":  VirtualEventRegistrationQuestionAnswerInputType_SingleChoice,
		"text":          VirtualEventRegistrationQuestionAnswerInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventRegistrationQuestionAnswerInputType(input)
	return &out, nil
}
