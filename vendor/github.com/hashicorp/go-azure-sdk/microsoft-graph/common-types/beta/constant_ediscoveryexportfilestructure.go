package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryExportFileStructure string

const (
	EdiscoveryExportFileStructure_Directory EdiscoveryExportFileStructure = "directory"
	EdiscoveryExportFileStructure_None      EdiscoveryExportFileStructure = "none"
	EdiscoveryExportFileStructure_Pst       EdiscoveryExportFileStructure = "pst"
)

func PossibleValuesForEdiscoveryExportFileStructure() []string {
	return []string{
		string(EdiscoveryExportFileStructure_Directory),
		string(EdiscoveryExportFileStructure_None),
		string(EdiscoveryExportFileStructure_Pst),
	}
}

func (s *EdiscoveryExportFileStructure) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryExportFileStructure(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryExportFileStructure(input string) (*EdiscoveryExportFileStructure, error) {
	vals := map[string]EdiscoveryExportFileStructure{
		"directory": EdiscoveryExportFileStructure_Directory,
		"none":      EdiscoveryExportFileStructure_None,
		"pst":       EdiscoveryExportFileStructure_Pst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryExportFileStructure(input)
	return &out, nil
}
