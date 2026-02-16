package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingCompletionDuration string

const (
	TrainingCompletionDuration_Fortnite TrainingCompletionDuration = "fortnite"
	TrainingCompletionDuration_Month    TrainingCompletionDuration = "month"
	TrainingCompletionDuration_Week     TrainingCompletionDuration = "week"
)

func PossibleValuesForTrainingCompletionDuration() []string {
	return []string{
		string(TrainingCompletionDuration_Fortnite),
		string(TrainingCompletionDuration_Month),
		string(TrainingCompletionDuration_Week),
	}
}

func (s *TrainingCompletionDuration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingCompletionDuration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingCompletionDuration(input string) (*TrainingCompletionDuration, error) {
	vals := map[string]TrainingCompletionDuration{
		"fortnite": TrainingCompletionDuration_Fortnite,
		"month":    TrainingCompletionDuration_Month,
		"week":     TrainingCompletionDuration_Week,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingCompletionDuration(input)
	return &out, nil
}
