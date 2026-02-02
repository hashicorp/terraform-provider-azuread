package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DerivedCredentialProviderType string

const (
	DerivedCredentialProviderType_EntrustDataCard DerivedCredentialProviderType = "entrustDataCard"
	DerivedCredentialProviderType_Intercede       DerivedCredentialProviderType = "intercede"
	DerivedCredentialProviderType_NotConfigured   DerivedCredentialProviderType = "notConfigured"
	DerivedCredentialProviderType_Purebred        DerivedCredentialProviderType = "purebred"
	DerivedCredentialProviderType_XTec            DerivedCredentialProviderType = "xTec"
)

func PossibleValuesForDerivedCredentialProviderType() []string {
	return []string{
		string(DerivedCredentialProviderType_EntrustDataCard),
		string(DerivedCredentialProviderType_Intercede),
		string(DerivedCredentialProviderType_NotConfigured),
		string(DerivedCredentialProviderType_Purebred),
		string(DerivedCredentialProviderType_XTec),
	}
}

func (s *DerivedCredentialProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDerivedCredentialProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDerivedCredentialProviderType(input string) (*DerivedCredentialProviderType, error) {
	vals := map[string]DerivedCredentialProviderType{
		"entrustdatacard": DerivedCredentialProviderType_EntrustDataCard,
		"intercede":       DerivedCredentialProviderType_Intercede,
		"notconfigured":   DerivedCredentialProviderType_NotConfigured,
		"purebred":        DerivedCredentialProviderType_Purebred,
		"xtec":            DerivedCredentialProviderType_XTec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DerivedCredentialProviderType(input)
	return &out, nil
}
