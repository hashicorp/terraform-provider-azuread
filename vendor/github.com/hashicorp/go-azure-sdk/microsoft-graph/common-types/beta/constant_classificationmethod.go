package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationMethod string

const (
	ClassificationMethod_ExactDataMatch  ClassificationMethod = "exactDataMatch"
	ClassificationMethod_Fingerprint     ClassificationMethod = "fingerprint"
	ClassificationMethod_MachineLearning ClassificationMethod = "machineLearning"
	ClassificationMethod_PatternMatch    ClassificationMethod = "patternMatch"
)

func PossibleValuesForClassificationMethod() []string {
	return []string{
		string(ClassificationMethod_ExactDataMatch),
		string(ClassificationMethod_Fingerprint),
		string(ClassificationMethod_MachineLearning),
		string(ClassificationMethod_PatternMatch),
	}
}

func (s *ClassificationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClassificationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClassificationMethod(input string) (*ClassificationMethod, error) {
	vals := map[string]ClassificationMethod{
		"exactdatamatch":  ClassificationMethod_ExactDataMatch,
		"fingerprint":     ClassificationMethod_Fingerprint,
		"machinelearning": ClassificationMethod_MachineLearning,
		"patternmatch":    ClassificationMethod_PatternMatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClassificationMethod(input)
	return &out, nil
}
