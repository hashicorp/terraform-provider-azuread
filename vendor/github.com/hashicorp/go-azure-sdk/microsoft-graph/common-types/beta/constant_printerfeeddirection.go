package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterFeedDirection string

const (
	PrinterFeedDirection_LongEdgeFirst  PrinterFeedDirection = "longEdgeFirst"
	PrinterFeedDirection_ShortEdgeFirst PrinterFeedDirection = "shortEdgeFirst"
)

func PossibleValuesForPrinterFeedDirection() []string {
	return []string{
		string(PrinterFeedDirection_LongEdgeFirst),
		string(PrinterFeedDirection_ShortEdgeFirst),
	}
}

func (s *PrinterFeedDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterFeedDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterFeedDirection(input string) (*PrinterFeedDirection, error) {
	vals := map[string]PrinterFeedDirection{
		"longedgefirst":  PrinterFeedDirection_LongEdgeFirst,
		"shortedgefirst": PrinterFeedDirection_ShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterFeedDirection(input)
	return &out, nil
}
