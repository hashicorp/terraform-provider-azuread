package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityFilePlanDescriptor{}

type SecurityFilePlanDescriptor struct {
	// Represents the file plan descriptor of type authority applied to a particular retention label.
	Authority *SecurityFilePlanAuthority `json:"authority,omitempty"`

	// Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
	AuthorityTemplate *SecurityAuthorityTemplate `json:"authorityTemplate,omitempty"`

	// Represents the file plan descriptor of type category applied to a particular retention label.
	Category *SecurityFilePlanAppliedCategory `json:"category,omitempty"`

	// Specifies a group of similar types of content in a particular department.
	CategoryTemplate *SecurityCategoryTemplate `json:"categoryTemplate,omitempty"`

	// Represents the file plan descriptor of type citation applied to a particular retention label.
	Citation *SecurityFilePlanCitation `json:"citation,omitempty"`

	// The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should
	// be retained or deleted.
	CitationTemplate *SecurityCitationTemplate `json:"citationTemplate,omitempty"`

	// Represents the file plan descriptor of type department applied to a particular retention label.
	Department *SecurityFilePlanDepartment `json:"department,omitempty"`

	// Specifies the department or business unit of an organization to which a label belongs.
	DepartmentTemplate *SecurityDepartmentTemplate `json:"departmentTemplate,omitempty"`

	// Represents the file plan descriptor of type filePlanReference applied to a particular retention label.
	FilePlanReference *SecurityFilePlanReference `json:"filePlanReference,omitempty"`

	// Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
	FilePlanReferenceTemplate *SecurityFilePlanReferenceTemplate `json:"filePlanReferenceTemplate,omitempty"`

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

func (s SecurityFilePlanDescriptor) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityFilePlanDescriptor{}

func (s SecurityFilePlanDescriptor) MarshalJSON() ([]byte, error) {
	type wrapper SecurityFilePlanDescriptor
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityFilePlanDescriptor: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityFilePlanDescriptor: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.filePlanDescriptor"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityFilePlanDescriptor: %+v", err)
	}

	return encoded, nil
}
