package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponseFeedbackType string

const (
	ResponseFeedbackType_Neutral        ResponseFeedbackType = "neutral"
	ResponseFeedbackType_None           ResponseFeedbackType = "none"
	ResponseFeedbackType_NotDetected    ResponseFeedbackType = "notDetected"
	ResponseFeedbackType_Pleasant       ResponseFeedbackType = "pleasant"
	ResponseFeedbackType_Unpleasant     ResponseFeedbackType = "unpleasant"
	ResponseFeedbackType_VeryPleasant   ResponseFeedbackType = "veryPleasant"
	ResponseFeedbackType_VeryUnpleasant ResponseFeedbackType = "veryUnpleasant"
)

func PossibleValuesForResponseFeedbackType() []string {
	return []string{
		string(ResponseFeedbackType_Neutral),
		string(ResponseFeedbackType_None),
		string(ResponseFeedbackType_NotDetected),
		string(ResponseFeedbackType_Pleasant),
		string(ResponseFeedbackType_Unpleasant),
		string(ResponseFeedbackType_VeryPleasant),
		string(ResponseFeedbackType_VeryUnpleasant),
	}
}

func (s *ResponseFeedbackType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResponseFeedbackType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResponseFeedbackType(input string) (*ResponseFeedbackType, error) {
	vals := map[string]ResponseFeedbackType{
		"neutral":        ResponseFeedbackType_Neutral,
		"none":           ResponseFeedbackType_None,
		"notdetected":    ResponseFeedbackType_NotDetected,
		"pleasant":       ResponseFeedbackType_Pleasant,
		"unpleasant":     ResponseFeedbackType_Unpleasant,
		"verypleasant":   ResponseFeedbackType_VeryPleasant,
		"veryunpleasant": ResponseFeedbackType_VeryUnpleasant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResponseFeedbackType(input)
	return &out, nil
}
