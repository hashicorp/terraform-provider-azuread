package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleAreaTextAlignmentType string

const (
	TitleAreaTextAlignmentType_Center TitleAreaTextAlignmentType = "center"
	TitleAreaTextAlignmentType_Left   TitleAreaTextAlignmentType = "left"
)

func PossibleValuesForTitleAreaTextAlignmentType() []string {
	return []string{
		string(TitleAreaTextAlignmentType_Center),
		string(TitleAreaTextAlignmentType_Left),
	}
}

func (s *TitleAreaTextAlignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTitleAreaTextAlignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTitleAreaTextAlignmentType(input string) (*TitleAreaTextAlignmentType, error) {
	vals := map[string]TitleAreaTextAlignmentType{
		"center": TitleAreaTextAlignmentType_Center,
		"left":   TitleAreaTextAlignmentType_Left,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TitleAreaTextAlignmentType(input)
	return &out, nil
}
