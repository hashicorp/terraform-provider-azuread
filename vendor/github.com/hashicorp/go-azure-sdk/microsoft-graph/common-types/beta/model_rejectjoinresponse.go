package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ParticipantJoiningResponse = RejectJoinResponse{}

type RejectJoinResponse struct {
	Reason *RejectReason `json:"reason,omitempty"`

	// Fields inherited from ParticipantJoiningResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s RejectJoinResponse) ParticipantJoiningResponse() BaseParticipantJoiningResponseImpl {
	return BaseParticipantJoiningResponseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RejectJoinResponse{}

func (s RejectJoinResponse) MarshalJSON() ([]byte, error) {
	type wrapper RejectJoinResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RejectJoinResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RejectJoinResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.rejectJoinResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RejectJoinResponse: %+v", err)
	}

	return encoded, nil
}
