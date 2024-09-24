package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = TeamCreatedEventMessageDetail{}

type TeamCreatedEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Description for the team.
	TeamDescription nullable.Type[string] `json:"teamDescription,omitempty"`

	// Display name of the team.
	TeamDisplayName nullable.Type[string] `json:"teamDisplayName,omitempty"`

	// Unique identifier of the team.
	TeamId nullable.Type[string] `json:"teamId,omitempty"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TeamCreatedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamCreatedEventMessageDetail{}

func (s TeamCreatedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper TeamCreatedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamCreatedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamCreatedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamCreatedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamCreatedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamCreatedEventMessageDetail{}

func (s *TeamCreatedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		TeamDescription nullable.Type[string] `json:"teamDescription,omitempty"`
		TeamDisplayName nullable.Type[string] `json:"teamDisplayName,omitempty"`
		TeamId          nullable.Type[string] `json:"teamId,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.TeamDescription = decoded.TeamDescription
	s.TeamDisplayName = decoded.TeamDisplayName
	s.TeamId = decoded.TeamId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamCreatedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'TeamCreatedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
