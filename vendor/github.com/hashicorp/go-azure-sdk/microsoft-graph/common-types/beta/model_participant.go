package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Participant{}

type Participant struct {
	Info                 *ParticipantInfo    `json:"info,omitempty"`
	IsIdentityAnonymized nullable.Type[bool] `json:"isIdentityAnonymized,omitempty"`

	// true if the participant is in lobby.
	IsInLobby *bool `json:"isInLobby,omitempty"`

	// true if the participant is muted (client or server muted).
	IsMuted *bool `json:"isMuted,omitempty"`

	// The list of media streams.
	MediaStreams *[]MediaStream `json:"mediaStreams,omitempty"`

	// A blob of data provided by the participant in the roster.
	Metadata nullable.Type[string] `json:"metadata,omitempty"`

	// The participant's preferred display name that overrides the original display name.
	PreferredDisplayName nullable.Type[string] `json:"preferredDisplayName,omitempty"`

	// Information on whether the participant has recording capability.
	RecordingInfo *RecordingInfo `json:"recordingInfo,omitempty"`

	// Indicates the reason why the participant was removed from the roster.
	RemovedState *RemovedState `json:"removedState,omitempty"`

	// Indicates the reason or reasons why media content from this participant is restricted.
	RestrictedExperience *OnlineMeetingRestricted `json:"restrictedExperience,omitempty"`

	// Indicates the roster sequence number the participant was last updated in.
	RosterSequenceNumber nullable.Type[int64] `json:"rosterSequenceNumber,omitempty"`

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

func (s Participant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Participant{}

func (s Participant) MarshalJSON() ([]byte, error) {
	type wrapper Participant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Participant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Participant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.participant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Participant: %+v", err)
	}

	return encoded, nil
}
