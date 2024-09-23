package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateClassification string

const (
	UpdateClassification_Important               UpdateClassification = "important"
	UpdateClassification_None                    UpdateClassification = "none"
	UpdateClassification_RecommendedAndImportant UpdateClassification = "recommendedAndImportant"
	UpdateClassification_UserDefined             UpdateClassification = "userDefined"
)

func PossibleValuesForUpdateClassification() []string {
	return []string{
		string(UpdateClassification_Important),
		string(UpdateClassification_None),
		string(UpdateClassification_RecommendedAndImportant),
		string(UpdateClassification_UserDefined),
	}
}

func (s *UpdateClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateClassification(input string) (*UpdateClassification, error) {
	vals := map[string]UpdateClassification{
		"important":               UpdateClassification_Important,
		"none":                    UpdateClassification_None,
		"recommendedandimportant": UpdateClassification_RecommendedAndImportant,
		"userdefined":             UpdateClassification_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateClassification(input)
	return &out, nil
}
