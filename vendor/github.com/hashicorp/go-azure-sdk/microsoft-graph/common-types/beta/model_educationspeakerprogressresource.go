package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationResource = EducationSpeakerProgressResource{}

type EducationSpeakerProgressResource struct {
	// The feedback types that students should receive from AI feedback. This property should only be provided if
	// isAiFeedbackEnabled is true.
	AiFeedbackCriteria *EducationAiFeedbackCriteria `json:"aiFeedbackCriteria,omitempty"`

	// Indicates whether AI feedback is enabled for the student submissions.
	IsAiFeedbackEnabled *bool `json:"isAiFeedbackEnabled,omitempty"`

	// Indicates whether video is required for the student recording.
	IsVideoRequired *bool `json:"isVideoRequired,omitempty"`

	// The maximum number of recording attempts available to the student. Specify 0 to set unlimited recording attempts.
	MaxRecordingAttempts *int64 `json:"maxRecordingAttempts,omitempty"`

	// The title of the speaker progress resource visible to students.
	PresentationTitle *string `json:"presentationTitle,omitempty"`

	// The time limit is in minutes for the student recording.
	RecordingTimeLimitInMinutes *int64 `json:"recordingTimeLimitInMinutes,omitempty"`

	// Allows students to view their rehearsal report before the assignment is graded.
	ShowRehearsalReportToStudentBeforeMediaUpload *bool `json:"showRehearsalReportToStudentBeforeMediaUpload,omitempty"`

	// The feedback types that students should receive from the Speaker Coach.
	SpeakerCoachSettings *EducationSpeakerCoachSettings `json:"speakerCoachSettings,omitempty"`

	// The spoken language for the student recording. For example, en-US.
	SpokenLanguageLocale *string `json:"spokenLanguageLocale,omitempty"`

	// Fields inherited from EducationResource

	// Who created the resource?
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Display name of resource.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Who was the last user to modify the resource?
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Moment in time when the resource was last modified. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationSpeakerProgressResource) EducationResource() BaseEducationResourceImpl {
	return BaseEducationResourceImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

var _ json.Marshaler = EducationSpeakerProgressResource{}

func (s EducationSpeakerProgressResource) MarshalJSON() ([]byte, error) {
	type wrapper EducationSpeakerProgressResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSpeakerProgressResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSpeakerProgressResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSpeakerProgressResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSpeakerProgressResource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationSpeakerProgressResource{}

func (s *EducationSpeakerProgressResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AiFeedbackCriteria                            *EducationAiFeedbackCriteria   `json:"aiFeedbackCriteria,omitempty"`
		IsAiFeedbackEnabled                           *bool                          `json:"isAiFeedbackEnabled,omitempty"`
		IsVideoRequired                               *bool                          `json:"isVideoRequired,omitempty"`
		MaxRecordingAttempts                          *int64                         `json:"maxRecordingAttempts,omitempty"`
		PresentationTitle                             *string                        `json:"presentationTitle,omitempty"`
		RecordingTimeLimitInMinutes                   *int64                         `json:"recordingTimeLimitInMinutes,omitempty"`
		ShowRehearsalReportToStudentBeforeMediaUpload *bool                          `json:"showRehearsalReportToStudentBeforeMediaUpload,omitempty"`
		SpeakerCoachSettings                          *EducationSpeakerCoachSettings `json:"speakerCoachSettings,omitempty"`
		SpokenLanguageLocale                          *string                        `json:"spokenLanguageLocale,omitempty"`
		CreatedDateTime                               nullable.Type[string]          `json:"createdDateTime,omitempty"`
		DisplayName                                   nullable.Type[string]          `json:"displayName,omitempty"`
		LastModifiedDateTime                          nullable.Type[string]          `json:"lastModifiedDateTime,omitempty"`
		ODataId                                       *string                        `json:"@odata.id,omitempty"`
		ODataType                                     *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AiFeedbackCriteria = decoded.AiFeedbackCriteria
	s.IsAiFeedbackEnabled = decoded.IsAiFeedbackEnabled
	s.IsVideoRequired = decoded.IsVideoRequired
	s.MaxRecordingAttempts = decoded.MaxRecordingAttempts
	s.PresentationTitle = decoded.PresentationTitle
	s.RecordingTimeLimitInMinutes = decoded.RecordingTimeLimitInMinutes
	s.ShowRehearsalReportToStudentBeforeMediaUpload = decoded.ShowRehearsalReportToStudentBeforeMediaUpload
	s.SpeakerCoachSettings = decoded.SpeakerCoachSettings
	s.SpokenLanguageLocale = decoded.SpokenLanguageLocale
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationSpeakerProgressResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationSpeakerProgressResource': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationSpeakerProgressResource': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
