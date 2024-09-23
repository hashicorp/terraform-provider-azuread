package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetectionSource string

const (
	SecurityDetectionSource_Antivirus                                    SecurityDetectionSource = "antivirus"
	SecurityDetectionSource_AppGovernanceDetection                       SecurityDetectionSource = "appGovernanceDetection"
	SecurityDetectionSource_AppGovernancePolicy                          SecurityDetectionSource = "appGovernancePolicy"
	SecurityDetectionSource_AutomatedInvestigation                       SecurityDetectionSource = "automatedInvestigation"
	SecurityDetectionSource_AzureAdIdentityProtection                    SecurityDetectionSource = "azureAdIdentityProtection"
	SecurityDetectionSource_BuiltInMl                                    SecurityDetectionSource = "builtInMl"
	SecurityDetectionSource_CloudAppSecurity                             SecurityDetectionSource = "cloudAppSecurity"
	SecurityDetectionSource_CustomDetection                              SecurityDetectionSource = "customDetection"
	SecurityDetectionSource_CustomTi                                     SecurityDetectionSource = "customTi"
	SecurityDetectionSource_Manual                                       SecurityDetectionSource = "manual"
	SecurityDetectionSource_Microsoft365Defender                         SecurityDetectionSource = "microsoft365Defender"
	SecurityDetectionSource_MicrosoftDataLossPrevention                  SecurityDetectionSource = "microsoftDataLossPrevention"
	SecurityDetectionSource_MicrosoftDefenderForApiManagement            SecurityDetectionSource = "microsoftDefenderForApiManagement"
	SecurityDetectionSource_MicrosoftDefenderForAppService               SecurityDetectionSource = "microsoftDefenderForAppService"
	SecurityDetectionSource_MicrosoftDefenderForCloud                    SecurityDetectionSource = "microsoftDefenderForCloud"
	SecurityDetectionSource_MicrosoftDefenderForContainers               SecurityDetectionSource = "microsoftDefenderForContainers"
	SecurityDetectionSource_MicrosoftDefenderForDNS                      SecurityDetectionSource = "microsoftDefenderForDNS"
	SecurityDetectionSource_MicrosoftDefenderForDatabases                SecurityDetectionSource = "microsoftDefenderForDatabases"
	SecurityDetectionSource_MicrosoftDefenderForEndpoint                 SecurityDetectionSource = "microsoftDefenderForEndpoint"
	SecurityDetectionSource_MicrosoftDefenderForIdentity                 SecurityDetectionSource = "microsoftDefenderForIdentity"
	SecurityDetectionSource_MicrosoftDefenderForIoT                      SecurityDetectionSource = "microsoftDefenderForIoT"
	SecurityDetectionSource_MicrosoftDefenderForKeyVault                 SecurityDetectionSource = "microsoftDefenderForKeyVault"
	SecurityDetectionSource_MicrosoftDefenderForNetwork                  SecurityDetectionSource = "microsoftDefenderForNetwork"
	SecurityDetectionSource_MicrosoftDefenderForOffice365                SecurityDetectionSource = "microsoftDefenderForOffice365"
	SecurityDetectionSource_MicrosoftDefenderForResourceManager          SecurityDetectionSource = "microsoftDefenderForResourceManager"
	SecurityDetectionSource_MicrosoftDefenderForServers                  SecurityDetectionSource = "microsoftDefenderForServers"
	SecurityDetectionSource_MicrosoftDefenderForStorage                  SecurityDetectionSource = "microsoftDefenderForStorage"
	SecurityDetectionSource_MicrosoftDefenderThreatIntelligenceAnalytics SecurityDetectionSource = "microsoftDefenderThreatIntelligenceAnalytics"
	SecurityDetectionSource_MicrosoftInsiderRiskManagement               SecurityDetectionSource = "microsoftInsiderRiskManagement"
	SecurityDetectionSource_MicrosoftSentinel                            SecurityDetectionSource = "microsoftSentinel"
	SecurityDetectionSource_MicrosoftThreatExperts                       SecurityDetectionSource = "microsoftThreatExperts"
	SecurityDetectionSource_NrtAlerts                                    SecurityDetectionSource = "nrtAlerts"
	SecurityDetectionSource_ScheduledAlerts                              SecurityDetectionSource = "scheduledAlerts"
	SecurityDetectionSource_SmartScreen                                  SecurityDetectionSource = "smartScreen"
	SecurityDetectionSource_Unknown                                      SecurityDetectionSource = "unknown"
)

