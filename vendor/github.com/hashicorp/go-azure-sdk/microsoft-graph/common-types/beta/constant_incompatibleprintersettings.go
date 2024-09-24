package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncompatiblePrinterSettings string

const (
	IncompatiblePrinterSettings_Hide IncompatiblePrinterSettings = "hide"
	IncompatiblePrinterSettings_Show IncompatiblePrinterSettings = "show"
)

func PossibleValuesForIncompatiblePrinterSettings() []string {
	return []string{
		string(IncompatiblePrinterSettings_Hide),
		string(IncompatiblePrinterSettings_Show),
	}
}

func (s *IncompatiblePrinterSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncompatiblePrinterSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncompatiblePrinterSettings(input string) (*IncompatiblePrinterSettings, error) {
	vals := map[string]IncompatiblePrinterSettings{
		"hide": IncompatiblePrinterSettings_Hide,
		"show": IncompatiblePrinterSettings_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncompatiblePrinterSettings(input)
	return &out, nil
}
