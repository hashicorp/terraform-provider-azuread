package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFeedback struct {
	// User who created the feedback.
	FeedbackBy IdentitySet `json:"feedbackBy"`

	// Moment in time when the feedback was given. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	FeedbackDateTime nullable.Type[string] `json:"feedbackDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Feedback.
	Text *EducationItemBody `json:"text,omitempty"`
}

var _ json.Unmarshaler = &EducationFeedback{}

func (s *EducationFeedback) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FeedbackDateTime nullable.Type[string] `json:"feedbackDateTime,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
		Text             *EducationItemBody    `json:"text,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FeedbackDateTime = decoded.FeedbackDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Text = decoded.Text

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationFeedback into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["feedbackBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'FeedbackBy' for 'EducationFeedback': %+v", err)
		}
		s.FeedbackBy = impl
	}

	return nil
}
