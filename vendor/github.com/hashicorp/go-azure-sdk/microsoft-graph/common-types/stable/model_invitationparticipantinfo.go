package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationParticipantInfo struct {
	// Optional. Whether to hide the participant from the roster.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

	Identity IdentitySet `json:"identity"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional. The ID of the target participant.
	ParticipantId nullable.Type[string] `json:"participantId,omitempty"`

	// Optional. Whether to remove them from the main mixer.
	RemoveFromDefaultAudioRoutingGroup nullable.Type[bool] `json:"removeFromDefaultAudioRoutingGroup,omitempty"`

	// Optional. The call which the target identity is currently a part of. For peer-to-peer case, the call will be dropped
	// once the participant is added successfully.
	ReplacesCallId nullable.Type[string] `json:"replacesCallId,omitempty"`
}

var _ json.Unmarshaler = &InvitationParticipantInfo{}

func (s *InvitationParticipantInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Hidden                             nullable.Type[bool]   `json:"hidden,omitempty"`
		ODataId                            *string               `json:"@odata.id,omitempty"`
		ODataType                          *string               `json:"@odata.type,omitempty"`
		ParticipantId                      nullable.Type[string] `json:"participantId,omitempty"`
		RemoveFromDefaultAudioRoutingGroup nullable.Type[bool]   `json:"removeFromDefaultAudioRoutingGroup,omitempty"`
		ReplacesCallId                     nullable.Type[string] `json:"replacesCallId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Hidden = decoded.Hidden
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ParticipantId = decoded.ParticipantId
	s.RemoveFromDefaultAudioRoutingGroup = decoded.RemoveFromDefaultAudioRoutingGroup
	s.ReplacesCallId = decoded.ReplacesCallId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling InvitationParticipantInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'InvitationParticipantInfo': %+v", err)
		}
		s.Identity = impl
	}

	return nil
}
