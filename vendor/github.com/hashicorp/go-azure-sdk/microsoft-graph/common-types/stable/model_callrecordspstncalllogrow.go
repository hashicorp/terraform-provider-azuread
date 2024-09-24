package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnCallLogRow struct {
	// The source of the call duration data. If the call uses a third-party telecommunications operator via the Operator
	// Connect Program, the operator can provide their own call duration data. In this case, the property value is operator.
	// Otherwise, the value is microsoft.
	CallDurationSource *CallRecordsPstnCallDurationSource `json:"callDurationSource,omitempty"`

	// Call identifier. Not guaranteed to be unique.
	CallId nullable.Type[string] `json:"callId,omitempty"`

	// Indicates whether the call was a PSTN outbound or inbound call and the type of call, such as a call placed by a user
	// or an audio conference.
	CallType nullable.Type[string] `json:"callType,omitempty"`

	// Number dialed in E.164 format.
	CalleeNumber nullable.Type[string] `json:"calleeNumber,omitempty"`

	// Number that received the call for inbound calls or the number dialed for outbound calls. E.164 format.
	CallerNumber nullable.Type[string] `json:"callerNumber,omitempty"`

	// Amount of money or cost of the call that is charged to your account.
	Charge nullable.Type[float64] `json:"charge,omitempty"`

	// ID of the audio conference.
	ConferenceId nullable.Type[string] `json:"conferenceId,omitempty"`

	// Connection fee price.
	ConnectionCharge nullable.Type[float64] `json:"connectionCharge,omitempty"`

	// Type of currency used to calculate the cost of the call. For details, see (ISO 4217.
	Currency nullable.Type[string] `json:"currency,omitempty"`

	// Whether the call was domestic (within a country or region) or international (outside a country or region), based on
	// the user's location.
	DestinationContext nullable.Type[string] `json:"destinationContext,omitempty"`

	// Country or region dialed.
	DestinationName nullable.Type[string] `json:"destinationName,omitempty"`

	// How long the call was connected, in seconds.
	Duration nullable.Type[int64] `json:"duration,omitempty"`

	// Call end time.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Unique call identifier. GUID.
	Id nullable.Type[string] `json:"id,omitempty"`

	// User's phone number type, such as a service of toll-free number.
	InventoryType nullable.Type[string] `json:"inventoryType,omitempty"`

	// The license used for the call.
	LicenseCapability nullable.Type[string] `json:"licenseCapability,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The telecommunications operator which provided PSTN services for this call. This might be Microsoft, or it might be a
	// third-party operator via the Operator Connect Program.
	Operator nullable.Type[string] `json:"operator,omitempty"`

	// Call start time.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Country code of the tenant. For details, see ISO 3166-1 alpha-2.
	TenantCountryCode nullable.Type[string] `json:"tenantCountryCode,omitempty"`

	// Country code of the user. For details, see ISO 3166-1 alpha-2.
	UsageCountryCode nullable.Type[string] `json:"usageCountryCode,omitempty"`

	// Display name of the user.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// Calling user's ID in Microsoft Graph. GUID. This and other user info will be null/empty for bot call types (ucapin,
	// ucapout).
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user principal name (sign-in name) in Microsoft Entra ID. This is usually the same as the user's SIP address, and
	// can be the same as the user's email address.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
