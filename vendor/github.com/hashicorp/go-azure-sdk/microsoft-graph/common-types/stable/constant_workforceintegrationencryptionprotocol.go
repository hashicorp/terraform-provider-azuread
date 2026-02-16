package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationEncryptionProtocol string

const (
	WorkforceIntegrationEncryptionProtocol_SharedSecret WorkforceIntegrationEncryptionProtocol = "sharedSecret"
)

func PossibleValuesForWorkforceIntegrationEncryptionProtocol() []string {
	return []string{
		string(WorkforceIntegrationEncryptionProtocol_SharedSecret),
	}
}

func (s *WorkforceIntegrationEncryptionProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkforceIntegrationEncryptionProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkforceIntegrationEncryptionProtocol(input string) (*WorkforceIntegrationEncryptionProtocol, error) {
	vals := map[string]WorkforceIntegrationEncryptionProtocol{
		"sharedsecret": WorkforceIntegrationEncryptionProtocol_SharedSecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationEncryptionProtocol(input)
	return &out, nil
}
