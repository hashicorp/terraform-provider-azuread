package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintScaling string

const (
	PrintScaling_Auto        PrintScaling = "auto"
	PrintScaling_Fill        PrintScaling = "fill"
	PrintScaling_Fit         PrintScaling = "fit"
	PrintScaling_None        PrintScaling = "none"
	PrintScaling_ShrinkToFit PrintScaling = "shrinkToFit"
)

func PossibleValuesForPrintScaling() []string {
	return []string{
		string(PrintScaling_Auto),
		string(PrintScaling_Fill),
		string(PrintScaling_Fit),
		string(PrintScaling_None),
		string(PrintScaling_ShrinkToFit),
	}
}

func (s *PrintScaling) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintScaling(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintScaling(input string) (*PrintScaling, error) {
	vals := map[string]PrintScaling{
		"auto":        PrintScaling_Auto,
		"fill":        PrintScaling_Fill,
		"fit":         PrintScaling_Fit,
		"none":        PrintScaling_None,
		"shrinktofit": PrintScaling_ShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintScaling(input)
	return &out, nil
}
