package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CategoryColor string

const (
	CategoryColor_None     CategoryColor = "none"
	CategoryColor_Preset0  CategoryColor = "preset0"
	CategoryColor_Preset1  CategoryColor = "preset1"
	CategoryColor_Preset10 CategoryColor = "preset10"
	CategoryColor_Preset11 CategoryColor = "preset11"
	CategoryColor_Preset12 CategoryColor = "preset12"
	CategoryColor_Preset13 CategoryColor = "preset13"
	CategoryColor_Preset14 CategoryColor = "preset14"
	CategoryColor_Preset15 CategoryColor = "preset15"
	CategoryColor_Preset16 CategoryColor = "preset16"
	CategoryColor_Preset17 CategoryColor = "preset17"
	CategoryColor_Preset18 CategoryColor = "preset18"
	CategoryColor_Preset19 CategoryColor = "preset19"
	CategoryColor_Preset2  CategoryColor = "preset2"
	CategoryColor_Preset20 CategoryColor = "preset20"
	CategoryColor_Preset21 CategoryColor = "preset21"
	CategoryColor_Preset22 CategoryColor = "preset22"
	CategoryColor_Preset23 CategoryColor = "preset23"
	CategoryColor_Preset24 CategoryColor = "preset24"
	CategoryColor_Preset3  CategoryColor = "preset3"
	CategoryColor_Preset4  CategoryColor = "preset4"
	CategoryColor_Preset5  CategoryColor = "preset5"
	CategoryColor_Preset6  CategoryColor = "preset6"
	CategoryColor_Preset7  CategoryColor = "preset7"
	CategoryColor_Preset8  CategoryColor = "preset8"
	CategoryColor_Preset9  CategoryColor = "preset9"
)

func PossibleValuesForCategoryColor() []string {
	return []string{
		string(CategoryColor_None),
		string(CategoryColor_Preset0),
		string(CategoryColor_Preset1),
		string(CategoryColor_Preset10),
		string(CategoryColor_Preset11),
		string(CategoryColor_Preset12),
		string(CategoryColor_Preset13),
		string(CategoryColor_Preset14),
		string(CategoryColor_Preset15),
		string(CategoryColor_Preset16),
		string(CategoryColor_Preset17),
		string(CategoryColor_Preset18),
		string(CategoryColor_Preset19),
		string(CategoryColor_Preset2),
		string(CategoryColor_Preset20),
		string(CategoryColor_Preset21),
		string(CategoryColor_Preset22),
		string(CategoryColor_Preset23),
		string(CategoryColor_Preset24),
		string(CategoryColor_Preset3),
		string(CategoryColor_Preset4),
		string(CategoryColor_Preset5),
		string(CategoryColor_Preset6),
		string(CategoryColor_Preset7),
		string(CategoryColor_Preset8),
		string(CategoryColor_Preset9),
	}
}

func (s *CategoryColor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategoryColor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategoryColor(input string) (*CategoryColor, error) {
	vals := map[string]CategoryColor{
		"none":     CategoryColor_None,
		"preset0":  CategoryColor_Preset0,
		"preset1":  CategoryColor_Preset1,
		"preset10": CategoryColor_Preset10,
		"preset11": CategoryColor_Preset11,
		"preset12": CategoryColor_Preset12,
		"preset13": CategoryColor_Preset13,
		"preset14": CategoryColor_Preset14,
		"preset15": CategoryColor_Preset15,
		"preset16": CategoryColor_Preset16,
		"preset17": CategoryColor_Preset17,
		"preset18": CategoryColor_Preset18,
		"preset19": CategoryColor_Preset19,
		"preset2":  CategoryColor_Preset2,
		"preset20": CategoryColor_Preset20,
		"preset21": CategoryColor_Preset21,
		"preset22": CategoryColor_Preset22,
		"preset23": CategoryColor_Preset23,
		"preset24": CategoryColor_Preset24,
		"preset3":  CategoryColor_Preset3,
		"preset4":  CategoryColor_Preset4,
		"preset5":  CategoryColor_Preset5,
		"preset6":  CategoryColor_Preset6,
		"preset7":  CategoryColor_Preset7,
		"preset8":  CategoryColor_Preset8,
		"preset9":  CategoryColor_Preset9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CategoryColor(input)
	return &out, nil
}
