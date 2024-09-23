package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintMultipageLayout string

const (
	PrintMultipageLayout_ClockwiseFromBottomLeft         PrintMultipageLayout = "clockwiseFromBottomLeft"
	PrintMultipageLayout_ClockwiseFromBottomRight        PrintMultipageLayout = "clockwiseFromBottomRight"
	PrintMultipageLayout_ClockwiseFromTopLeft            PrintMultipageLayout = "clockwiseFromTopLeft"
	PrintMultipageLayout_ClockwiseFromTopRight           PrintMultipageLayout = "clockwiseFromTopRight"
	PrintMultipageLayout_CounterclockwiseFromBottomLeft  PrintMultipageLayout = "counterclockwiseFromBottomLeft"
	PrintMultipageLayout_CounterclockwiseFromBottomRight PrintMultipageLayout = "counterclockwiseFromBottomRight"
	PrintMultipageLayout_CounterclockwiseFromTopLeft     PrintMultipageLayout = "counterclockwiseFromTopLeft"
	PrintMultipageLayout_CounterclockwiseFromTopRight    PrintMultipageLayout = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrintMultipageLayout() []string {
	return []string{
		string(PrintMultipageLayout_ClockwiseFromBottomLeft),
		string(PrintMultipageLayout_ClockwiseFromBottomRight),
		string(PrintMultipageLayout_ClockwiseFromTopLeft),
		string(PrintMultipageLayout_ClockwiseFromTopRight),
		string(PrintMultipageLayout_CounterclockwiseFromBottomLeft),
		string(PrintMultipageLayout_CounterclockwiseFromBottomRight),
		string(PrintMultipageLayout_CounterclockwiseFromTopLeft),
		string(PrintMultipageLayout_CounterclockwiseFromTopRight),
	}
}

func (s *PrintMultipageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintMultipageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintMultipageLayout(input string) (*PrintMultipageLayout, error) {
	vals := map[string]PrintMultipageLayout{
		"clockwisefrombottomleft":         PrintMultipageLayout_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrintMultipageLayout_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrintMultipageLayout_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrintMultipageLayout_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrintMultipageLayout_CounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrintMultipageLayout_CounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrintMultipageLayout_CounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrintMultipageLayout_CounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintMultipageLayout(input)
	return &out, nil
}
