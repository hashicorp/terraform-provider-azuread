package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelectionLikelihoodInfo string

const (
	SelectionLikelihoodInfo_High         SelectionLikelihoodInfo = "high"
	SelectionLikelihoodInfo_NotSpecified SelectionLikelihoodInfo = "notSpecified"
)

func PossibleValuesForSelectionLikelihoodInfo() []string {
	return []string{
		string(SelectionLikelihoodInfo_High),
		string(SelectionLikelihoodInfo_NotSpecified),
	}
}

func (s *SelectionLikelihoodInfo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSelectionLikelihoodInfo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSelectionLikelihoodInfo(input string) (*SelectionLikelihoodInfo, error) {
	vals := map[string]SelectionLikelihoodInfo{
		"high":         SelectionLikelihoodInfo_High,
		"notspecified": SelectionLikelihoodInfo_NotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SelectionLikelihoodInfo(input)
	return &out, nil
}
