package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingType string

const (
	TrainingType_Phishing TrainingType = "phishing"
	TrainingType_Unknown  TrainingType = "unknown"
)

func PossibleValuesForTrainingType() []string {
	return []string{
		string(TrainingType_Phishing),
		string(TrainingType_Unknown),
	}
}

func (s *TrainingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingType(input string) (*TrainingType, error) {
	vals := map[string]TrainingType{
		"phishing": TrainingType_Phishing,
		"unknown":  TrainingType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingType(input)
	return &out, nil
}
