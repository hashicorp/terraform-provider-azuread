package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityWatermarkLayout string

const (
	SecurityWatermarkLayout_Diagonal   SecurityWatermarkLayout = "diagonal"
	SecurityWatermarkLayout_Horizontal SecurityWatermarkLayout = "horizontal"
)

func PossibleValuesForSecurityWatermarkLayout() []string {
	return []string{
		string(SecurityWatermarkLayout_Diagonal),
		string(SecurityWatermarkLayout_Horizontal),
	}
}

func (s *SecurityWatermarkLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityWatermarkLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityWatermarkLayout(input string) (*SecurityWatermarkLayout, error) {
	vals := map[string]SecurityWatermarkLayout{
		"diagonal":   SecurityWatermarkLayout_Diagonal,
		"horizontal": SecurityWatermarkLayout_Horizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityWatermarkLayout(input)
	return &out, nil
}
