package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ReflectCheckInResponse{}

type ReflectCheckInResponse struct {
	// Identifier for the Reflect check-in.
	CheckInId nullable.Type[string] `json:"checkInId,omitempty"`

	// The question or prompt of the Reflect check-in that this response addresses.
	CheckInTitle *string `json:"checkInTitle,omitempty"`

	// ID of the class associated with the Reflect check-in.
	ClassId nullable.Type[string] `json:"classId,omitempty"`

	// Date and time when the Reflect check-in was created. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// ID of the user who created the Reflect check-in.
	CreatorId nullable.Type[string] `json:"creatorId,omitempty"`

	// Indicates whether the Reflect check-in is closed (true) or open (false).
	IsClosed *bool `json:"isClosed,omitempty"`

	// ID of the user who responded to the Reflect check-in.
	ResponderId nullable.Type[string] `json:"responderId,omitempty"`

	ResponseEmotion  *ResponseEmotionType  `json:"responseEmotion,omitempty"`
	ResponseFeedback *ResponseFeedbackType `json:"responseFeedback,omitempty"`

	// Date and time when the response to the Reflect check-in was submitted. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	SubmitDateTime *string `json:"submitDateTime,omitempty"`

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

func (s ReflectCheckInResponse) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ReflectCheckInResponse{}

func (s ReflectCheckInResponse) MarshalJSON() ([]byte, error) {
	type wrapper ReflectCheckInResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ReflectCheckInResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ReflectCheckInResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.reflectCheckInResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ReflectCheckInResponse: %+v", err)
	}

	return encoded, nil
}
