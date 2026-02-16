package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterFeedOrientation string

const (
	PrinterFeedOrientation_LongEdgeFirst  PrinterFeedOrientation = "longEdgeFirst"
	PrinterFeedOrientation_ShortEdgeFirst PrinterFeedOrientation = "shortEdgeFirst"
)

func PossibleValuesForPrinterFeedOrientation() []string {
	return []string{
		string(PrinterFeedOrientation_LongEdgeFirst),
		string(PrinterFeedOrientation_ShortEdgeFirst),
	}
}

func (s *PrinterFeedOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterFeedOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterFeedOrientation(input string) (*PrinterFeedOrientation, error) {
	vals := map[string]PrinterFeedOrientation{
		"longedgefirst":  PrinterFeedOrientation_LongEdgeFirst,
		"shortedgefirst": PrinterFeedOrientation_ShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterFeedOrientation(input)
	return &out, nil
}
