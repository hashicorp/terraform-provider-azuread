package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CallAiInsight{}

type CallAiInsight struct {
	// The collection of AI-generated action items. Read-only.
	ActionItems *[]ActionItem `json:"actionItems,omitempty"`

	// The ID for the online meeting call for which the callAiInsight was generated. Read-only.
	CallId nullable.Type[string] `json:"callId,omitempty"`

	// The unique ID that correlates the transcript from which the insights were generated. Read-only.
	ContentCorrelationId nullable.Type[string] `json:"contentCorrelationId,omitempty"`

	// Date and time at which the corresponding transcript was created. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Date and time at which the corresponding transcription ends. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Read-only.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The collection of AI-generated meeting notes. Read-only.
	MeetingNotes *[]MeetingNote `json:"meetingNotes,omitempty"`

	// The caller-specific properties of the callAiInsight entity. Read-only.
	Viewpoint *CallAiInsightViewPoint `json:"viewpoint,omitempty"`

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

func (s CallAiInsight) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallAiInsight{}

func (s CallAiInsight) MarshalJSON() ([]byte, error) {
	type wrapper CallAiInsight
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallAiInsight: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallAiInsight: %+v", err)
	}

	delete(decoded, "actionItems")
	delete(decoded, "callId")
	delete(decoded, "contentCorrelationId")
	delete(decoded, "createdDateTime")
	delete(decoded, "endDateTime")
	delete(decoded, "meetingNotes")
	delete(decoded, "viewpoint")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callAiInsight"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallAiInsight: %+v", err)
	}

	return encoded, nil
}
