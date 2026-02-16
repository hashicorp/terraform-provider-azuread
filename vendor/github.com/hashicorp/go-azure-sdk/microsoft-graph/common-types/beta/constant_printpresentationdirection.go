package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintPresentationDirection string

const (
	PrintPresentationDirection_ClockwiseFromBottomLeft         PrintPresentationDirection = "clockwiseFromBottomLeft"
	PrintPresentationDirection_ClockwiseFromBottomRight        PrintPresentationDirection = "clockwiseFromBottomRight"
	PrintPresentationDirection_ClockwiseFromTopLeft            PrintPresentationDirection = "clockwiseFromTopLeft"
	PrintPresentationDirection_ClockwiseFromTopRight           PrintPresentationDirection = "clockwiseFromTopRight"
	PrintPresentationDirection_CounterClockwiseFromBottomLeft  PrintPresentationDirection = "counterClockwiseFromBottomLeft"
	PrintPresentationDirection_CounterClockwiseFromBottomRight PrintPresentationDirection = "counterClockwiseFromBottomRight"
	PrintPresentationDirection_CounterClockwiseFromTopLeft     PrintPresentationDirection = "counterClockwiseFromTopLeft"
	PrintPresentationDirection_CounterClockwiseFromTopRight    PrintPresentationDirection = "counterClockwiseFromTopRight"
)

func PossibleValuesForPrintPresentationDirection() []string {
	return []string{
		string(PrintPresentationDirection_ClockwiseFromBottomLeft),
		string(PrintPresentationDirection_ClockwiseFromBottomRight),
		string(PrintPresentationDirection_ClockwiseFromTopLeft),
		string(PrintPresentationDirection_ClockwiseFromTopRight),
		string(PrintPresentationDirection_CounterClockwiseFromBottomLeft),
		string(PrintPresentationDirection_CounterClockwiseFromBottomRight),
		string(PrintPresentationDirection_CounterClockwiseFromTopLeft),
		string(PrintPresentationDirection_CounterClockwiseFromTopRight),
	}
}

func (s *PrintPresentationDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintPresentationDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintPresentationDirection(input string) (*PrintPresentationDirection, error) {
	vals := map[string]PrintPresentationDirection{
		"clockwisefrombottomleft":         PrintPresentationDirection_ClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrintPresentationDirection_ClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrintPresentationDirection_ClockwiseFromTopLeft,
		"clockwisefromtopright":           PrintPresentationDirection_ClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrintPresentationDirection_CounterClockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrintPresentationDirection_CounterClockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrintPresentationDirection_CounterClockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrintPresentationDirection_CounterClockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintPresentationDirection(input)
	return &out, nil
}
