package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDocumentVersion string

const (
	SecurityDocumentVersion_All       SecurityDocumentVersion = "all"
	SecurityDocumentVersion_Latest    SecurityDocumentVersion = "latest"
	SecurityDocumentVersion_Recent10  SecurityDocumentVersion = "recent10"
	SecurityDocumentVersion_Recent100 SecurityDocumentVersion = "recent100"
)

func PossibleValuesForSecurityDocumentVersion() []string {
	return []string{
		string(SecurityDocumentVersion_All),
		string(SecurityDocumentVersion_Latest),
		string(SecurityDocumentVersion_Recent10),
		string(SecurityDocumentVersion_Recent100),
	}
}

func (s *SecurityDocumentVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDocumentVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDocumentVersion(input string) (*SecurityDocumentVersion, error) {
	vals := map[string]SecurityDocumentVersion{
		"all":       SecurityDocumentVersion_All,
		"latest":    SecurityDocumentVersion_Latest,
		"recent10":  SecurityDocumentVersion_Recent10,
		"recent100": SecurityDocumentVersion_Recent100,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDocumentVersion(input)
	return &out, nil
}
