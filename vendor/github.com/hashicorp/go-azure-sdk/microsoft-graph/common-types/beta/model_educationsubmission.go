package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationSubmission{}

type EducationSubmission struct {
	ExcusedBy       *IdentitySet          `json:"excusedBy,omitempty"`
	ExcusedDateTime nullable.Type[string] `json:"excusedDateTime,omitempty"`
	Outcomes        *[]EducationOutcome   `json:"outcomes,omitempty"`

	// User who moved the status of this submission to reassigned.
	ReassignedBy *IdentitySet `json:"reassignedBy,omitempty"`

	// Moment in time when the submission was reassigned. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ReassignedDateTime nullable.Type[string] `json:"reassignedDateTime,omitempty"`

	// Who this submission is assigned to.
	Recipient EducationSubmissionRecipient `json:"recipient"`

	Resources *[]EducationSubmissionResource `json:"resources,omitempty"`

	// Folder where all file resources for this submission need to be stored.
	ResourcesFolderUrl nullable.Type[string] `json:"resourcesFolderUrl,omitempty"`

	// User who moved the status of this submission to returned.
	ReturnedBy *IdentitySet `json:"returnedBy,omitempty"`

	// Moment in time when the submission was returned. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ReturnedDateTime nullable.Type[string] `json:"returnedDateTime,omitempty"`

	// Read-only. Possible values are: working, submitted, returned, unknownFutureValue, reassigned, and excused. Note that
	// you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable
	// enum: reassigned, and excused.
	Status *EducationSubmissionStatus `json:"status,omitempty"`

	// User who moved the resource into the submitted state.
	SubmittedBy *IdentitySet `json:"submittedBy,omitempty"`

	// Moment in time when the submission was moved into the submitted state. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	SubmittedDateTime nullable.Type[string] `json:"submittedDateTime,omitempty"`

	SubmittedResources *[]EducationSubmissionResource `json:"submittedResources,omitempty"`

	// User who moved the resource from submitted into the working state.
	UnsubmittedBy *IdentitySet `json:"unsubmittedBy,omitempty"`

	// Moment in time when the submission was moved from submitted into the working state. The timestamp type represents
	// date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	UnsubmittedDateTime nullable.Type[string] `json:"unsubmittedDateTime,omitempty"`

	// The deep link URL for the given submission.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

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

func (s EducationSubmission) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationSubmission{}

func (s EducationSubmission) MarshalJSON() ([]byte, error) {
	type wrapper EducationSubmission
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSubmission: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSubmission: %+v", err)
	}

	delete(decoded, "excusedBy")
	delete(decoded, "excusedDateTime")
	delete(decoded, "reassignedBy")
	delete(decoded, "reassignedDateTime")
	delete(decoded, "resourcesFolderUrl")
	delete(decoded, "returnedBy")
	delete(decoded, "returnedDateTime")
	delete(decoded, "status")
	delete(decoded, "submittedBy")
	delete(decoded, "submittedDateTime")
	delete(decoded, "unsubmittedBy")
	delete(decoded, "unsubmittedDateTime")
	delete(decoded, "webUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSubmission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSubmission: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationSubmission{}

func (s *EducationSubmission) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExcusedDateTime     nullable.Type[string]          `json:"excusedDateTime,omitempty"`
		ReassignedDateTime  nullable.Type[string]          `json:"reassignedDateTime,omitempty"`
		Resources           *[]EducationSubmissionResource `json:"resources,omitempty"`
		ResourcesFolderUrl  nullable.Type[string]          `json:"resourcesFolderUrl,omitempty"`
		ReturnedDateTime    nullable.Type[string]          `json:"returnedDateTime,omitempty"`
		Status              *EducationSubmissionStatus     `json:"status,omitempty"`
		SubmittedDateTime   nullable.Type[string]          `json:"submittedDateTime,omitempty"`
		SubmittedResources  *[]EducationSubmissionResource `json:"submittedResources,omitempty"`
		UnsubmittedDateTime nullable.Type[string]          `json:"unsubmittedDateTime,omitempty"`
		WebUrl              nullable.Type[string]          `json:"webUrl,omitempty"`
		Id                  *string                        `json:"id,omitempty"`
		ODataId             *string                        `json:"@odata.id,omitempty"`
		ODataType           *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExcusedDateTime = decoded.ExcusedDateTime
	s.ReassignedDateTime = decoded.ReassignedDateTime
	s.Resources = decoded.Resources
	s.ResourcesFolderUrl = decoded.ResourcesFolderUrl
	s.ReturnedDateTime = decoded.ReturnedDateTime
	s.Status = decoded.Status
	s.SubmittedDateTime = decoded.SubmittedDateTime
	s.SubmittedResources = decoded.SubmittedResources
	s.UnsubmittedDateTime = decoded.UnsubmittedDateTime
	s.WebUrl = decoded.WebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationSubmission into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["excusedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExcusedBy' for 'EducationSubmission': %+v", err)
		}
		s.ExcusedBy = &impl
	}

	if v, ok := temp["outcomes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Outcomes into list []json.RawMessage: %+v", err)
		}

		output := make([]EducationOutcome, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEducationOutcomeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Outcomes' for 'EducationSubmission': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Outcomes = &output
	}

	if v, ok := temp["reassignedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReassignedBy' for 'EducationSubmission': %+v", err)
		}
		s.ReassignedBy = &impl
	}

	if v, ok := temp["recipient"]; ok {
		impl, err := UnmarshalEducationSubmissionRecipientImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Recipient' for 'EducationSubmission': %+v", err)
		}
		s.Recipient = impl
	}

	if v, ok := temp["returnedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReturnedBy' for 'EducationSubmission': %+v", err)
		}
		s.ReturnedBy = &impl
	}

	if v, ok := temp["submittedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SubmittedBy' for 'EducationSubmission': %+v", err)
		}
		s.SubmittedBy = &impl
	}

	if v, ok := temp["unsubmittedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'UnsubmittedBy' for 'EducationSubmission': %+v", err)
		}
		s.UnsubmittedBy = &impl
	}

	return nil
}
