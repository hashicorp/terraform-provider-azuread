package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureEncryption string

const (
	AzureEncryption_Customer          AzureEncryption = "customer"
	AzureEncryption_MicrosoftKeyVault AzureEncryption = "microsoftKeyVault"
	AzureEncryption_MicrosoftStorage  AzureEncryption = "microsoftStorage"
)

func PossibleValuesForAzureEncryption() []string {
	return []string{
		string(AzureEncryption_Customer),
		string(AzureEncryption_MicrosoftKeyVault),
		string(AzureEncryption_MicrosoftStorage),
	}
}

func (s *AzureEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureEncryption(input string) (*AzureEncryption, error) {
	vals := map[string]AzureEncryption{
		"customer":          AzureEncryption_Customer,
		"microsoftkeyvault": AzureEncryption_MicrosoftKeyVault,
		"microsoftstorage":  AzureEncryption_MicrosoftStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureEncryption(input)
	return &out, nil
}
