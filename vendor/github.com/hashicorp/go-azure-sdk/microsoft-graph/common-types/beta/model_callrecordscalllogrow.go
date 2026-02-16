package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsCallLogRow interface {
	CallRecordsCallLogRow() BaseCallRecordsCallLogRowImpl
}

var _ CallRecordsCallLogRow = BaseCallRecordsCallLogRowImpl{}

type BaseCallRecordsCallLogRowImpl struct {
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

func (s BaseCallRecordsCallLogRowImpl) CallRecordsCallLogRow() BaseCallRecordsCallLogRowImpl {
	return s
}

var _ CallRecordsCallLogRow = RawCallRecordsCallLogRowImpl{}

// RawCallRecordsCallLogRowImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCallRecordsCallLogRowImpl struct {
	callRecordsCallLogRow BaseCallRecordsCallLogRowImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawCallRecordsCallLogRowImpl) CallRecordsCallLogRow() BaseCallRecordsCallLogRowImpl {
	return s.callRecordsCallLogRow
}

func UnmarshalCallRecordsCallLogRowImplementation(input []byte) (CallRecordsCallLogRow, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsCallLogRow into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.directRoutingLogRow") {
		var out CallRecordsDirectRoutingLogRow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsDirectRoutingLogRow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.pstnCallLogRow") {
		var out CallRecordsPstnCallLogRow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsPstnCallLogRow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.smsLogRow") {
		var out CallRecordsSmsLogRow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsSmsLogRow: %+v", err)
		}
		return out, nil
	}

	var parent BaseCallRecordsCallLogRowImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCallRecordsCallLogRowImpl: %+v", err)
	}

	return RawCallRecordsCallLogRowImpl{
		callRecordsCallLogRow: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