func PossibleValuesForSecurityDetectionSource() []string {
	return []string{
		string(SecurityDetectionSource_Antivirus),
		string(SecurityDetectionSource_AppGovernanceDetection),
		string(SecurityDetectionSource_AppGovernancePolicy),
		string(SecurityDetectionSource_AutomatedInvestigation),
		string(SecurityDetectionSource_AzureAdIdentityProtection),
		string(SecurityDetectionSource_BuiltInMl),
		string(SecurityDetectionSource_CloudAppSecurity),
		string(SecurityDetectionSource_CustomDetection),
		string(SecurityDetectionSource_CustomTi),
		string(SecurityDetectionSource_Manual),
		string(SecurityDetectionSource_Microsoft365Defender),
		string(SecurityDetectionSource_MicrosoftDataLossPrevention),
		string(SecurityDetectionSource_MicrosoftDefenderForApiManagement),
		string(SecurityDetectionSource_MicrosoftDefenderForAppService),
		string(SecurityDetectionSource_MicrosoftDefenderForCloud),
		string(SecurityDetectionSource_MicrosoftDefenderForContainers),
		string(SecurityDetectionSource_MicrosoftDefenderForDNS),
		string(SecurityDetectionSource_MicrosoftDefenderForDatabases),
		string(SecurityDetectionSource_MicrosoftDefenderForEndpoint),
		string(SecurityDetectionSource_MicrosoftDefenderForIdentity),
		string(SecurityDetectionSource_MicrosoftDefenderForIoT),
		string(SecurityDetectionSource_MicrosoftDefenderForKeyVault),
		string(SecurityDetectionSource_MicrosoftDefenderForNetwork),
		string(SecurityDetectionSource_MicrosoftDefenderForOffice365),
		string(SecurityDetectionSource_MicrosoftDefenderForResourceManager),
		string(SecurityDetectionSource_MicrosoftDefenderForServers),
		string(SecurityDetectionSource_MicrosoftDefenderForStorage),
		string(SecurityDetectionSource_MicrosoftDefenderThreatIntelligenceAnalytics),
		string(SecurityDetectionSource_MicrosoftInsiderRiskManagement),
		string(SecurityDetectionSource_MicrosoftSentinel),
		string(SecurityDetectionSource_MicrosoftThreatExperts),
		string(SecurityDetectionSource_NrtAlerts),
		string(SecurityDetectionSource_ScheduledAlerts),
		string(SecurityDetectionSource_SmartScreen),
		string(SecurityDetectionSource_Unknown),
	}
}

func (s *SecurityDetectionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDetectionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDetectionSource(input string) (*SecurityDetectionSource, error) {
	vals := map[string]SecurityDetectionSource{
		"antivirus":                                    SecurityDetectionSource_Antivirus,
		"appgovernancedetection":                       SecurityDetectionSource_AppGovernanceDetection,
		"appgovernancepolicy":                          SecurityDetectionSource_AppGovernancePolicy,
		"automatedinvestigation":                       SecurityDetectionSource_AutomatedInvestigation,
		"azureadidentityprotection":                    SecurityDetectionSource_AzureAdIdentityProtection,
		"builtinml":                                    SecurityDetectionSource_BuiltInMl,
		"cloudappsecurity":                             SecurityDetectionSource_CloudAppSecurity,
		"customdetection":                              SecurityDetectionSource_CustomDetection,
		"customti":                                     SecurityDetectionSource_CustomTi,
		"manual":                                       SecurityDetectionSource_Manual,
		"microsoft365defender":                         SecurityDetectionSource_Microsoft365Defender,
		"microsoftdatalossprevention":                  SecurityDetectionSource_MicrosoftDataLossPrevention,
		"microsoftdefenderforapimanagement":            SecurityDetectionSource_MicrosoftDefenderForApiManagement,
		"microsoftdefenderforappservice":               SecurityDetectionSource_MicrosoftDefenderForAppService,
		"microsoftdefenderforcloud":                    SecurityDetectionSource_MicrosoftDefenderForCloud,
		"microsoftdefenderforcontainers":               SecurityDetectionSource_MicrosoftDefenderForContainers,
		"microsoftdefenderfordns":                      SecurityDetectionSource_MicrosoftDefenderForDNS,
		"microsoftdefenderfordatabases":                SecurityDetectionSource_MicrosoftDefenderForDatabases,
		"microsoftdefenderforendpoint":                 SecurityDetectionSource_MicrosoftDefenderForEndpoint,
		"microsoftdefenderforidentity":                 SecurityDetectionSource_MicrosoftDefenderForIdentity,
		"microsoftdefenderforiot":                      SecurityDetectionSource_MicrosoftDefenderForIoT,
		"microsoftdefenderforkeyvault":                 SecurityDetectionSource_MicrosoftDefenderForKeyVault,
		"microsoftdefenderfornetwork":                  SecurityDetectionSource_MicrosoftDefenderForNetwork,
		"microsoftdefenderforoffice365":                SecurityDetectionSource_MicrosoftDefenderForOffice365,
		"microsoftdefenderforresourcemanager":          SecurityDetectionSource_MicrosoftDefenderForResourceManager,
		"microsoftdefenderforservers":                  SecurityDetectionSource_MicrosoftDefenderForServers,
		"microsoftdefenderforstorage":                  SecurityDetectionSource_MicrosoftDefenderForStorage,
		"microsoftdefenderthreatintelligenceanalytics": SecurityDetectionSource_MicrosoftDefenderThreatIntelligenceAnalytics,
		"microsoftinsiderriskmanagement":               SecurityDetectionSource_MicrosoftInsiderRiskManagement,
		"microsoftsentinel":                            SecurityDetectionSource_MicrosoftSentinel,
		"microsoftthreatexperts":                       SecurityDetectionSource_MicrosoftThreatExperts,
		"nrtalerts":                                    SecurityDetectionSource_NrtAlerts,
		"scheduledalerts":                              SecurityDetectionSource_ScheduledAlerts,
		"smartscreen":                                  SecurityDetectionSource_SmartScreen,
		"unknown":                                      SecurityDetectionSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDetectionSource(input)
	return &out, nil
}
