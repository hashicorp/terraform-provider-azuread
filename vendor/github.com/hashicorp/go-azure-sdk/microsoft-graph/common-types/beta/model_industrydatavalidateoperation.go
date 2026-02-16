package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataValidateOperation interface {
	Entity
	LongRunningOperation
	IndustryDataValidateOperation() BaseIndustryDataValidateOperationImpl
}

var _ IndustryDataValidateOperation = BaseIndustryDataValidateOperationImpl{}

type BaseIndustryDataValidateOperationImpl struct {
	// Set of errors discovered through validation.
	Errors *[]PublicError `json:"errors,omitempty"`

	// Set of warnings discovered through validation.
	Warnings *[]PublicError `json:"warnings,omitempty"`

	// Fields inherited from LongRunningOperation

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

func (s BaseIndustryDataValidateOperationImpl) IndustryDataValidateOperation() BaseIndustryDataValidateOperationImpl {
	return s
}

func (s BaseIndustryDataValidateOperationImpl) LongRunningOperation() BaseLongRunningOperationImpl {
	return BaseLongRunningOperationImpl{
		CreatedDateTime:    s.CreatedDateTime,
		LastActionDateTime: s.LastActionDateTime,
		ResourceLocation:   s.ResourceLocation,
		Status:             s.Status,
		StatusDetail:       s.StatusDetail,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s BaseIndustryDataValidateOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataValidateOperation = RawIndustryDataValidateOperationImpl{}

// RawIndustryDataValidateOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataValidateOperationImpl struct {
	industryDataValidateOperation BaseIndustryDataValidateOperationImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawIndustryDataValidateOperationImpl) IndustryDataValidateOperation() BaseIndustryDataValidateOperationImpl {
	return s.industryDataValidateOperation
}

func (s RawIndustryDataValidateOperationImpl) LongRunningOperation() BaseLongRunningOperationImpl {
	return s.industryDataValidateOperation.LongRunningOperation()
}

func (s RawIndustryDataValidateOperationImpl) Entity() BaseEntityImpl {
	return s.industryDataValidateOperation.Entity()
}

var _ json.Marshaler = BaseIndustryDataValidateOperationImpl{}

func (s BaseIndustryDataValidateOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataValidateOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataValidateOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataValidateOperationImpl: %+v", err)
	}

	delete(decoded, "errors")
	delete(decoded, "warnings")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.validateOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataValidateOperationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataValidateOperationImplementation(input []byte) (IndustryDataValidateOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataValidateOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.fileValidateOperation") {
		var out IndustryDataFileValidateOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataFileValidateOperation: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataValidateOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataValidateOperationImpl: %+v", err)
	}

	return RawIndustryDataValidateOperationImpl{
		industryDataValidateOperation: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
