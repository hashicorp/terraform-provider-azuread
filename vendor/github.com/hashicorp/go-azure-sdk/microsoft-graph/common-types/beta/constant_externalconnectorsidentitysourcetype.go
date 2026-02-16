package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsIdentitySourceType string

const (
	ExternalConnectorsIdentitySourceType_AzureActiveDirectory ExternalConnectorsIdentitySourceType = "azureActiveDirectory"
	ExternalConnectorsIdentitySourceType_External             ExternalConnectorsIdentitySourceType = "external"
)

func PossibleValuesForExternalConnectorsIdentitySourceType() []string {
	return []string{
		string(ExternalConnectorsIdentitySourceType_AzureActiveDirectory),
		string(ExternalConnectorsIdentitySourceType_External),
	}
}

func (s *ExternalConnectorsIdentitySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsIdentitySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsIdentitySourceType(input string) (*ExternalConnectorsIdentitySourceType, error) {
	vals := map[string]ExternalConnectorsIdentitySourceType{
		"azureactivedirectory": ExternalConnectorsIdentitySourceType_AzureActiveDirectory,
		"external":             ExternalConnectorsIdentitySourceType_External,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsIdentitySourceType(input)
	return &out, nil
}
