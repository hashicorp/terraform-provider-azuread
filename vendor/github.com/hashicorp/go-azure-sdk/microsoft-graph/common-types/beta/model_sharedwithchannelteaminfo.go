package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamInfo = SharedWithChannelTeamInfo{}

type SharedWithChannelTeamInfo struct {
	// A collection of team members who have access to the shared channel.
	AllowedMembers *[]ConversationMember `json:"allowedMembers,omitempty"`

	// Indicates whether the team is the host of the channel.
	IsHostTeam nullable.Type[bool] `json:"isHostTeam,omitempty"`

	// Fields inherited from TeamInfo

	// The name of the team.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	Team *Team `json:"team,omitempty"`

	// The ID of the Microsoft Entra tenant.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

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

func (s SharedWithChannelTeamInfo) TeamInfo() BaseTeamInfoImpl {
	return BaseTeamInfoImpl{
		DisplayName: s.DisplayName,
		Team:        s.Team,
		TenantId:    s.TenantId,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s SharedWithChannelTeamInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SharedWithChannelTeamInfo{}

func (s SharedWithChannelTeamInfo) MarshalJSON() ([]byte, error) {
	type wrapper SharedWithChannelTeamInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharedWithChannelTeamInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharedWithChannelTeamInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sharedWithChannelTeamInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharedWithChannelTeamInfo: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SharedWithChannelTeamInfo{}

func (s *SharedWithChannelTeamInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsHostTeam  nullable.Type[bool]   `json:"isHostTeam,omitempty"`
		DisplayName nullable.Type[string] `json:"displayName,omitempty"`
		Team        *Team                 `json:"team,omitempty"`
		TenantId    nullable.Type[string] `json:"tenantId,omitempty"`
		Id          *string               `json:"id,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsHostTeam = decoded.IsHostTeam
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Team = decoded.Team
	s.TenantId = decoded.TenantId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharedWithChannelTeamInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allowedMembers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AllowedMembers into list []json.RawMessage: %+v", err)
		}

		output := make([]ConversationMember, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalConversationMemberImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AllowedMembers' for 'SharedWithChannelTeamInfo': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AllowedMembers = &output
	}

	return nil
}
