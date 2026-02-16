package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerDelta = PlannerTaskDetails{}

type PlannerTaskDetails struct {
	// Detailed information about the approval that is attached to the task.
	ApprovalAttachment PlannerBaseApprovalAttachment `json:"approvalAttachment"`

	// The collection of checklist items on the task.
	Checklist *PlannerChecklistItems `json:"checklist,omitempty"`

	// Contains detailed information about requirements on the task.
	CompletionRequirements *PlannerTaskCompletionRequirementDetails `json:"completionRequirements,omitempty"`

	// Description of the task.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Read-only. Represents a dictionary of data about the forms associated with a task. Each entry in the dictionary is a
	// key-value pair, and the value is a plannerFormReference object.
	Forms *PlannerFormsDictionary `json:"forms,omitempty"`

	// Rich text description of the task. To be used by HTML-aware clients. For backwards compatibility, a plain-text
	// version of the HTML description will be synced to the 'description' field. If this field hasn't previously been set
	// but 'description' has been, the existing description is synchronized to 'notes' with minimal whitespace-preserving
	// HTML markup. Setting both 'description' and 'notes' is an error and will result in an exception.
	Notes *ItemBody `json:"notes,omitempty"`

	// This sets the type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist,
	// description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
	PreviewType *PlannerPreviewType `json:"previewType,omitempty"`

	// The collection of references on the task.
	References *PlannerExternalReferences `json:"references,omitempty"`

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

func (s PlannerTaskDetails) PlannerDelta() BasePlannerDeltaImpl {
	return BasePlannerDeltaImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s PlannerTaskDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerTaskDetails{}

func (s PlannerTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper PlannerTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerTaskDetails: %+v", err)
	}

	delete(decoded, "forms")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerTaskDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerTaskDetails: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerTaskDetails{}

func (s *PlannerTaskDetails) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Checklist              *PlannerChecklistItems                   `json:"checklist,omitempty"`
		CompletionRequirements *PlannerTaskCompletionRequirementDetails `json:"completionRequirements,omitempty"`
		Description            nullable.Type[string]                    `json:"description,omitempty"`
		Forms                  *PlannerFormsDictionary                  `json:"forms,omitempty"`
		Notes                  *ItemBody                                `json:"notes,omitempty"`
		PreviewType            *PlannerPreviewType                      `json:"previewType,omitempty"`
		References             *PlannerExternalReferences               `json:"references,omitempty"`
		Id                     *string                                  `json:"id,omitempty"`
		ODataId                *string                                  `json:"@odata.id,omitempty"`
		ODataType              *string                                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Checklist = decoded.Checklist
	s.CompletionRequirements = decoded.CompletionRequirements
	s.Description = decoded.Description
	s.Forms = decoded.Forms
	s.Notes = decoded.Notes
	s.PreviewType = decoded.PreviewType
	s.References = decoded.References
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerTaskDetails into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["approvalAttachment"]; ok {
		impl, err := UnmarshalPlannerBaseApprovalAttachmentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ApprovalAttachment' for 'PlannerTaskDetails': %+v", err)
		}
		s.ApprovalAttachment = impl
	}

	return nil
}
