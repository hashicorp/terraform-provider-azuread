package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationAssignment{}

type EducationAssignment struct {
	// Optional field to control the assignment behavior for adding assignments to students' and teachers' calendars when
	// the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners,
	// unknownFutureValue, and studentsOnly. Use the Prefer: include-unknown-enum-members request header to get the
	// following value(s) in this evolvable enum: studentsOnly. The default value is none.
	AddToCalendarAction *EducationAddToCalendarOptions `json:"addToCalendarAction,omitempty"`

	// Optional field to control the assignment behavior for students who are added after the assignment is published. If
	// not specified, defaults to none. Supported values are: none, assignIfOpen. For example, a teacher can use
	// assignIfOpen to indicate that an assignment should be assigned to any new student who joins the class while the
	// assignment is still open, and none to indicate that an assignment shouldn't be assigned to new students.
	AddedStudentAction *EducationAddedStudentAction `json:"addedStudentAction,omitempty"`

	// Identifies whether students can submit after the due date. If this property isn't specified during create, it
	// defaults to true.
	AllowLateSubmissions nullable.Type[bool] `json:"allowLateSubmissions,omitempty"`

	// Identifies whether students can add their own resources to a submission or if they can only modify resources added by
	// the teacher.
	AllowStudentsToAddResourcesToSubmission nullable.Type[bool] `json:"allowStudentsToAddResourcesToSubmission,omitempty"`

	// The date when the assignment should become active. If in the future, the assignment isn't shown to the student until
	// this date. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	AssignDateTime nullable.Type[string] `json:"assignDateTime,omitempty"`

	// Which users, or whole class should receive a submission object once the assignment is published.
	AssignTo EducationAssignmentRecipient `json:"assignTo"`

	// The moment that the assignment was published to students and the assignment shows up on the students timeline. The
	// Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	AssignedDateTime nullable.Type[string] `json:"assignedDateTime,omitempty"`

	// When set, enables users to easily find assignments of a given type. Read-only. Nullable.
	Categories *[]EducationCategory `json:"categories,omitempty"`

	// Class to which this assignment belongs.
	ClassId nullable.Type[string] `json:"classId,omitempty"`

	// Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment
	// doesn't allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the
	// closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z
	CloseDateTime nullable.Type[string] `json:"closeDateTime,omitempty"`

	// Who created the assignment.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Moment when the assignment was created. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Name of the assignment.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Date when the students assignment is due. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	DueDateTime nullable.Type[string] `json:"dueDateTime,omitempty"`

	// Folder URL where all the feedback file resources for this assignment are stored.
	FeedbackResourcesFolderUrl nullable.Type[string] `json:"feedbackResourcesFolderUrl,omitempty"`

	// How the assignment will be graded.
	Grading EducationAssignmentGradeType `json:"grading"`

	// When set, enables users to weight assignments differently when computing a class average grade.
	GradingCategory *EducationGradingCategory `json:"gradingCategory,omitempty"`

	GradingScheme *EducationGradingScheme `json:"gradingScheme,omitempty"`

	// Instructions for the assignment. This property and the display name tell the student what to do.
	Instructions *EducationItemBody `json:"instructions,omitempty"`

	// Specifies the language in which UI notifications for the assignment are displayed. If languageTag isn't provided, the
	// default language is en-US. Optional.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// Who last modified the assignment.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The date and time on which the assignment was modified. A student submission doesn't modify the assignment; only
	// teachers can update assignments. The Timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The URL of the module from which to access the assignment.
	ModuleUrl nullable.Type[string] `json:"moduleUrl,omitempty"`

	// Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or
	// null, defaults to the General channel. This field only applies to assignments where the assignTo value is
	// educationAssignmentClassRecipient. Updating the notificationChannelUrl isn't allowed after the assignment has been
	// published.
	NotificationChannelUrl nullable.Type[string] `json:"notificationChannelUrl,omitempty"`

	// Learning objects that are associated with this assignment. Only teachers can modify this list. Nullable.
	Resources *[]EducationAssignmentResource `json:"resources,omitempty"`

	// Folder URL where all the file resources for this assignment are stored.
	ResourcesFolderUrl nullable.Type[string] `json:"resourcesFolderUrl,omitempty"`

	// When set, the grading rubric attached to this assignment.
	Rubric *EducationRubric `json:"rubric,omitempty"`

	// Status of the assignment. You can't PATCH this value. Possible values are: draft, scheduled, published, assigned,
	// unknownFutureValue, inactive. Use the Prefer: include-unknown-enum-members request header to get the following
	// value(s) in this evolvable enum: inactive.
	Status *EducationAssignmentStatus `json:"status,omitempty"`

	// Once published, there is a submission object for each student representing their work and grade. Read-only. Nullable.
	Submissions *[]EducationSubmission `json:"submissions,omitempty"`

	// The deep link URL for the given assignment.
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

func (s EducationAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignment{}

func (s EducationAssignment) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignment: %+v", err)
	}

	delete(decoded, "assignDateTime")
	delete(decoded, "assignedDateTime")
	delete(decoded, "categories")
	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "feedbackResourcesFolderUrl")
	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "resourcesFolderUrl")
	delete(decoded, "status")
	delete(decoded, "submissions")
	delete(decoded, "webUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationAssignment{}

func (s *EducationAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AddToCalendarAction                     *EducationAddToCalendarOptions `json:"addToCalendarAction,omitempty"`
		AddedStudentAction                      *EducationAddedStudentAction   `json:"addedStudentAction,omitempty"`
		AllowLateSubmissions                    nullable.Type[bool]            `json:"allowLateSubmissions,omitempty"`
		AllowStudentsToAddResourcesToSubmission nullable.Type[bool]            `json:"allowStudentsToAddResourcesToSubmission,omitempty"`
		AssignDateTime                          nullable.Type[string]          `json:"assignDateTime,omitempty"`
		AssignedDateTime                        nullable.Type[string]          `json:"assignedDateTime,omitempty"`
		Categories                              *[]EducationCategory           `json:"categories,omitempty"`
		ClassId                                 nullable.Type[string]          `json:"classId,omitempty"`
		CloseDateTime                           nullable.Type[string]          `json:"closeDateTime,omitempty"`
		CreatedDateTime                         nullable.Type[string]          `json:"createdDateTime,omitempty"`
		DisplayName                             nullable.Type[string]          `json:"displayName,omitempty"`
		DueDateTime                             nullable.Type[string]          `json:"dueDateTime,omitempty"`
		FeedbackResourcesFolderUrl              nullable.Type[string]          `json:"feedbackResourcesFolderUrl,omitempty"`
		GradingCategory                         *EducationGradingCategory      `json:"gradingCategory,omitempty"`
		GradingScheme                           *EducationGradingScheme        `json:"gradingScheme,omitempty"`
		Instructions                            *EducationItemBody             `json:"instructions,omitempty"`
		LanguageTag                             nullable.Type[string]          `json:"languageTag,omitempty"`
		LastModifiedDateTime                    nullable.Type[string]          `json:"lastModifiedDateTime,omitempty"`
		ModuleUrl                               nullable.Type[string]          `json:"moduleUrl,omitempty"`
		NotificationChannelUrl                  nullable.Type[string]          `json:"notificationChannelUrl,omitempty"`
		Resources                               *[]EducationAssignmentResource `json:"resources,omitempty"`
		ResourcesFolderUrl                      nullable.Type[string]          `json:"resourcesFolderUrl,omitempty"`
		Rubric                                  *EducationRubric               `json:"rubric,omitempty"`
		Status                                  *EducationAssignmentStatus     `json:"status,omitempty"`
		Submissions                             *[]EducationSubmission         `json:"submissions,omitempty"`
		WebUrl                                  nullable.Type[string]          `json:"webUrl,omitempty"`
		Id                                      *string                        `json:"id,omitempty"`
		ODataId                                 *string                        `json:"@odata.id,omitempty"`
		ODataType                               *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AddToCalendarAction = decoded.AddToCalendarAction
	s.AddedStudentAction = decoded.AddedStudentAction
	s.AllowLateSubmissions = decoded.AllowLateSubmissions
	s.AllowStudentsToAddResourcesToSubmission = decoded.AllowStudentsToAddResourcesToSubmission
	s.AssignDateTime = decoded.AssignDateTime
	s.AssignedDateTime = decoded.AssignedDateTime
	s.Categories = decoded.Categories
	s.ClassId = decoded.ClassId
	s.CloseDateTime = decoded.CloseDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.DueDateTime = decoded.DueDateTime
	s.FeedbackResourcesFolderUrl = decoded.FeedbackResourcesFolderUrl
	s.GradingCategory = decoded.GradingCategory
	s.GradingScheme = decoded.GradingScheme
	s.Instructions = decoded.Instructions
	s.LanguageTag = decoded.LanguageTag
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ModuleUrl = decoded.ModuleUrl
	s.NotificationChannelUrl = decoded.NotificationChannelUrl
	s.Resources = decoded.Resources
	s.ResourcesFolderUrl = decoded.ResourcesFolderUrl
	s.Rubric = decoded.Rubric
	s.Status = decoded.Status
	s.Submissions = decoded.Submissions
	s.WebUrl = decoded.WebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["assignTo"]; ok {
		impl, err := UnmarshalEducationAssignmentRecipientImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssignTo' for 'EducationAssignment': %+v", err)
		}
		s.AssignTo = impl
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationAssignment': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["grading"]; ok {
		impl, err := UnmarshalEducationAssignmentGradeTypeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Grading' for 'EducationAssignment': %+v", err)
		}
		s.Grading = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationAssignment': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
