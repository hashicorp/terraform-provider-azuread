package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureBootWithDMAType string

const (
	SecureBootWithDMAType_NotConfigured SecureBootWithDMAType = "notConfigured"
	SecureBootWithDMAType_WithDMA       SecureBootWithDMAType = "withDMA"
	SecureBootWithDMAType_WithoutDMA    SecureBootWithDMAType = "withoutDMA"
)

func PossibleValuesForSecureBootWithDMAType() []string {
	return []string{
		string(SecureBootWithDMAType_NotConfigured),
		string(SecureBootWithDMAType_WithDMA),
		string(SecureBootWithDMAType_WithoutDMA),
	}
}

func (s *SecureBootWithDMAType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecureBootWithDMAType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecureBootWithDMAType(input string) (*SecureBootWithDMAType, error) {
	vals := map[string]SecureBootWithDMAType{
		"notconfigured": SecureBootWithDMAType_NotConfigured,
		"withdma":       SecureBootWithDMAType_WithDMA,
		"withoutdma":    SecureBootWithDMAType_WithoutDMA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecureBootWithDMAType(input)
	return &out, nil
}
