package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityLabelsRoot{}

type SecurityLabelsRoot struct {
	// Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
	Authorities *[]SecurityAuthorityTemplate `json:"authorities,omitempty"`

	// Specifies a group of similar types of content in a particular department.
	Categories *[]SecurityCategoryTemplate `json:"categories,omitempty"`

	// The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should
	// be retained or deleted.
	Citations *[]SecurityCitationTemplate `json:"citations,omitempty"`

	// Specifies the department or business unit of an organization to which a label belongs.
	Departments *[]SecurityDepartmentTemplate `json:"departments,omitempty"`

	// Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
	FilePlanReferences *[]SecurityFilePlanReferenceTemplate `json:"filePlanReferences,omitempty"`

	// Represents how customers can manage their data, whether and for how long to retain or delete it.
	RetentionLabels *[]SecurityRetentionLabel `json:"retentionLabels,omitempty"`

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

func (s SecurityLabelsRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityLabelsRoot{}

func (s SecurityLabelsRoot) MarshalJSON() ([]byte, error) {
	type wrapper SecurityLabelsRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityLabelsRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityLabelsRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.labelsRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityLabelsRoot: %+v", err)
	}

	return encoded, nil
}
