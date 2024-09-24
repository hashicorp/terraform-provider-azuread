package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationClass{}

type EducationClass struct {
	// All categories associated with this class. Nullable.
	AssignmentCategories *[]EducationCategory `json:"assignmentCategories,omitempty"`

	// Specifies class-level defaults respected by new assignments created in the class.
	AssignmentDefaults *EducationAssignmentDefaults `json:"assignmentDefaults,omitempty"`

	// Specifies class-level assignments settings.
	AssignmentSettings *EducationAssignmentSettings `json:"assignmentSettings,omitempty"`

	// All assignments associated with this class. Nullable.
	Assignments *[]EducationAssignment `json:"assignments,omitempty"`

	// Class code used by the school to identify the class.
	ClassCode nullable.Type[string] `json:"classCode,omitempty"`

	Course *EducationCourse `json:"course,omitempty"`

	// Entity who created the class
	CreatedBy IdentitySet `json:"createdBy"`

	// Description of the class.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the class.
	DisplayName *string `json:"displayName,omitempty"`

	// ID of the class from the syncing system.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Name of the class in the syncing system.
	ExternalName nullable.Type[string] `json:"externalName,omitempty"`

	// How this class was created. Possible values are: sis, manual.
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`

	// The name of the external source this resource was generated from.
	ExternalSourceDetail nullable.Type[string] `json:"externalSourceDetail,omitempty"`

	// Grade level of the class.
	Grade nullable.Type[string] `json:"grade,omitempty"`

	// The underlying Microsoft 365 group object.
	Group *Group `json:"group,omitempty"`

	// Mail name for sending email to all members, if this is enabled.
	MailNickname *string `json:"mailNickname,omitempty"`

	// All users in the class. Nullable.
	Members *[]EducationUser `json:"members,omitempty"`

	// All modules in the class. Nullable.
	Modules *[]EducationModule `json:"modules,omitempty"`

	// All schools that this class is associated with. Nullable.
	Schools *[]EducationSchool `json:"schools,omitempty"`

	// All teachers in the class. Nullable.
	Teachers *[]EducationUser `json:"teachers,omitempty"`

	// Term for this class.
	Term *EducationTerm `json:"term,omitempty"`

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

func (s EducationClass) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationClass{}

func (s EducationClass) MarshalJSON() ([]byte, error) {
	type wrapper EducationClass
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationClass: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationClass: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationClass"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationClass: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationClass{}

func (s *EducationClass) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignmentCategories *[]EducationCategory         `json:"assignmentCategories,omitempty"`
		AssignmentDefaults   *EducationAssignmentDefaults `json:"assignmentDefaults,omitempty"`
		AssignmentSettings   *EducationAssignmentSettings `json:"assignmentSettings,omitempty"`
		Assignments          *[]EducationAssignment       `json:"assignments,omitempty"`
		ClassCode            nullable.Type[string]        `json:"classCode,omitempty"`
		Course               *EducationCourse             `json:"course,omitempty"`
		Description          nullable.Type[string]        `json:"description,omitempty"`
		DisplayName          *string                      `json:"displayName,omitempty"`
		ExternalId           nullable.Type[string]        `json:"externalId,omitempty"`
		ExternalName         nullable.Type[string]        `json:"externalName,omitempty"`
		ExternalSource       *EducationExternalSource     `json:"externalSource,omitempty"`
		ExternalSourceDetail nullable.Type[string]        `json:"externalSourceDetail,omitempty"`
		Grade                nullable.Type[string]        `json:"grade,omitempty"`
		Group                *Group                       `json:"group,omitempty"`
		MailNickname         *string                      `json:"mailNickname,omitempty"`
		Members              *[]EducationUser             `json:"members,omitempty"`
		Modules              *[]EducationModule           `json:"modules,omitempty"`
		Schools              *[]EducationSchool           `json:"schools,omitempty"`
		Teachers             *[]EducationUser             `json:"teachers,omitempty"`
		Term                 *EducationTerm               `json:"term,omitempty"`
		Id                   *string                      `json:"id,omitempty"`
		ODataId              *string                      `json:"@odata.id,omitempty"`
		ODataType            *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignmentCategories = decoded.AssignmentCategories
	s.AssignmentDefaults = decoded.AssignmentDefaults
	s.AssignmentSettings = decoded.AssignmentSettings
	s.Assignments = decoded.Assignments
	s.ClassCode = decoded.ClassCode
	s.Course = decoded.Course
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.ExternalName = decoded.ExternalName
	s.ExternalSource = decoded.ExternalSource
	s.ExternalSourceDetail = decoded.ExternalSourceDetail
	s.Grade = decoded.Grade
	s.Group = decoded.Group
	s.MailNickname = decoded.MailNickname
	s.Members = decoded.Members
	s.Modules = decoded.Modules
	s.Schools = decoded.Schools
	s.Teachers = decoded.Teachers
	s.Term = decoded.Term
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationClass into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationClass': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
