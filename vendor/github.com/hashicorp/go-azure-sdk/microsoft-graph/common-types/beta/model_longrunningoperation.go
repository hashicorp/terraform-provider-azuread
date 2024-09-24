package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LongRunningOperation interface {
	Entity
	LongRunningOperation() BaseLongRunningOperationImpl
}

var _ LongRunningOperation = BaseLongRunningOperationImpl{}

type BaseLongRunningOperationImpl struct {
	// The start time of the operation. The timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The time of the last action in the operation. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// URI of the resource that the operation is performed on.
	ResourceLocation nullable.Type[string] `json:"resourceLocation,omitempty"`

	// The status of the operation. The possible values are: notStarted, running, succeeded, failed, skipped,
	// unknownFutureValue.
	Status *LongRunningOperationStatus `json:"status,omitempty"`

	// Details about the status of the operation.
	StatusDetail nullable.Type[string] `json:"statusDetail,omitempty"`

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

func (s BaseLongRunningOperationImpl) LongRunningOperation() BaseLongRunningOperationImpl {
	return s
}

func (s BaseLongRunningOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ LongRunningOperation = RawLongRunningOperationImpl{}

// RawLongRunningOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLongRunningOperationImpl struct {
	longRunningOperation BaseLongRunningOperationImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawLongRunningOperationImpl) LongRunningOperation() BaseLongRunningOperationImpl {
	return s.longRunningOperation
}

func (s RawLongRunningOperationImpl) Entity() BaseEntityImpl {
	return s.longRunningOperation.Entity()
}

var _ json.Marshaler = BaseLongRunningOperationImpl{}

func (s BaseLongRunningOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseLongRunningOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseLongRunningOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseLongRunningOperationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.longRunningOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseLongRunningOperationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalLongRunningOperationImplementation(input []byte) (LongRunningOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling LongRunningOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.attackSimulationOperation") {
		var out AttackSimulationOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttackSimulationOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.engagementAsyncOperation") {
		var out EngagementAsyncOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EngagementAsyncOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.goalsExportJob") {
		var out GoalsExportJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GoalsExportJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.validateOperation") {
		var out IndustryDataValidateOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataValidateOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.richLongRunningOperation") {
		var out RichLongRunningOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RichLongRunningOperation: %+v", err)
		}
		return out, nil
	}

	var parent BaseLongRunningOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseLongRunningOperationImpl: %+v", err)
	}

	return RawLongRunningOperationImpl{
		longRunningOperation: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
