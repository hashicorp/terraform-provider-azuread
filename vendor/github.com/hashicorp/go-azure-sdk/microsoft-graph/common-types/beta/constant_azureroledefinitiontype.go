package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureRoleDefinitionType string

const (
	AzureRoleDefinitionType_Custom AzureRoleDefinitionType = "custom"
	AzureRoleDefinitionType_System AzureRoleDefinitionType = "system"
)

func PossibleValuesForAzureRoleDefinitionType() []string {
	return []string{
		string(AzureRoleDefinitionType_Custom),
		string(AzureRoleDefinitionType_System),
	}
}

func (s *AzureRoleDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureRoleDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureRoleDefinitionType(input string) (*AzureRoleDefinitionType, error) {
	vals := map[string]AzureRoleDefinitionType{
		"custom": AzureRoleDefinitionType_Custom,
		"system": AzureRoleDefinitionType_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureRoleDefinitionType(input)
	return &out, nil
}
