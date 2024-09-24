package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2bIdentityProvidersType string

const (
	B2bIdentityProvidersType_AzureActiveDirectory    B2bIdentityProvidersType = "azureActiveDirectory"
	B2bIdentityProvidersType_DefaultConfiguredIdp    B2bIdentityProvidersType = "defaultConfiguredIdp"
	B2bIdentityProvidersType_EmailOneTimePasscode    B2bIdentityProvidersType = "emailOneTimePasscode"
	B2bIdentityProvidersType_ExternalFederation      B2bIdentityProvidersType = "externalFederation"
	B2bIdentityProvidersType_MicrosoftAccount        B2bIdentityProvidersType = "microsoftAccount"
	B2bIdentityProvidersType_SocialIdentityProviders B2bIdentityProvidersType = "socialIdentityProviders"
)

func PossibleValuesForB2bIdentityProvidersType() []string {
	return []string{
		string(B2bIdentityProvidersType_AzureActiveDirectory),
		string(B2bIdentityProvidersType_DefaultConfiguredIdp),
		string(B2bIdentityProvidersType_EmailOneTimePasscode),
		string(B2bIdentityProvidersType_ExternalFederation),
		string(B2bIdentityProvidersType_MicrosoftAccount),
		string(B2bIdentityProvidersType_SocialIdentityProviders),
	}
}

func (s *B2bIdentityProvidersType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseB2bIdentityProvidersType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseB2bIdentityProvidersType(input string) (*B2bIdentityProvidersType, error) {
	vals := map[string]B2bIdentityProvidersType{
		"azureactivedirectory":    B2bIdentityProvidersType_AzureActiveDirectory,
		"defaultconfiguredidp":    B2bIdentityProvidersType_DefaultConfiguredIdp,
		"emailonetimepasscode":    B2bIdentityProvidersType_EmailOneTimePasscode,
		"externalfederation":      B2bIdentityProvidersType_ExternalFederation,
		"microsoftaccount":        B2bIdentityProvidersType_MicrosoftAccount,
		"socialidentityproviders": B2bIdentityProvidersType_SocialIdentityProviders,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := B2bIdentityProvidersType(input)
	return &out, nil
}
