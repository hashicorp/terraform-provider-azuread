package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobResponseBase interface {
	Entity
	JobResponseBase() BaseJobResponseBaseImpl
}

var _ JobResponseBase = BaseJobResponseBaseImpl{}

type BaseJobResponseBaseImpl struct {
	CreationDateTime nullable.Type[string] `json:"creationDateTime,omitempty"`
	EndDateTime      nullable.Type[string] `json:"endDateTime,omitempty"`
	Error            *ClassificationError  `json:"error,omitempty"`
	StartDateTime    nullable.Type[string] `json:"startDateTime,omitempty"`
	Status           nullable.Type[string] `json:"status,omitempty"`
	TenantId         nullable.Type[string] `json:"tenantId,omitempty"`
	Type             nullable.Type[string] `json:"type,omitempty"`
	UserId           nullable.Type[string] `json:"userId,omitempty"`

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

func (s BaseJobResponseBaseImpl) JobResponseBase() BaseJobResponseBaseImpl {
	return s
}

func (s BaseJobResponseBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ JobResponseBase = RawJobResponseBaseImpl{}

// RawJobResponseBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobResponseBaseImpl struct {
	jobResponseBase BaseJobResponseBaseImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawJobResponseBaseImpl) JobResponseBase() BaseJobResponseBaseImpl {
	return s.jobResponseBase
}

func (s RawJobResponseBaseImpl) Entity() BaseEntityImpl {
	return s.jobResponseBase.Entity()
}

var _ json.Marshaler = BaseJobResponseBaseImpl{}

func (s BaseJobResponseBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseJobResponseBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseJobResponseBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseJobResponseBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.jobResponseBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseJobResponseBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalJobResponseBaseImplementation(input []byte) (JobResponseBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobResponseBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.classificationJobResponse") {
		var out ClassificationJobResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClassificationJobResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.dlpEvaluatePoliciesJobResponse") {
		var out DlpEvaluatePoliciesJobResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DlpEvaluatePoliciesJobResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.evaluateLabelJobResponse") {
		var out EvaluateLabelJobResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EvaluateLabelJobResponse: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobResponseBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobResponseBaseImpl: %+v", err)
	}

	return RawJobResponseBaseImpl{
		jobResponseBase: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
