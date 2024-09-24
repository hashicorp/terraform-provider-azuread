package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityServiceSource string

const (
	SecurityServiceSource_AzureAdIdentityProtection      SecurityServiceSource = "azureAdIdentityProtection"
	SecurityServiceSource_DataLossPrevention             SecurityServiceSource = "dataLossPrevention"
	SecurityServiceSource_Microsoft365Defender           SecurityServiceSource = "microsoft365Defender"
	SecurityServiceSource_MicrosoftAppGovernance         SecurityServiceSource = "microsoftAppGovernance"
	SecurityServiceSource_MicrosoftDefenderForCloud      SecurityServiceSource = "microsoftDefenderForCloud"
	SecurityServiceSource_MicrosoftDefenderForCloudApps  SecurityServiceSource = "microsoftDefenderForCloudApps"
	SecurityServiceSource_MicrosoftDefenderForEndpoint   SecurityServiceSource = "microsoftDefenderForEndpoint"
	SecurityServiceSource_MicrosoftDefenderForIdentity   SecurityServiceSource = "microsoftDefenderForIdentity"
	SecurityServiceSource_MicrosoftDefenderForOffice365  SecurityServiceSource = "microsoftDefenderForOffice365"
	SecurityServiceSource_MicrosoftInsiderRiskManagement SecurityServiceSource = "microsoftInsiderRiskManagement"
	SecurityServiceSource_MicrosoftSentinel              SecurityServiceSource = "microsoftSentinel"
	SecurityServiceSource_Unknown                        SecurityServiceSource = "unknown"
)

func PossibleValuesForSecurityServiceSource() []string {
	return []string{
		string(SecurityServiceSource_AzureAdIdentityProtection),
		string(SecurityServiceSource_DataLossPrevention),
		string(SecurityServiceSource_Microsoft365Defender),
		string(SecurityServiceSource_MicrosoftAppGovernance),
		string(SecurityServiceSource_MicrosoftDefenderForCloud),
		string(SecurityServiceSource_MicrosoftDefenderForCloudApps),
		string(SecurityServiceSource_MicrosoftDefenderForEndpoint),
		string(SecurityServiceSource_MicrosoftDefenderForIdentity),
		string(SecurityServiceSource_MicrosoftDefenderForOffice365),
		string(SecurityServiceSource_MicrosoftInsiderRiskManagement),
		string(SecurityServiceSource_MicrosoftSentinel),
		string(SecurityServiceSource_Unknown),
	}
}

func (s *SecurityServiceSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityServiceSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityServiceSource(input string) (*SecurityServiceSource, error) {
	vals := map[string]SecurityServiceSource{
		"azureadidentityprotection":      SecurityServiceSource_AzureAdIdentityProtection,
		"datalossprevention":             SecurityServiceSource_DataLossPrevention,
		"microsoft365defender":           SecurityServiceSource_Microsoft365Defender,
		"microsoftappgovernance":         SecurityServiceSource_MicrosoftAppGovernance,
		"microsoftdefenderforcloud":      SecurityServiceSource_MicrosoftDefenderForCloud,
		"microsoftdefenderforcloudapps":  SecurityServiceSource_MicrosoftDefenderForCloudApps,
		"microsoftdefenderforendpoint":   SecurityServiceSource_MicrosoftDefenderForEndpoint,
		"microsoftdefenderforidentity":   SecurityServiceSource_MicrosoftDefenderForIdentity,
		"microsoftdefenderforoffice365":  SecurityServiceSource_MicrosoftDefenderForOffice365,
		"microsoftinsiderriskmanagement": SecurityServiceSource_MicrosoftInsiderRiskManagement,
		"microsoftsentinel":              SecurityServiceSource_MicrosoftSentinel,
		"unknown":                        SecurityServiceSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityServiceSource(input)
	return &out, nil
}
