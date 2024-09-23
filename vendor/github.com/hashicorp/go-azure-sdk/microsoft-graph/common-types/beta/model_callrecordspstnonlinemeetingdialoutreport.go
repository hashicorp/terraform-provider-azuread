package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnOnlineMeetingDialoutReport struct {
	// Currency used to calculate the cost of the call. For details, see ISO 4217.
	Currency nullable.Type[string] `json:"currency,omitempty"`

	// Indicates whether the call was Domestic (within a country or region) or International (outside a country or region)
	// based on the user's location.
	DestinationContext nullable.Type[string] `json:"destinationContext,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total costs of all the calls within the selected time range, including call charges and connection fees.
	TotalCallCharge nullable.Type[float64] `json:"totalCallCharge,omitempty"`

	// Total duration of all the calls within the selected time range, in seconds.
	TotalCallSeconds nullable.Type[int64] `json:"totalCallSeconds,omitempty"`

	// Total number of dial-out calls within the selected time range.
	TotalCalls nullable.Type[int64] `json:"totalCalls,omitempty"`

	// Country code of the user. For details, see ISO 3166-1 alpha-2.
	UsageLocation nullable.Type[string] `json:"usageLocation,omitempty"`

	// Display name of the user.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// The unique identifier (GUID) of the user in Microsoft Entra ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user principal name (sign-in name) in Microsoft Entra ID. This is usually the same as the user's SIP address, and
	// can be same as the user's e-mail address.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
