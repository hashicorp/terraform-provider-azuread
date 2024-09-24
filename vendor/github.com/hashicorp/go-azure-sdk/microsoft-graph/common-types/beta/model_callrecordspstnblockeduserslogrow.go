package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnBlockedUsersLogRow struct {
	// The date and time when the user was blocked/unblocked from making PSTN calls. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	BlockDateTime nullable.Type[string] `json:"blockDateTime,omitempty"`

	// The reason why the user is blocked/unblocked from making calls.
	BlockReason nullable.Type[string] `json:"blockReason,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique identifier (GUID) for the blocking/unblocking action.
	RemediationId nullable.Type[string] `json:"remediationId,omitempty"`

	// Indicates whether the user is blocked or unblocked from making PSTN calls in Microsoft Teams. The possible values
	// are: blocked, unblocked, unknownFutureValue.
	UserBlockMode *CallRecordsPstnUserBlockMode `json:"userBlockMode,omitempty"`

	// Display name of the user.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// The unique identifier (GUID) of the user in Microsoft Entra ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user principal name (sign-in name) in Microsoft Entra ID. This is usually the same as the user's SIP address, and
	// can be same as the user's e-mail address.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// User's blocked number. For details, see E.164.
	UserTelephoneNumber nullable.Type[string] `json:"userTelephoneNumber,omitempty"`
}
