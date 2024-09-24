package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceClassificationType string

const (
	InferenceClassificationType_Focused InferenceClassificationType = "focused"
	InferenceClassificationType_Other   InferenceClassificationType = "other"
)

func PossibleValuesForInferenceClassificationType() []string {
	return []string{
		string(InferenceClassificationType_Focused),
		string(InferenceClassificationType_Other),
	}
}

func (s *InferenceClassificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInferenceClassificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInferenceClassificationType(input string) (*InferenceClassificationType, error) {
	vals := map[string]InferenceClassificationType{
		"focused": InferenceClassificationType_Focused,
		"other":   InferenceClassificationType_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InferenceClassificationType(input)
	return &out, nil
}
