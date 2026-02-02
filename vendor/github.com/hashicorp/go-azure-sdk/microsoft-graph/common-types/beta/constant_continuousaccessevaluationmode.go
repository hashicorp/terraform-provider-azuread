package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinuousAccessEvaluationMode string

const (
	ContinuousAccessEvaluationMode_Disabled          ContinuousAccessEvaluationMode = "disabled"
	ContinuousAccessEvaluationMode_StrictEnforcement ContinuousAccessEvaluationMode = "strictEnforcement"
	ContinuousAccessEvaluationMode_StrictLocation    ContinuousAccessEvaluationMode = "strictLocation"
)

func PossibleValuesForContinuousAccessEvaluationMode() []string {
	return []string{
		string(ContinuousAccessEvaluationMode_Disabled),
		string(ContinuousAccessEvaluationMode_StrictEnforcement),
		string(ContinuousAccessEvaluationMode_StrictLocation),
	}
}

func (s *ContinuousAccessEvaluationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContinuousAccessEvaluationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContinuousAccessEvaluationMode(input string) (*ContinuousAccessEvaluationMode, error) {
	vals := map[string]ContinuousAccessEvaluationMode{
		"disabled":          ContinuousAccessEvaluationMode_Disabled,
		"strictenforcement": ContinuousAccessEvaluationMode_StrictEnforcement,
		"strictlocation":    ContinuousAccessEvaluationMode_StrictLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContinuousAccessEvaluationMode(input)
	return &out, nil
}
