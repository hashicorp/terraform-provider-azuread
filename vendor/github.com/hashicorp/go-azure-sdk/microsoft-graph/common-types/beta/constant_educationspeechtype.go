package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSpeechType string

const (
	EducationSpeechType_Informative EducationSpeechType = "informative"
	EducationSpeechType_Personal    EducationSpeechType = "personal"
	EducationSpeechType_Persuasive  EducationSpeechType = "persuasive"
)

func PossibleValuesForEducationSpeechType() []string {
	return []string{
		string(EducationSpeechType_Informative),
		string(EducationSpeechType_Personal),
		string(EducationSpeechType_Persuasive),
	}
}

func (s *EducationSpeechType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationSpeechType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationSpeechType(input string) (*EducationSpeechType, error) {
	vals := map[string]EducationSpeechType{
		"informative": EducationSpeechType_Informative,
		"personal":    EducationSpeechType_Personal,
		"persuasive":  EducationSpeechType_Persuasive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSpeechType(input)
	return &out, nil
}
