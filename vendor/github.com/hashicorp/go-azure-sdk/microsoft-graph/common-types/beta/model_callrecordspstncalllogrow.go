package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsCallLogRow = CallRecordsPstnCallLogRow{}

type CallRecordsPstnCallLogRow struct {
	// The source of the call duration data. If the call uses a third-party telecommunications operator via the Operator
	// Connect Program, the operator may provide their own call duration data. In this case, the property value is operator.
	// Otherwise, the value is microsoft.
	CallDurationSource *CallRecordsPstnCallDurationSource `json:"callDurationSource,omitempty"`

	// Call identifier. Not guaranteed to be unique.
	CallId nullable.Type[string] `json:"callId,omitempty"`

	// Indicates whether the call was a PSTN outbound or inbound call and the type of call such as a call placed by a user
	// or an audio conference.
	CallType nullable.Type[string] `json:"callType,omitempty"`

	// Number of the user or bot who received the call (E.164).
	CalleeNumber nullable.Type[string] `json:"calleeNumber,omitempty"`

	// Number of the user or bot who made the call (E.164).
	CallerNumber nullable.Type[string] `json:"callerNumber,omitempty"`

	// Amount of money or cost of the call that is charged to your account.
	Charge nullable.Type[float64] `json:"charge,omitempty"`

	// Local IPv4 of the client that is retrieved from the operating system of the client.
	ClientLocalIPV4Address nullable.Type[string] `json:"clientLocalIpV4Address,omitempty"`

	// Local IPv6 of the client that is retrieved from the operating system of the client.
	ClientLocalIPV6Address nullable.Type[string] `json:"clientLocalIpV6Address,omitempty"`

	// Public IPv4 of the client that can be used to determine the location of the client.
	ClientPublicIPV4Address nullable.Type[string] `json:"clientPublicIpV4Address,omitempty"`

	// Public IPv6 of the client that can be used to determine the location of the client.
	ClientPublicIPV6Address nullable.Type[string] `json:"clientPublicIpV6Address,omitempty"`

	// ID of the audio conference.
	ConferenceId nullable.Type[string] `json:"conferenceId,omitempty"`

	// Connection fee price.
	ConnectionCharge nullable.Type[float64] `json:"connectionCharge,omitempty"`

	// Type of currency used to calculate the cost of the call (ISO 4217).
	Currency nullable.Type[string] `json:"currency,omitempty"`

	// Indicates whether the call was Domestic (within a country or region) or International (outside a country or region)
	// based on the user's location.
	DestinationContext nullable.Type[string] `json:"destinationContext,omitempty"`

	// Country or region dialed.
	DestinationName nullable.Type[string] `json:"destinationName,omitempty"`

	// How long the call was connected, in seconds.
	Duration nullable.Type[int64] `json:"duration,omitempty"`

	// Call end time.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// User's phone number type, such as a service of toll-free number.
	InventoryType nullable.Type[string] `json:"inventoryType,omitempty"`

	// The license used for the call.
	LicenseCapability nullable.Type[string] `json:"licenseCapability,omitempty"`

	// The telecommunications operator that provided PSTN services for this call. It may be Microsoft, or it may be a
	// third-party operator via the Operator Connect Program.
	Operator nullable.Type[string] `json:"operator,omitempty"`

	// Call start time.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Country code of the tenant. For details, see ISO 3166-1 alpha-2.
	TenantCountryCode nullable.Type[string] `json:"tenantCountryCode,omitempty"`

	// Country code of the user. For details, see ISO 3166-1 alpha-2.
	UsageCountryCode nullable.Type[string] `json:"usageCountryCode,omitempty"`

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

func (s CallRecordsPstnCallLogRow) CallRecordsCallLogRow() BaseCallRecordsCallLogRowImpl {
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

var _ json.Marshaler = CallRecordsPstnCallLogRow{}

func (s CallRecordsPstnCallLogRow) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsPstnCallLogRow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsPstnCallLogRow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsPstnCallLogRow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.pstnCallLogRow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsPstnCallLogRow: %+v", err)
	}

	return encoded, nil
}
