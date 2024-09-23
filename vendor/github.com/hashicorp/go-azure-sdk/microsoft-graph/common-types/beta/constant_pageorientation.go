package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PageOrientation string

const (
	PageOrientation_Diagonal   PageOrientation = "diagonal"
	PageOrientation_Horizontal PageOrientation = "horizontal"
)

func PossibleValuesForPageOrientation() []string {
	return []string{
		string(PageOrientation_Diagonal),
		string(PageOrientation_Horizontal),
	}
}

func (s *PageOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePageOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePageOrientation(input string) (*PageOrientation, error) {
	vals := map[string]PageOrientation{
		"diagonal":   PageOrientation_Diagonal,
		"horizontal": PageOrientation_Horizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PageOrientation(input)
	return &out, nil
}
