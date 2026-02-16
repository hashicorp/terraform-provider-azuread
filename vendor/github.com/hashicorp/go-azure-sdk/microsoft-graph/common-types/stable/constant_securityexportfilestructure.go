package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportFileStructure string

const (
	SecurityExportFileStructure_Directory SecurityExportFileStructure = "directory"
	SecurityExportFileStructure_Msg       SecurityExportFileStructure = "msg"
	SecurityExportFileStructure_None      SecurityExportFileStructure = "none"
	SecurityExportFileStructure_Pst       SecurityExportFileStructure = "pst"
)

func PossibleValuesForSecurityExportFileStructure() []string {
	return []string{
		string(SecurityExportFileStructure_Directory),
		string(SecurityExportFileStructure_Msg),
		string(SecurityExportFileStructure_None),
		string(SecurityExportFileStructure_Pst),
	}
}

func (s *SecurityExportFileStructure) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityExportFileStructure(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityExportFileStructure(input string) (*SecurityExportFileStructure, error) {
	vals := map[string]SecurityExportFileStructure{
		"directory": SecurityExportFileStructure_Directory,
		"msg":       SecurityExportFileStructure_Msg,
		"none":      SecurityExportFileStructure_None,
		"pst":       SecurityExportFileStructure_Pst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityExportFileStructure(input)
	return &out, nil
}
