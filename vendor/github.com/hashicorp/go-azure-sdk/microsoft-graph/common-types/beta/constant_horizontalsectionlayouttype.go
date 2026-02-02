package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HorizontalSectionLayoutType string

const (
	HorizontalSectionLayoutType_FullWidth           HorizontalSectionLayoutType = "fullWidth"
	HorizontalSectionLayoutType_None                HorizontalSectionLayoutType = "none"
	HorizontalSectionLayoutType_OneColumn           HorizontalSectionLayoutType = "oneColumn"
	HorizontalSectionLayoutType_OneThirdLeftColumn  HorizontalSectionLayoutType = "oneThirdLeftColumn"
	HorizontalSectionLayoutType_OneThirdRightColumn HorizontalSectionLayoutType = "oneThirdRightColumn"
	HorizontalSectionLayoutType_ThreeColumns        HorizontalSectionLayoutType = "threeColumns"
	HorizontalSectionLayoutType_TwoColumns          HorizontalSectionLayoutType = "twoColumns"
)

func PossibleValuesForHorizontalSectionLayoutType() []string {
	return []string{
		string(HorizontalSectionLayoutType_FullWidth),
		string(HorizontalSectionLayoutType_None),
		string(HorizontalSectionLayoutType_OneColumn),
		string(HorizontalSectionLayoutType_OneThirdLeftColumn),
		string(HorizontalSectionLayoutType_OneThirdRightColumn),
		string(HorizontalSectionLayoutType_ThreeColumns),
		string(HorizontalSectionLayoutType_TwoColumns),
	}
}

func (s *HorizontalSectionLayoutType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHorizontalSectionLayoutType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHorizontalSectionLayoutType(input string) (*HorizontalSectionLayoutType, error) {
	vals := map[string]HorizontalSectionLayoutType{
		"fullwidth":           HorizontalSectionLayoutType_FullWidth,
		"none":                HorizontalSectionLayoutType_None,
		"onecolumn":           HorizontalSectionLayoutType_OneColumn,
		"onethirdleftcolumn":  HorizontalSectionLayoutType_OneThirdLeftColumn,
		"onethirdrightcolumn": HorizontalSectionLayoutType_OneThirdRightColumn,
		"threecolumns":        HorizontalSectionLayoutType_ThreeColumns,
		"twocolumns":          HorizontalSectionLayoutType_TwoColumns,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HorizontalSectionLayoutType(input)
	return &out, nil
}
