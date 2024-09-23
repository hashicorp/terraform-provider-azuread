package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsCallLogRow = CallRecordsSmsLogRow{}

type CallRecordsSmsLogRow struct {
	// Amount of money or cost of the SMS that is charged.
	CallCharge nullable.Type[float64] `json:"callCharge,omitempty"`

	// Currency used to calculate the cost of the call. For details, see ISO 4217.
	Currency nullable.Type[string] `json:"currency,omitempty"`

	// Indicates whether the SMS was Domestic (within a country or region) or International (outside a country or region)
	// based on the user's location.
	DestinationContext nullable.Type[string] `json:"destinationContext,omitempty"`

	// Country or region of a phone number that received the SMS.
	DestinationName nullable.Type[string] `json:"destinationName,omitempty"`

	// Partially obfuscated phone number that received the SMS. For details, see E.164.
	DestinationNumber nullable.Type[string] `json:"destinationNumber,omitempty"`

	// The license used for the SMS.
	LicenseCapability nullable.Type[string] `json:"licenseCapability,omitempty"`

	// The date and time when the SMS was sent.
	SentDateTime nullable.Type[string] `json:"sentDateTime,omitempty"`

	// SMS identifier. Not guaranteed to be unique.
	SmsId nullable.Type[string] `json:"smsId,omitempty"`

	// Type of SMS such as outbound or inbound.
	SmsType nullable.Type[string] `json:"smsType,omitempty"`

	// Number of SMS units sent/received.
	SmsUnits nullable.Type[int64] `json:"smsUnits,omitempty"`

	// Partially obfuscated phone number that sent the SMS. For details, see E.164.
	SourceNumber nullable.Type[string] `json:"sourceNumber,omitempty"`

	// Country code of the tenant. For details, see ISO 3166-1 alpha-2.
	TenantCountryCode nullable.Type[string] `json:"tenantCountryCode,omitempty"`

	// Country code of the user. For details, see ISO 3166-1 alpha-2.
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

func (s CallRecordsSmsLogRow) CallRecordsCallLogRow() BaseCallRecordsCallLogRowImpl {
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

var _ json.Marshaler = CallRecordsSmsLogRow{}

func (s CallRecordsSmsLogRow) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsSmsLogRow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsSmsLogRow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsSmsLogRow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.smsLogRow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsSmsLogRow: %+v", err)
	}

	return encoded, nil
}
