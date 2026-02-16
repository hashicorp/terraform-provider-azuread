package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportFormat string

const (
	SecurityExportFormat_Eml SecurityExportFormat = "eml"
	SecurityExportFormat_Msg SecurityExportFormat = "msg"
	SecurityExportFormat_Pst SecurityExportFormat = "pst"
)

func PossibleValuesForSecurityExportFormat() []string {
	return []string{
		string(SecurityExportFormat_Eml),
		string(SecurityExportFormat_Msg),
		string(SecurityExportFormat_Pst),
	}
}

func (s *SecurityExportFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityExportFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityExportFormat(input string) (*SecurityExportFormat, error) {
	vals := map[string]SecurityExportFormat{
		"eml": SecurityExportFormat_Eml,
		"msg": SecurityExportFormat_Msg,
		"pst": SecurityExportFormat_Pst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityExportFormat(input)
	return &out, nil
}
