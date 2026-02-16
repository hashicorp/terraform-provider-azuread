package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintDuplexConfiguration string

const (
	PrintDuplexConfiguration_OneSided          PrintDuplexConfiguration = "oneSided"
	PrintDuplexConfiguration_TwoSidedLongEdge  PrintDuplexConfiguration = "twoSidedLongEdge"
	PrintDuplexConfiguration_TwoSidedShortEdge PrintDuplexConfiguration = "twoSidedShortEdge"
)

func PossibleValuesForPrintDuplexConfiguration() []string {
	return []string{
		string(PrintDuplexConfiguration_OneSided),
		string(PrintDuplexConfiguration_TwoSidedLongEdge),
		string(PrintDuplexConfiguration_TwoSidedShortEdge),
	}
}

func (s *PrintDuplexConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintDuplexConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintDuplexConfiguration(input string) (*PrintDuplexConfiguration, error) {
	vals := map[string]PrintDuplexConfiguration{
		"onesided":          PrintDuplexConfiguration_OneSided,
		"twosidedlongedge":  PrintDuplexConfiguration_TwoSidedLongEdge,
		"twosidedshortedge": PrintDuplexConfiguration_TwoSidedShortEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintDuplexConfiguration(input)
	return &out, nil
}
