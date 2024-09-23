package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintColorMode string

const (
	PrintColorMode_Auto          PrintColorMode = "auto"
	PrintColorMode_BlackAndWhite PrintColorMode = "blackAndWhite"
	PrintColorMode_Color         PrintColorMode = "color"
	PrintColorMode_Grayscale     PrintColorMode = "grayscale"
)

func PossibleValuesForPrintColorMode() []string {
	return []string{
		string(PrintColorMode_Auto),
		string(PrintColorMode_BlackAndWhite),
		string(PrintColorMode_Color),
		string(PrintColorMode_Grayscale),
	}
}

func (s *PrintColorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintColorMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintColorMode(input string) (*PrintColorMode, error) {
	vals := map[string]PrintColorMode{
		"auto":          PrintColorMode_Auto,
		"blackandwhite": PrintColorMode_BlackAndWhite,
		"color":         PrintColorMode_Color,
		"grayscale":     PrintColorMode_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintColorMode(input)
	return &out, nil
}
