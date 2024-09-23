package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsContentExperienceType string

const (
	ExternalConnectorsContentExperienceType_Compliance ExternalConnectorsContentExperienceType = "compliance"
	ExternalConnectorsContentExperienceType_Search     ExternalConnectorsContentExperienceType = "search"
)

func PossibleValuesForExternalConnectorsContentExperienceType() []string {
	return []string{
		string(ExternalConnectorsContentExperienceType_Compliance),
		string(ExternalConnectorsContentExperienceType_Search),
	}
}

func (s *ExternalConnectorsContentExperienceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsContentExperienceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsContentExperienceType(input string) (*ExternalConnectorsContentExperienceType, error) {
	vals := map[string]ExternalConnectorsContentExperienceType{
		"compliance": ExternalConnectorsContentExperienceType_Compliance,
		"search":     ExternalConnectorsContentExperienceType_Search,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsContentExperienceType(input)
	return &out, nil
}
