package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInConditions struct {
	// Type of authentication flow. The possible value is: deviceCodeFlow or authenticationTransfer. Default value is none.
	AuthenticationFlow *AuthenticationFlow `json:"authenticationFlow,omitempty"`

	// Client application type. The possible value is: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync,
	// easSupported, other, unknownFutureValue. Default value is all.
	ClientAppType *ConditionalAccessClientApp `json:"clientAppType,omitempty"`

	// Country from where the identity is authenticating.
	Country nullable.Type[string] `json:"country,omitempty"`

	// Information about the device used for the sign-in.
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`

	// Device platform. The possible value is: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue, linux.
	// Default value is all.
	DevicePlatform *ConditionalAccessDevicePlatform `json:"devicePlatform,omitempty"`

	// Ip address of the authenticating identity.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// Insider risk associated with the authenticating user. The possible value is: none, minor, moderate, elevated,
	// unknownFutureValue. Default value is none.
	InsiderRiskLevel *InsiderRiskLevel `json:"insiderRiskLevel,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Risk associated with the service principal. The possible value is: low, medium, high, hidden, none,
	// unknownFutureValue. Default value is none.
	ServicePrincipalRiskLevel *RiskLevel `json:"servicePrincipalRiskLevel,omitempty"`

	// Sign-in risk associated with the user. The possible value is: low, medium, high, hidden, none, unknownFutureValue.
	// Default value is none.
	SignInRiskLevel *RiskLevel `json:"signInRiskLevel,omitempty"`

	// The authenticating user's risk level. The possible value is: low, medium, high, hidden, none, unknownFutureValue.
	// Default value is none.
	UserRiskLevel *RiskLevel `json:"userRiskLevel,omitempty"`
}
