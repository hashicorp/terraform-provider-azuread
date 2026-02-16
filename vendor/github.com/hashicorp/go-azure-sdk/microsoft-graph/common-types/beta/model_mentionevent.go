package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MentionEvent struct {
	// The date and time of the mention event. The timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EventDateTime nullable.Type[string] `json:"eventDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The speaker who mentioned the user.
	Speaker IdentitySet `json:"speaker"`

	// The utterance in the online meeting transcript that contains the mention event.
	TranscriptUtterance nullable.Type[string] `json:"transcriptUtterance,omitempty"`
}

var _ json.Unmarshaler = &MentionEvent{}

func (s *MentionEvent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EventDateTime       nullable.Type[string] `json:"eventDateTime,omitempty"`
		ODataId             *string               `json:"@odata.id,omitempty"`
		ODataType           *string               `json:"@odata.type,omitempty"`
		TranscriptUtterance nullable.Type[string] `json:"transcriptUtterance,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EventDateTime = decoded.EventDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.TranscriptUtterance = decoded.TranscriptUtterance

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MentionEvent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["speaker"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Speaker' for 'MentionEvent': %+v", err)
		}
		s.Speaker = impl
	}

	return nil
}
