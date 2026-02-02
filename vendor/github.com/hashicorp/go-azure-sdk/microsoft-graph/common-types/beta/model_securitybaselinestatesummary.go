package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineStateSummary interface {
	Entity
	SecurityBaselineStateSummary() BaseSecurityBaselineStateSummaryImpl
}

var _ SecurityBaselineStateSummary = BaseSecurityBaselineStateSummaryImpl{}

type BaseSecurityBaselineStateSummaryImpl struct {
	// Number of conflict devices
	ConflictCount *int64 `json:"conflictCount,omitempty"`

	// Number of error devices
	ErrorCount *int64 `json:"errorCount,omitempty"`

	// Number of not applicable devices
	NotApplicableCount *int64 `json:"notApplicableCount,omitempty"`

	// Number of not secure devices
	NotSecureCount *int64 `json:"notSecureCount,omitempty"`

	// Number of secure devices
	SecureCount *int64 `json:"secureCount,omitempty"`

	// Number of unknown devices
	UnknownCount *int64 `json:"unknownCount,omitempty"`

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

func (s BaseSecurityBaselineStateSummaryImpl) SecurityBaselineStateSummary() BaseSecurityBaselineStateSummaryImpl {
	return s
}

func (s BaseSecurityBaselineStateSummaryImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityBaselineStateSummary = RawSecurityBaselineStateSummaryImpl{}

// RawSecurityBaselineStateSummaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityBaselineStateSummaryImpl struct {
	securityBaselineStateSummary BaseSecurityBaselineStateSummaryImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawSecurityBaselineStateSummaryImpl) SecurityBaselineStateSummary() BaseSecurityBaselineStateSummaryImpl {
	return s.securityBaselineStateSummary
}

func (s RawSecurityBaselineStateSummaryImpl) Entity() BaseEntityImpl {
	return s.securityBaselineStateSummary.Entity()
}

var _ json.Marshaler = BaseSecurityBaselineStateSummaryImpl{}

func (s BaseSecurityBaselineStateSummaryImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityBaselineStateSummaryImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityBaselineStateSummaryImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityBaselineStateSummaryImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.securityBaselineStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityBaselineStateSummaryImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityBaselineStateSummaryImplementation(input []byte) (SecurityBaselineStateSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityBaselineStateSummary into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineCategoryStateSummary") {
		var out SecurityBaselineCategoryStateSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineCategoryStateSummary: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityBaselineStateSummaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityBaselineStateSummaryImpl: %+v", err)
	}

	return RawSecurityBaselineStateSummaryImpl{
		securityBaselineStateSummary: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
