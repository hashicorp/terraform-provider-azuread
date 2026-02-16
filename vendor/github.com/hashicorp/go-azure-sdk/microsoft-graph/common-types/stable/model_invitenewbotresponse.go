package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ParticipantJoiningResponse = InviteNewBotResponse{}

type InviteNewBotResponse struct {
	// URI to receive new incoming call notification.
	InviteUri nullable.Type[string] `json:"inviteUri,omitempty"`

	// Fields inherited from ParticipantJoiningResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s InviteNewBotResponse) ParticipantJoiningResponse() BaseParticipantJoiningResponseImpl {
	return BaseParticipantJoiningResponseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InviteNewBotResponse{}

func (s InviteNewBotResponse) MarshalJSON() ([]byte, error) {
	type wrapper InviteNewBotResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InviteNewBotResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InviteNewBotResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.inviteNewBotResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InviteNewBotResponse: %+v", err)
	}

	return encoded, nil
}
