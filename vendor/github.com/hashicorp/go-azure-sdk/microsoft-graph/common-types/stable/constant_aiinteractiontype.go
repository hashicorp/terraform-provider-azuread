package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AiInteractionType string

const (
	AiInteractionType_AiResponse AiInteractionType = "aiResponse"
	AiInteractionType_UserPrompt AiInteractionType = "userPrompt"
)

func PossibleValuesForAiInteractionType() []string {
	return []string{
		string(AiInteractionType_AiResponse),
		string(AiInteractionType_UserPrompt),
	}
}

func (s *AiInteractionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAiInteractionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAiInteractionType(input string) (*AiInteractionType, error) {
	vals := map[string]AiInteractionType{
		"airesponse": AiInteractionType_AiResponse,
		"userprompt": AiInteractionType_UserPrompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AiInteractionType(input)
	return &out, nil
}
