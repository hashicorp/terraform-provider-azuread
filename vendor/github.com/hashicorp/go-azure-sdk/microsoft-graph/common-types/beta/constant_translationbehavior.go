package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationBehavior string

const (
	TranslationBehavior_Ask TranslationBehavior = "Ask"
	TranslationBehavior_No  TranslationBehavior = "No"
	TranslationBehavior_Yes TranslationBehavior = "Yes"
)

func PossibleValuesForTranslationBehavior() []string {
	return []string{
		string(TranslationBehavior_Ask),
		string(TranslationBehavior_No),
		string(TranslationBehavior_Yes),
	}
}

func (s *TranslationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTranslationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTranslationBehavior(input string) (*TranslationBehavior, error) {
	vals := map[string]TranslationBehavior{
		"ask": TranslationBehavior_Ask,
		"no":  TranslationBehavior_No,
		"yes": TranslationBehavior_Yes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TranslationBehavior(input)
	return &out, nil
}
