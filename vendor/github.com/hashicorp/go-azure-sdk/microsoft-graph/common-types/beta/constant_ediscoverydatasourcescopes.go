package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceScopes string

const (
	EdiscoveryDataSourceScopes_AllCaseCustodians              EdiscoveryDataSourceScopes = "allCaseCustodians"
	EdiscoveryDataSourceScopes_AllCaseNoncustodialDataSources EdiscoveryDataSourceScopes = "allCaseNoncustodialDataSources"
	EdiscoveryDataSourceScopes_AllTenantMailboxes             EdiscoveryDataSourceScopes = "allTenantMailboxes"
	EdiscoveryDataSourceScopes_AllTenantSites                 EdiscoveryDataSourceScopes = "allTenantSites"
	EdiscoveryDataSourceScopes_None                           EdiscoveryDataSourceScopes = "none"
)

func PossibleValuesForEdiscoveryDataSourceScopes() []string {
	return []string{
		string(EdiscoveryDataSourceScopes_AllCaseCustodians),
		string(EdiscoveryDataSourceScopes_AllCaseNoncustodialDataSources),
		string(EdiscoveryDataSourceScopes_AllTenantMailboxes),
		string(EdiscoveryDataSourceScopes_AllTenantSites),
		string(EdiscoveryDataSourceScopes_None),
	}
}

func (s *EdiscoveryDataSourceScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryDataSourceScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryDataSourceScopes(input string) (*EdiscoveryDataSourceScopes, error) {
	vals := map[string]EdiscoveryDataSourceScopes{
		"allcasecustodians":              EdiscoveryDataSourceScopes_AllCaseCustodians,
		"allcasenoncustodialdatasources": EdiscoveryDataSourceScopes_AllCaseNoncustodialDataSources,
		"alltenantmailboxes":             EdiscoveryDataSourceScopes_AllTenantMailboxes,
		"alltenantsites":                 EdiscoveryDataSourceScopes_AllTenantSites,
		"none":                           EdiscoveryDataSourceScopes_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryDataSourceScopes(input)
	return &out, nil
}
