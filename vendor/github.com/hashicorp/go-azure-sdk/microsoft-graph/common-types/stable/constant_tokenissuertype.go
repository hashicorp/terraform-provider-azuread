package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenIssuerType string

const (
	TokenIssuerType_ADFederationServices           TokenIssuerType = "ADFederationServices"
	TokenIssuerType_ADFederationServicesMFAAdapter TokenIssuerType = "ADFederationServicesMFAAdapter"
	TokenIssuerType_AzureAD                        TokenIssuerType = "AzureAD"
	TokenIssuerType_AzureADBackupAuth              TokenIssuerType = "AzureADBackupAuth"
	TokenIssuerType_NPSExtension                   TokenIssuerType = "NPSExtension"
)

func PossibleValuesForTokenIssuerType() []string {
	return []string{
		string(TokenIssuerType_ADFederationServices),
		string(TokenIssuerType_ADFederationServicesMFAAdapter),
		string(TokenIssuerType_AzureAD),
		string(TokenIssuerType_AzureADBackupAuth),
		string(TokenIssuerType_NPSExtension),
	}
}

func (s *TokenIssuerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTokenIssuerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTokenIssuerType(input string) (*TokenIssuerType, error) {
	vals := map[string]TokenIssuerType{
		"adfederationservices":           TokenIssuerType_ADFederationServices,
		"adfederationservicesmfaadapter": TokenIssuerType_ADFederationServicesMFAAdapter,
		"azuread":                        TokenIssuerType_AzureAD,
		"azureadbackupauth":              TokenIssuerType_AzureADBackupAuth,
		"npsextension":                   TokenIssuerType_NPSExtension,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TokenIssuerType(input)
	return &out, nil
}
