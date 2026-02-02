package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ReadingAssignmentSubmission{}

type ReadingAssignmentSubmission struct {
	// Indicates whether the submission is an attempt by the student or a miscue edit done by the educator. The possible
	// values are Attempt and EditMiscue.
	Action *string `json:"action,omitempty"`

	// ID of the assignment with which this submission is associated.
	AssignmentId *string `json:"assignmentId,omitempty"`

	// List of words that the student found challenging during the reading session.
	ChallengingWords *[]ChallengingWord `json:"challengingWords,omitempty"`

	// ID of the class this reading progress is associated with.
	ClassId nullable.Type[string] `json:"classId,omitempty"`

	// Insertions of the reading progress.
	Insertions *int64 `json:"insertions,omitempty"`

	// Mispronunciations of the reading progress.
	Mispronunciations *int64 `json:"mispronunciations,omitempty"`

	// Number of exclamation marks missed in the reading passage.
	MissedExclamationMarks *int64 `json:"missedExclamationMarks,omitempty"`

	// Number of periods missed in the reading passage.
	MissedPeriods *int64 `json:"missedPeriods,omitempty"`

	// Number of question marks missed in the reading passage.
	MissedQuestionMarks *int64 `json:"missedQuestionMarks,omitempty"`

	// Number of short words missed during the reading session.
	MissedShorts *int64 `json:"missedShorts,omitempty"`

	// Omissions of the reading progress.
	Omissions *int64 `json:"omissions,omitempty"`

	// Number of times the student repeated words or phrases during the reading session.
	Repetitions *int64 `json:"repetitions,omitempty"`

	// Number of times the student self-corrected their reading errors.
	SelfCorrections *int64 `json:"selfCorrections,omitempty"`

	// ID of the user this reading progress is associated with.
	StudentId nullable.Type[string] `json:"studentId,omitempty"`

	// Date and time of the submission this reading progress is associated with. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	SubmissionDateTime *string `json:"submissionDateTime,omitempty"`

	// ID of the submission this reading progress is associated with.
	SubmissionId nullable.Type[string] `json:"submissionId,omitempty"`

	// Number of unexpected pauses made during the reading session.
	UnexpectedPauses *int64 `json:"unexpectedPauses,omitempty"`

	// Words count of the reading progress.
	WordCount *int64 `json:"wordCount,omitempty"`

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

func (s ReadingAssignmentSubmission) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ReadingAssignmentSubmission{}

func (s ReadingAssignmentSubmission) MarshalJSON() ([]byte, error) {
	type wrapper ReadingAssignmentSubmission
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ReadingAssignmentSubmission: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ReadingAssignmentSubmission: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.readingAssignmentSubmission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ReadingAssignmentSubmission: %+v", err)
	}

	return encoded, nil
}
