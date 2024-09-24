package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseOperation interface {
	Entity
	SecurityCaseOperation() BaseSecurityCaseOperationImpl
}

var _ SecurityCaseOperation = BaseSecurityCaseOperationImpl{}

type BaseSecurityCaseOperationImpl struct {
	// The type of action the operation represents. Possible values are:
	// addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
	Action *SecurityCaseAction `json:"action,omitempty"`

	// The date and time the operation was completed.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The user that created the operation.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the operation was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The progress of the operation.
	PercentProgress nullable.Type[int64] `json:"percentProgress,omitempty"`

	// Contains success and failure-specific result information.
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`

	// The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded,
	// partiallySucceeded, failed.
	Status *SecurityCaseOperationStatus `json:"status,omitempty"`

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

func (s BaseSecurityCaseOperationImpl) SecurityCaseOperation() BaseSecurityCaseOperationImpl {
	return s
}

func (s BaseSecurityCaseOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityCaseOperation = RawSecurityCaseOperationImpl{}

// RawSecurityCaseOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityCaseOperationImpl struct {
	securityCaseOperation BaseSecurityCaseOperationImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawSecurityCaseOperationImpl) SecurityCaseOperation() BaseSecurityCaseOperationImpl {
	return s.securityCaseOperation
}

func (s RawSecurityCaseOperationImpl) Entity() BaseEntityImpl {
	return s.securityCaseOperation.Entity()
}

var _ json.Marshaler = BaseSecurityCaseOperationImpl{}

func (s BaseSecurityCaseOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityCaseOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityCaseOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityCaseOperationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.caseOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityCaseOperationImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityCaseOperationImpl{}

func (s *BaseSecurityCaseOperationImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action            *SecurityCaseAction          `json:"action,omitempty"`
		CompletedDateTime nullable.Type[string]        `json:"completedDateTime,omitempty"`
		CreatedDateTime   nullable.Type[string]        `json:"createdDateTime,omitempty"`
		PercentProgress   nullable.Type[int64]         `json:"percentProgress,omitempty"`
		ResultInfo        *ResultInfo                  `json:"resultInfo,omitempty"`
		Status            *SecurityCaseOperationStatus `json:"status,omitempty"`
		Id                *string                      `json:"id,omitempty"`
		ODataId           *string                      `json:"@odata.id,omitempty"`
		ODataType         *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.PercentProgress = decoded.PercentProgress
	s.ResultInfo = decoded.ResultInfo
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecurityCaseOperationImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseSecurityCaseOperationImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalSecurityCaseOperationImplementation(input []byte) (SecurityCaseOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityCaseOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryAddToReviewSetOperation") {
		var out SecurityEdiscoveryAddToReviewSetOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryAddToReviewSetOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryEstimateOperation") {
		var out SecurityEdiscoveryEstimateOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryEstimateOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryExportOperation") {
		var out SecurityEdiscoveryExportOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryExportOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryHoldOperation") {
		var out SecurityEdiscoveryHoldOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryHoldOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryIndexOperation") {
		var out SecurityEdiscoveryIndexOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryIndexOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryPurgeDataOperation") {
		var out SecurityEdiscoveryPurgeDataOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryPurgeDataOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryTagOperation") {
		var out SecurityEdiscoveryTagOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryTagOperation: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityCaseOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityCaseOperationImpl: %+v", err)
	}

	return RawSecurityCaseOperationImpl{
		securityCaseOperation: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
