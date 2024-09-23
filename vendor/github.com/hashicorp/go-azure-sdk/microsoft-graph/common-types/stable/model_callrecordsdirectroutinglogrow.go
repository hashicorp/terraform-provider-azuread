package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsDirectRoutingLogRow struct {
	// In addition to the SIP codes, Microsoft has subcodes that indicate the specific issue.
	CallEndSubReason nullable.Type[int64] `json:"callEndSubReason,omitempty"`

	// Call type and direction.
	CallType nullable.Type[string] `json:"callType,omitempty"`

	// Number of the user or bot who received the call. E.164 format, but might include other data.
	CalleeNumber nullable.Type[string] `json:"calleeNumber,omitempty"`

	// Number of the user or bot who made the call. E.164 format, but might include other data.
	CallerNumber nullable.Type[string] `json:"callerNumber,omitempty"`

	// Identifier for the call that you can use when calling Microsoft Support. GUID.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Duration of the call in seconds.
	Duration nullable.Type[int64] `json:"duration,omitempty"`

	// Only exists for successful (fully established) calls. Time when call ended.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Only exists for failed (not fully established) calls.
	FailureDateTime nullable.Type[string] `json:"failureDateTime,omitempty"`

	// The final response code with which the call ended. For more information, see RFC 3261.
	FinalSipCode nullable.Type[int64] `json:"finalSipCode,omitempty"`

	// Description of the SIP code and Microsoft subcode.
	FinalSipCodePhrase nullable.Type[string] `json:"finalSipCodePhrase,omitempty"`

	// Unique call identifier. GUID.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The date and time when the initial invite was sent.
	InviteDateTime nullable.Type[string] `json:"inviteDateTime,omitempty"`

	// Indicates whether the trunk was enabled for media bypass.
	MediaBypassEnabled nullable.Type[bool] `json:"mediaBypassEnabled,omitempty"`

	// The datacenter used for media path in a nonbypass call.
	MediaPathLocation nullable.Type[string] `json:"mediaPathLocation,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The datacenter used for signaling for both bypass and nonbypass calls.
	SignalingLocation nullable.Type[string] `json:"signalingLocation,omitempty"`

	// Call start time.For failed and unanswered calls, this value can be equal to the invite or failure time.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Success or attempt.
	SuccessfulCall nullable.Type[bool] `json:"successfulCall,omitempty"`

	// Fully qualified domain name of the session border controller.
	TrunkFullyQualifiedDomainName nullable.Type[string] `json:"trunkFullyQualifiedDomainName,omitempty"`

	// Display name of the user.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// Calling user's ID in Microsoft Graph. This and other user information is null/empty for bot call types. GUID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// UserPrincipalName (sign-in name) in Microsoft Entra ID. This value is usually the same as the user's SIP Address, and
	// can be the same as the user's email address.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
