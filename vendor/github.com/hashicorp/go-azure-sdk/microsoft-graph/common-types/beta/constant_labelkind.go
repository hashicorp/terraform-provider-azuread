package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelKind string

const (
	LabelKind_All        LabelKind = "all"
	LabelKind_Enumerated LabelKind = "enumerated"
)

func PossibleValuesForLabelKind() []string {
	return []string{
		string(LabelKind_All),
		string(LabelKind_Enumerated),
	}
}

func (s *LabelKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLabelKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLabelKind(input string) (*LabelKind, error) {
	vals := map[string]LabelKind{
		"all":        LabelKind_All,
		"enumerated": LabelKind_Enumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LabelKind(input)
	return &out, nil
}
