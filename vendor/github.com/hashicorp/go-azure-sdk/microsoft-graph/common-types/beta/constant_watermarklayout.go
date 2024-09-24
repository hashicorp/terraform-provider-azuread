package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WatermarkLayout string

const (
	WatermarkLayout_Diagonal   WatermarkLayout = "diagonal"
	WatermarkLayout_Horizontal WatermarkLayout = "horizontal"
)

func PossibleValuesForWatermarkLayout() []string {
	return []string{
		string(WatermarkLayout_Diagonal),
		string(WatermarkLayout_Horizontal),
	}
}

func (s *WatermarkLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWatermarkLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWatermarkLayout(input string) (*WatermarkLayout, error) {
	vals := map[string]WatermarkLayout{
		"diagonal":   WatermarkLayout_Diagonal,
		"horizontal": WatermarkLayout_Horizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WatermarkLayout(input)
	return &out, nil
}
