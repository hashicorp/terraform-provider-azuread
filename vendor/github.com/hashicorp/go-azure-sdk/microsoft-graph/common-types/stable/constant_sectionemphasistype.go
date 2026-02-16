package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SectionEmphasisType string

const (
	SectionEmphasisType_Neutral SectionEmphasisType = "neutral"
	SectionEmphasisType_None    SectionEmphasisType = "none"
	SectionEmphasisType_Soft    SectionEmphasisType = "soft"
	SectionEmphasisType_Strong  SectionEmphasisType = "strong"
)

func PossibleValuesForSectionEmphasisType() []string {
	return []string{
		string(SectionEmphasisType_Neutral),
		string(SectionEmphasisType_None),
		string(SectionEmphasisType_Soft),
		string(SectionEmphasisType_Strong),
	}
}

func (s *SectionEmphasisType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSectionEmphasisType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSectionEmphasisType(input string) (*SectionEmphasisType, error) {
	vals := map[string]SectionEmphasisType{
		"neutral": SectionEmphasisType_Neutral,
		"none":    SectionEmphasisType_None,
		"soft":    SectionEmphasisType_Soft,
		"strong":  SectionEmphasisType_Strong,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SectionEmphasisType(input)
	return &out, nil
}
