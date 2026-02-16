package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentitySourceType string

const (
	IdentitySourceType_AzureActiveDirectory IdentitySourceType = "azureActiveDirectory"
	IdentitySourceType_External             IdentitySourceType = "external"
)

func PossibleValuesForIdentitySourceType() []string {
	return []string{
		string(IdentitySourceType_AzureActiveDirectory),
		string(IdentitySourceType_External),
	}
}

func (s *IdentitySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentitySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentitySourceType(input string) (*IdentitySourceType, error) {
	vals := map[string]IdentitySourceType{
		"azureactivedirectory": IdentitySourceType_AzureActiveDirectory,
		"external":             IdentitySourceType_External,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentitySourceType(input)
	return &out, nil
}
