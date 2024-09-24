package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintQuality string

const (
	PrintQuality_High   PrintQuality = "high"
	PrintQuality_Low    PrintQuality = "low"
	PrintQuality_Medium PrintQuality = "medium"
)

func PossibleValuesForPrintQuality() []string {
	return []string{
		string(PrintQuality_High),
		string(PrintQuality_Low),
		string(PrintQuality_Medium),
	}
}

func (s *PrintQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintQuality(input string) (*PrintQuality, error) {
	vals := map[string]PrintQuality{
		"high":   PrintQuality_High,
		"low":    PrintQuality_Low,
		"medium": PrintQuality_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintQuality(input)
	return &out, nil
}
