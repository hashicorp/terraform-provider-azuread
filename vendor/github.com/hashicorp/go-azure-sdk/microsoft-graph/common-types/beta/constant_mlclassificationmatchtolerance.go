package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MlClassificationMatchTolerance string

const (
	MlClassificationMatchTolerance_Exact MlClassificationMatchTolerance = "exact"
	MlClassificationMatchTolerance_Near  MlClassificationMatchTolerance = "near"
)

func PossibleValuesForMlClassificationMatchTolerance() []string {
	return []string{
		string(MlClassificationMatchTolerance_Exact),
		string(MlClassificationMatchTolerance_Near),
	}
}

func (s *MlClassificationMatchTolerance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMlClassificationMatchTolerance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMlClassificationMatchTolerance(input string) (*MlClassificationMatchTolerance, error) {
	vals := map[string]MlClassificationMatchTolerance{
		"exact": MlClassificationMatchTolerance_Exact,
		"near":  MlClassificationMatchTolerance_Near,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MlClassificationMatchTolerance(input)
	return &out, nil
}
