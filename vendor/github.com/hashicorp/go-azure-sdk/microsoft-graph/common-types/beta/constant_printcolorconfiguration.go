package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintColorConfiguration string

const (
	PrintColorConfiguration_Auto          PrintColorConfiguration = "auto"
	PrintColorConfiguration_BlackAndWhite PrintColorConfiguration = "blackAndWhite"
	PrintColorConfiguration_Color         PrintColorConfiguration = "color"
	PrintColorConfiguration_Grayscale     PrintColorConfiguration = "grayscale"
)

func PossibleValuesForPrintColorConfiguration() []string {
	return []string{
		string(PrintColorConfiguration_Auto),
		string(PrintColorConfiguration_BlackAndWhite),
		string(PrintColorConfiguration_Color),
		string(PrintColorConfiguration_Grayscale),
	}
}

func (s *PrintColorConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintColorConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintColorConfiguration(input string) (*PrintColorConfiguration, error) {
	vals := map[string]PrintColorConfiguration{
		"auto":          PrintColorConfiguration_Auto,
		"blackandwhite": PrintColorConfiguration_BlackAndWhite,
		"color":         PrintColorConfiguration_Color,
		"grayscale":     PrintColorConfiguration_Grayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintColorConfiguration(input)
	return &out, nil
}
