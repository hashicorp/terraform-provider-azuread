package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintOrientation string

const (
	PrintOrientation_Landscape        PrintOrientation = "landscape"
	PrintOrientation_Portrait         PrintOrientation = "portrait"
	PrintOrientation_ReverseLandscape PrintOrientation = "reverseLandscape"
	PrintOrientation_ReversePortrait  PrintOrientation = "reversePortrait"
)

func PossibleValuesForPrintOrientation() []string {
	return []string{
		string(PrintOrientation_Landscape),
		string(PrintOrientation_Portrait),
		string(PrintOrientation_ReverseLandscape),
		string(PrintOrientation_ReversePortrait),
	}
}

func (s *PrintOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintOrientation(input string) (*PrintOrientation, error) {
	vals := map[string]PrintOrientation{
		"landscape":        PrintOrientation_Landscape,
		"portrait":         PrintOrientation_Portrait,
		"reverselandscape": PrintOrientation_ReverseLandscape,
		"reverseportrait":  PrintOrientation_ReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintOrientation(input)
	return &out, nil
}
