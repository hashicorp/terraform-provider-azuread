package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessWhatIfConditions struct {
	AuthenticationFlow *AuthenticationFlow              `json:"authenticationFlow,omitempty"`
	ClientAppType      *ConditionalAccessClientApp      `json:"clientAppType,omitempty"`
	Country            nullable.Type[string]            `json:"country,omitempty"`
	DeviceInfo         *DeviceInfo                      `json:"deviceInfo,omitempty"`
	DevicePlatform     *ConditionalAccessDevicePlatform `json:"devicePlatform,omitempty"`
	IPAddress          nullable.Type[string]            `json:"ipAddress,omitempty"`
	InsiderRiskLevel   *InsiderRiskLevel                `json:"insiderRiskLevel,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ServicePrincipalRiskLevel *RiskLevel `json:"servicePrincipalRiskLevel,omitempty"`
	SignInRiskLevel           *RiskLevel `json:"signInRiskLevel,omitempty"`
	UserRiskLevel             *RiskLevel `json:"userRiskLevel,omitempty"`
}
