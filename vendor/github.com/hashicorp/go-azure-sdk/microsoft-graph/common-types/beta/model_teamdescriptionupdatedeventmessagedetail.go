package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = TeamDescriptionUpdatedEventMessageDetail{}

type TeamDescriptionUpdatedEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// The updated description for the team.
	TeamDescription nullable.Type[string] `json:"teamDescription,omitempty"`

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

func (s TeamDescriptionUpdatedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamDescriptionUpdatedEventMessageDetail{}

func (s TeamDescriptionUpdatedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper TeamDescriptionUpdatedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamDescriptionUpdatedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamDescriptionUpdatedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamDescriptionUpdatedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamDescriptionUpdatedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamDescriptionUpdatedEventMessageDetail{}

func (s *TeamDescriptionUpdatedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		TeamDescription nullable.Type[string] `json:"teamDescription,omitempty"`
		TeamId          nullable.Type[string] `json:"teamId,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.TeamDescription = decoded.TeamDescription
	s.TeamId = decoded.TeamId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamDescriptionUpdatedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'TeamDescriptionUpdatedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
