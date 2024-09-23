package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnersBillingOperation interface {
	Entity
	PartnersBillingOperation() BasePartnersBillingOperationImpl
}

var _ PartnersBillingOperation = BasePartnersBillingOperationImpl{}

type BasePartnersBillingOperationImpl struct {
	// The start time of the operation. The timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The time of the last action of the operation. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastActionDateTime *string `json:"lastActionDateTime,omitempty"`

	// The status of the operation. Possible values are: notStarted, running, completed, failed, unknownFutureValue.
	Status *LongRunningOperationStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePartnersBillingOperationImpl) PartnersBillingOperation() BasePartnersBillingOperationImpl {
	return s
}

func (s BasePartnersBillingOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PartnersBillingOperation = RawPartnersBillingOperationImpl{}

// RawPartnersBillingOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPartnersBillingOperationImpl struct {
	partnersBillingOperation BasePartnersBillingOperationImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawPartnersBillingOperationImpl) PartnersBillingOperation() BasePartnersBillingOperationImpl {
	return s.partnersBillingOperation
}

func (s RawPartnersBillingOperationImpl) Entity() BaseEntityImpl {
	return s.partnersBillingOperation.Entity()
}

var _ json.Marshaler = BasePartnersBillingOperationImpl{}

func (s BasePartnersBillingOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePartnersBillingOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePartnersBillingOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePartnersBillingOperationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partners.billing.operation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePartnersBillingOperationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPartnersBillingOperationImplementation(input []byte) (PartnersBillingOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnersBillingOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.exportSuccessOperation") {
		var out PartnersBillingExportSuccessOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingExportSuccessOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.failedOperation") {
		var out PartnersBillingFailedOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingFailedOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partners.billing.runningOperation") {
		var out PartnersBillingRunningOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnersBillingRunningOperation: %+v", err)
		}
		return out, nil
	}

	var parent BasePartnersBillingOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePartnersBillingOperationImpl: %+v", err)
	}

	return RawPartnersBillingOperationImpl{
		partnersBillingOperation: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
