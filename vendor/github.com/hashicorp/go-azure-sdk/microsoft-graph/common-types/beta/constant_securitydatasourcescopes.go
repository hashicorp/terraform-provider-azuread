package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSourceScopes string

const (
	SecurityDataSourceScopes_AllCaseCustodians              SecurityDataSourceScopes = "allCaseCustodians"
	SecurityDataSourceScopes_AllCaseNoncustodialDataSources SecurityDataSourceScopes = "allCaseNoncustodialDataSources"
	SecurityDataSourceScopes_AllTenantMailboxes             SecurityDataSourceScopes = "allTenantMailboxes"
	SecurityDataSourceScopes_AllTenantSites                 SecurityDataSourceScopes = "allTenantSites"
	SecurityDataSourceScopes_None                           SecurityDataSourceScopes = "none"
)

func PossibleValuesForSecurityDataSourceScopes() []string {
	return []string{
		string(SecurityDataSourceScopes_AllCaseCustodians),
		string(SecurityDataSourceScopes_AllCaseNoncustodialDataSources),
		string(SecurityDataSourceScopes_AllTenantMailboxes),
		string(SecurityDataSourceScopes_AllTenantSites),
		string(SecurityDataSourceScopes_None),
	}
}

func (s *SecurityDataSourceScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDataSourceScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDataSourceScopes(input string) (*SecurityDataSourceScopes, error) {
	vals := map[string]SecurityDataSourceScopes{
		"allcasecustodians":              SecurityDataSourceScopes_AllCaseCustodians,
		"allcasenoncustodialdatasources": SecurityDataSourceScopes_AllCaseNoncustodialDataSources,
		"alltenantmailboxes":             SecurityDataSourceScopes_AllTenantMailboxes,
		"alltenantsites":                 SecurityDataSourceScopes_AllTenantSites,
		"none":                           SecurityDataSourceScopes_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDataSourceScopes(input)
	return &out, nil
}
