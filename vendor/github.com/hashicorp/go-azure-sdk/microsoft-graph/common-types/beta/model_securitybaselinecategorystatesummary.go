package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityBaselineStateSummary = SecurityBaselineCategoryStateSummary{}

type SecurityBaselineCategoryStateSummary struct {
	// The category name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Fields inherited from SecurityBaselineStateSummary

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

func (s SecurityBaselineCategoryStateSummary) SecurityBaselineStateSummary() BaseSecurityBaselineStateSummaryImpl {
	return BaseSecurityBaselineStateSummaryImpl{
		ConflictCount:      s.ConflictCount,
		ErrorCount:         s.ErrorCount,
		NotApplicableCount: s.NotApplicableCount,
		NotSecureCount:     s.NotSecureCount,
		SecureCount:        s.SecureCount,
		UnknownCount:       s.UnknownCount,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s SecurityBaselineCategoryStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityBaselineCategoryStateSummary{}

func (s SecurityBaselineCategoryStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper SecurityBaselineCategoryStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityBaselineCategoryStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityBaselineCategoryStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.securityBaselineCategoryStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityBaselineCategoryStateSummary: %+v", err)
	}

	return encoded, nil
}
