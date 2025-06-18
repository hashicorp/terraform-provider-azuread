package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppCategory string

const (
	SecurityAppCategory_AccountingAndFinance     SecurityAppCategory = "accountingAndFinance"
	SecurityAppCategory_Advertising              SecurityAppCategory = "advertising"
	SecurityAppCategory_AiModelProvider          SecurityAppCategory = "aiModelProvider"
	SecurityAppCategory_BusinessIntelligence     SecurityAppCategory = "businessIntelligence"
	SecurityAppCategory_BusinessManagement       SecurityAppCategory = "businessManagement"
	SecurityAppCategory_ClientAiApp              SecurityAppCategory = "clientAiApp"
	SecurityAppCategory_CloudComputingPlatform   SecurityAppCategory = "cloudComputingPlatform"
	SecurityAppCategory_CloudStorage             SecurityAppCategory = "cloudStorage"
	SecurityAppCategory_CodeHosting              SecurityAppCategory = "codeHosting"
	SecurityAppCategory_Collaboration            SecurityAppCategory = "collaboration"
	SecurityAppCategory_Communications           SecurityAppCategory = "communications"
	SecurityAppCategory_ContentManagement        SecurityAppCategory = "contentManagement"
	SecurityAppCategory_ContentSharing           SecurityAppCategory = "contentSharing"
	SecurityAppCategory_Crm                      SecurityAppCategory = "crm"
	SecurityAppCategory_CustomerSupport          SecurityAppCategory = "customerSupport"
	SecurityAppCategory_DataAnalytics            SecurityAppCategory = "dataAnalytics"
	SecurityAppCategory_DevelopmentTools         SecurityAppCategory = "developmentTools"
	SecurityAppCategory_ECommerce                SecurityAppCategory = "eCommerce"
	SecurityAppCategory_Education                SecurityAppCategory = "education"
	SecurityAppCategory_Forums                   SecurityAppCategory = "forums"
	SecurityAppCategory_GenerativeAi             SecurityAppCategory = "generativeAi"
	SecurityAppCategory_Health                   SecurityAppCategory = "health"
	SecurityAppCategory_HostingServices          SecurityAppCategory = "hostingServices"
	SecurityAppCategory_HumanResourceManagement  SecurityAppCategory = "humanResourceManagement"
	SecurityAppCategory_InternetOfThings         SecurityAppCategory = "internetOfThings"
	SecurityAppCategory_ItServices               SecurityAppCategory = "itServices"
	SecurityAppCategory_Marketing                SecurityAppCategory = "marketing"
	SecurityAppCategory_McpServer                SecurityAppCategory = "mcpServer"
	SecurityAppCategory_NewsAndEntertainment     SecurityAppCategory = "newsAndEntertainment"
	SecurityAppCategory_OnlineMeetings           SecurityAppCategory = "onlineMeetings"
	SecurityAppCategory_OperationsManagement     SecurityAppCategory = "operationsManagement"
	SecurityAppCategory_PersonalInstantMessaging SecurityAppCategory = "personalInstantMessaging"
	SecurityAppCategory_ProductDesign            SecurityAppCategory = "productDesign"
	SecurityAppCategory_Productivity             SecurityAppCategory = "productivity"
	SecurityAppCategory_ProjectManagement        SecurityAppCategory = "projectManagement"
	SecurityAppCategory_PropertyManagement       SecurityAppCategory = "propertyManagement"
	SecurityAppCategory_Sales                    SecurityAppCategory = "sales"
	SecurityAppCategory_Security                 SecurityAppCategory = "security"
	SecurityAppCategory_SocialNetwork            SecurityAppCategory = "socialNetwork"
	SecurityAppCategory_SupplyChainAndLogistics  SecurityAppCategory = "supplyChainAndLogistics"
	SecurityAppCategory_TransportationAndTravel  SecurityAppCategory = "transportationAndTravel"
	SecurityAppCategory_Unknown                  SecurityAppCategory = "unknown"
	SecurityAppCategory_VendorManagementSystems  SecurityAppCategory = "vendorManagementSystems"
	SecurityAppCategory_WebAnalytics             SecurityAppCategory = "webAnalytics"
	SecurityAppCategory_Webemail                 SecurityAppCategory = "webemail"
	SecurityAppCategory_WebsiteMonitoring        SecurityAppCategory = "websiteMonitoring"
)

