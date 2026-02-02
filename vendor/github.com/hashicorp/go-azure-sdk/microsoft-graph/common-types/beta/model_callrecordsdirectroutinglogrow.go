package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsCallLogRow = CallRecordsDirectRoutingLogRow{}

type CallRecordsDirectRoutingLogRow struct {
	// In addition to the SIP codes, Microsoft has own subcodes that indicate the specific issue.
	CallEndSubReason nullable.Type[int64] `json:"callEndSubReason,omitempty"`

	// Call type and direction.
	CallType nullable.Type[string] `json:"callType,omitempty"`

	// Number of the user or bot who received the call (E.164 format, but might include more data).
	CalleeNumber nullable.Type[string] `json:"calleeNumber,omitempty"`

	// Number of the user or bot who made the call (E.164 format, but might include more data).
	CallerNumber nullable.Type[string] `json:"callerNumber,omitempty"`

	// Identifier (GUID) for the call that you can use when calling Microsoft Support.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Duration of the call in seconds.
	Duration nullable.Type[int64] `json:"duration,omitempty"`

	// Only exists for successful (fully established) calls. The time when the call ended.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Only exists for failed (not fully established) calls.
	FailureDateTime nullable.Type[string] `json:"failureDateTime,omitempty"`

	// The final response code with which the call ended (RFC 3261).
	FinalSipCode nullable.Type[int64] `json:"finalSipCode,omitempty"`

	// Description of the SIP code and Microsoft subcode.
	FinalSipCodePhrase nullable.Type[string] `json:"finalSipCodePhrase,omitempty"`

	// The date and time when the initial invite was sent.
	InviteDateTime nullable.Type[string] `json:"inviteDateTime,omitempty"`

	// Indicates if the trunk was enabled for media bypass or not.
	MediaBypassEnabled nullable.Type[bool] `json:"mediaBypassEnabled,omitempty"`

	// The data center used for media path in non-bypass call.
	MediaPathLocation nullable.Type[string] `json:"mediaPathLocation,omitempty"`

	// The data center used for signaling for both bypass and non-bypass calls.
	SignalingLocation nullable.Type[string] `json:"signalingLocation,omitempty"`

	// Call start time.For failed and unanswered calls, this value can be equal to invite or failure time.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Success or attempt.
	SuccessfulCall nullable.Type[bool] `json:"successfulCall,omitempty"`

	// Correlation ID of the call to the transferor.
	TransferorCorrelationId nullable.Type[string] `json:"transferorCorrelationId,omitempty"`

	// Fully qualified domain name of the session border controller.
	TrunkFullyQualifiedDomainName nullable.Type[string] `json:"trunkFullyQualifiedDomainName,omitempty"`

	// Country/region code of the user. For details, see ISO 3166-1 alpha-2.
	UserCountryCode nullable.Type[string] `json:"userCountryCode,omitempty"`

	// Fields inherited from CallRecordsCallLogRow

	AdministrativeUnitInfos *[]CallRecordsAdministrativeUnitInfo `json:"administrativeUnitInfos,omitempty"`
	Id                      nullable.Type[string]                `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	OtherPartyCountryCode nullable.Type[string] `json:"otherPartyCountryCode,omitempty"`
	UserDisplayName       nullable.Type[string] `json:"userDisplayName,omitempty"`
	UserId                nullable.Type[string] `json:"userId,omitempty"`
	UserPrincipalName     nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CallRecordsDirectRoutingLogRow) CallRecordsCallLogRow() BaseCallRecordsCallLogRowImpl {
	return BaseCallRecordsCallLogRowImpl{
		AdministrativeUnitInfos: s.AdministrativeUnitInfos,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
		OtherPartyCountryCode:   s.OtherPartyCountryCode,
		UserDisplayName:         s.UserDisplayName,
		UserId:                  s.UserId,
		UserPrincipalName:       s.UserPrincipalName,
	}
}

var _ json.Marshaler = CallRecordsDirectRoutingLogRow{}

func (s CallRecordsDirectRoutingLogRow) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsDirectRoutingLogRow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsDirectRoutingLogRow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsDirectRoutingLogRow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.directRoutingLogRow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsDirectRoutingLogRow: %+v", err)
	}

	return encoded, nil
}
