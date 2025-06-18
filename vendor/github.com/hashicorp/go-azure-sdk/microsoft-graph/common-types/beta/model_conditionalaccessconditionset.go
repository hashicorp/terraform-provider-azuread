package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSet struct {
	// Applications and user actions included in and excluded from the policy. Required.
	Applications ConditionalAccessApplications `json:"applications"`

	// Authentication flows included in the policy scope. For more information, see Conditional Access: Authentication
	// flows.
	AuthenticationFlows *ConditionalAccessAuthenticationFlows `json:"authenticationFlows,omitempty"`

	// Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients,
	// exchangeActiveSync, easSupported, other. Required. The easUnsupported enumeration member is deprecated in favor of
	// exchangeActiveSync, which includes EAS supported and unsupported platforms.
	ClientAppTypes []ConditionalAccessClientApp `json:"clientAppTypes"`

	// Client applications (service principals and workload identities) included in and excluded from the policy. Either
	// users or clientApplications is required.
	ClientApplications *ConditionalAccessClientApplications `json:"clientApplications,omitempty"`

	// Device states in the policy. To be deprecated and removed. Use the devices property instead.
	DeviceStates *ConditionalAccessDeviceStates `json:"deviceStates,omitempty"`

	// Devices in the policy.
	Devices *ConditionalAccessDevices `json:"devices,omitempty"`

	// Insider risk levels included in the policy. The possible values are: minor, moderate, elevated, unknownFutureValue.
	InsiderRiskLevels *ConditionalAccessInsiderRiskLevels `json:"insiderRiskLevels,omitempty"`

	// Locations included in and excluded from the policy.
	Locations *ConditionalAccessLocations `json:"locations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Platforms included in and excluded from the policy.
	Platforms *ConditionalAccessPlatforms `json:"platforms,omitempty"`

	// Service principal risk levels included in the policy. Possible values are: low, medium, high, none,
	// unknownFutureValue.
	ServicePrincipalRiskLevels *[]RiskLevel `json:"servicePrincipalRiskLevels,omitempty"`

	// Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
	// Required.
	SignInRiskLevels []RiskLevel `json:"signInRiskLevels"`

	// User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
	// Required.
	UserRiskLevels []RiskLevel `json:"userRiskLevels"`

	// Users, groups, and roles included in and excluded from the policy. Either users or clientApplications is required.
	Users *ConditionalAccessUsers `json:"users,omitempty"`
}