func PossibleValuesForSecurityAppCategory() []string {
	return []string{
		string(SecurityAppCategory_AccountingAndFinance),
		string(SecurityAppCategory_Advertising),
		string(SecurityAppCategory_AiModelProvider),
		string(SecurityAppCategory_BusinessIntelligence),
		string(SecurityAppCategory_BusinessManagement),
		string(SecurityAppCategory_ClientAiApp),
		string(SecurityAppCategory_CloudComputingPlatform),
		string(SecurityAppCategory_CloudStorage),
		string(SecurityAppCategory_CodeHosting),
		string(SecurityAppCategory_Collaboration),
		string(SecurityAppCategory_Communications),
		string(SecurityAppCategory_ContentManagement),
		string(SecurityAppCategory_ContentSharing),
		string(SecurityAppCategory_Crm),
		string(SecurityAppCategory_CustomerSupport),
		string(SecurityAppCategory_DataAnalytics),
		string(SecurityAppCategory_DevelopmentTools),
		string(SecurityAppCategory_ECommerce),
		string(SecurityAppCategory_Education),
		string(SecurityAppCategory_Forums),
		string(SecurityAppCategory_GenerativeAi),
		string(SecurityAppCategory_Health),
		string(SecurityAppCategory_HostingServices),
		string(SecurityAppCategory_HumanResourceManagement),
		string(SecurityAppCategory_InternetOfThings),
		string(SecurityAppCategory_ItServices),
		string(SecurityAppCategory_Marketing),
		string(SecurityAppCategory_McpServer),
		string(SecurityAppCategory_NewsAndEntertainment),
		string(SecurityAppCategory_OnlineMeetings),
		string(SecurityAppCategory_OperationsManagement),
		string(SecurityAppCategory_PersonalInstantMessaging),
		string(SecurityAppCategory_ProductDesign),
		string(SecurityAppCategory_Productivity),
		string(SecurityAppCategory_ProjectManagement),
		string(SecurityAppCategory_PropertyManagement),
		string(SecurityAppCategory_Sales),
		string(SecurityAppCategory_Security),
		string(SecurityAppCategory_SocialNetwork),
		string(SecurityAppCategory_SupplyChainAndLogistics),
		string(SecurityAppCategory_TransportationAndTravel),
		string(SecurityAppCategory_Unknown),
		string(SecurityAppCategory_VendorManagementSystems),
		string(SecurityAppCategory_WebAnalytics),
		string(SecurityAppCategory_Webemail),
		string(SecurityAppCategory_WebsiteMonitoring),
	}
}

func (s *SecurityAppCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppCategory(input string) (*SecurityAppCategory, error) {
	vals := map[string]SecurityAppCategory{
		"accountingandfinance":     SecurityAppCategory_AccountingAndFinance,
		"advertising":              SecurityAppCategory_Advertising,
		"aimodelprovider":          SecurityAppCategory_AiModelProvider,
		"businessintelligence":     SecurityAppCategory_BusinessIntelligence,
		"businessmanagement":       SecurityAppCategory_BusinessManagement,
		"clientaiapp":              SecurityAppCategory_ClientAiApp,
		"cloudcomputingplatform":   SecurityAppCategory_CloudComputingPlatform,
		"cloudstorage":             SecurityAppCategory_CloudStorage,
		"codehosting":              SecurityAppCategory_CodeHosting,
		"collaboration":            SecurityAppCategory_Collaboration,
		"communications":           SecurityAppCategory_Communications,
		"contentmanagement":        SecurityAppCategory_ContentManagement,
		"contentsharing":           SecurityAppCategory_ContentSharing,
		"crm":                      SecurityAppCategory_Crm,
		"customersupport":          SecurityAppCategory_CustomerSupport,
		"dataanalytics":            SecurityAppCategory_DataAnalytics,
		"developmenttools":         SecurityAppCategory_DevelopmentTools,
		"ecommerce":                SecurityAppCategory_ECommerce,
		"education":                SecurityAppCategory_Education,
		"forums":                   SecurityAppCategory_Forums,
		"generativeai":             SecurityAppCategory_GenerativeAi,
		"health":                   SecurityAppCategory_Health,
		"hostingservices":          SecurityAppCategory_HostingServices,
		"humanresourcemanagement":  SecurityAppCategory_HumanResourceManagement,
		"internetofthings":         SecurityAppCategory_InternetOfThings,
		"itservices":               SecurityAppCategory_ItServices,
		"marketing":                SecurityAppCategory_Marketing,
		"mcpserver":                SecurityAppCategory_McpServer,
		"newsandentertainment":     SecurityAppCategory_NewsAndEntertainment,
		"onlinemeetings":           SecurityAppCategory_OnlineMeetings,
		"operationsmanagement":     SecurityAppCategory_OperationsManagement,
		"personalinstantmessaging": SecurityAppCategory_PersonalInstantMessaging,
		"productdesign":            SecurityAppCategory_ProductDesign,
		"productivity":             SecurityAppCategory_Productivity,
		"projectmanagement":        SecurityAppCategory_ProjectManagement,
		"propertymanagement":       SecurityAppCategory_PropertyManagement,
		"sales":                    SecurityAppCategory_Sales,
		"security":                 SecurityAppCategory_Security,
		"socialnetwork":            SecurityAppCategory_SocialNetwork,
		"supplychainandlogistics":  SecurityAppCategory_SupplyChainAndLogistics,
		"transportationandtravel":  SecurityAppCategory_TransportationAndTravel,
		"unknown":                  SecurityAppCategory_Unknown,
		"vendormanagementsystems":  SecurityAppCategory_VendorManagementSystems,
		"webanalytics":             SecurityAppCategory_WebAnalytics,
		"webemail":                 SecurityAppCategory_Webemail,
		"websitemonitoring":        SecurityAppCategory_WebsiteMonitoring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppCategory(input)
	return &out, nil
}
