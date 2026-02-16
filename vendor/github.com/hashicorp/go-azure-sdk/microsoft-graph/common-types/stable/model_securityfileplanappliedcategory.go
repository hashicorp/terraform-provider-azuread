package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityFilePlanDescriptorBase = SecurityFilePlanAppliedCategory{}

type SecurityFilePlanAppliedCategory struct {
	// Represents the file plan descriptor for a subcategory under a specific category, which has been assigned to a
	// particular retention label.
	Subcategory *SecurityFilePlanSubcategory `json:"subcategory,omitempty"`

	// Fields inherited from SecurityFilePlanDescriptorBase

	// Unique string that defines the name for the file plan descriptor associated with a particular retention label.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityFilePlanAppliedCategory) SecurityFilePlanDescriptorBase() BaseSecurityFilePlanDescriptorBaseImpl {
	return BaseSecurityFilePlanDescriptorBaseImpl{
		DisplayName: s.DisplayName,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = SecurityFilePlanAppliedCategory{}

func (s SecurityFilePlanAppliedCategory) MarshalJSON() ([]byte, error) {
	type wrapper SecurityFilePlanAppliedCategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityFilePlanAppliedCategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityFilePlanAppliedCategory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.filePlanAppliedCategory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityFilePlanAppliedCategory: %+v", err)
	}

	return encoded, nil
}
