package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintDuplexMode string

const (
	PrintDuplexMode_FlipOnLongEdge  PrintDuplexMode = "flipOnLongEdge"
	PrintDuplexMode_FlipOnShortEdge PrintDuplexMode = "flipOnShortEdge"
	PrintDuplexMode_OneSided        PrintDuplexMode = "oneSided"
)

func PossibleValuesForPrintDuplexMode() []string {
	return []string{
		string(PrintDuplexMode_FlipOnLongEdge),
		string(PrintDuplexMode_FlipOnShortEdge),
		string(PrintDuplexMode_OneSided),
	}
}

func (s *PrintDuplexMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintDuplexMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintDuplexMode(input string) (*PrintDuplexMode, error) {
	vals := map[string]PrintDuplexMode{
		"fliponlongedge":  PrintDuplexMode_FlipOnLongEdge,
		"fliponshortedge": PrintDuplexMode_FlipOnShortEdge,
		"onesided":        PrintDuplexMode_OneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintDuplexMode(input)
	return &out, nil
}
