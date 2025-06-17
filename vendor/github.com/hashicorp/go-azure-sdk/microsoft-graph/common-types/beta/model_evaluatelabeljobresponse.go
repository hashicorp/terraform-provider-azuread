package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobResponseBase = EvaluateLabelJobResponse{}

type EvaluateLabelJobResponse struct {
	Result *EvaluateLabelJobResultGroup `json:"result,omitempty"`

	// Fields inherited from JobResponseBase

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

func (s EvaluateLabelJobResponse) JobResponseBase() BaseJobResponseBaseImpl {
	return BaseJobResponseBaseImpl{
		CreationDateTime: s.CreationDateTime,
		EndDateTime:      s.EndDateTime,
		Error:            s.Error,
		StartDateTime:    s.StartDateTime,
		Status:           s.Status,
		TenantId:         s.TenantId,
		Type:             s.Type,
		UserId:           s.UserId,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s EvaluateLabelJobResponse) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EvaluateLabelJobResponse{}

func (s EvaluateLabelJobResponse) MarshalJSON() ([]byte, error) {
	type wrapper EvaluateLabelJobResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EvaluateLabelJobResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EvaluateLabelJobResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.evaluateLabelJobResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EvaluateLabelJobResponse: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EvaluateLabelJobResponse{}

func (s *EvaluateLabelJobResponse) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Result           *EvaluateLabelJobResultGroup `json:"result,omitempty"`
		CreationDateTime nullable.Type[string]        `json:"creationDateTime,omitempty"`
		EndDateTime      nullable.Type[string]        `json:"endDateTime,omitempty"`
		StartDateTime    nullable.Type[string]        `json:"startDateTime,omitempty"`
		Status           nullable.Type[string]        `json:"status,omitempty"`
		TenantId         nullable.Type[string]        `json:"tenantId,omitempty"`
		Type             nullable.Type[string]        `json:"type,omitempty"`
		UserId           nullable.Type[string]        `json:"userId,omitempty"`
		Id               *string                      `json:"id,omitempty"`
		ODataId          *string                      `json:"@odata.id,omitempty"`
		ODataType        *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Result = decoded.Result
	s.CreationDateTime = decoded.CreationDateTime
	s.EndDateTime = decoded.EndDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.StartDateTime = decoded.StartDateTime
	s.Status = decoded.Status
	s.TenantId = decoded.TenantId
	s.Type = decoded.Type
	s.UserId = decoded.UserId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EvaluateLabelJobResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["error"]; ok {
		impl, err := UnmarshalClassificationErrorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Error' for 'EvaluateLabelJobResponse': %+v", err)
		}
		s.Error = &impl
	}

	return nil
}
