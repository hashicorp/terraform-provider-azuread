package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Channel{}

type Channel struct {
	// A collection of membership records associated with the channel, including both direct and indirect members of shared
	// channels.
	AllMembers *[]ConversationMember `json:"allMembers,omitempty"`

	// Read only. Timestamp at which the channel was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Optional textual description for the channel.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Channel name as it will appear to the user in Microsoft Teams. The maximum length is 50 characters.
	DisplayName *string `json:"displayName,omitempty"`

	// The email address for sending messages to the channel. Read-only.
	Email nullable.Type[string] `json:"email,omitempty"`

	// Metadata for the location where the channel's files are stored.
	FilesFolder *DriveItem `json:"filesFolder,omitempty"`

	// Indicates whether the channel is archived. Read-only.
	IsArchived nullable.Type[bool] `json:"isArchived,omitempty"`

	// Indicates whether the channel should be marked as recommended for all members of the team to show in their channel
	// list. Note: All recommended channels automatically show in the channels list for education and frontline worker
	// users. The property can only be set programmatically via the Create team method. The default value is false.
	IsFavoriteByDefault nullable.Type[bool] `json:"isFavoriteByDefault,omitempty"`

	// A collection of membership records associated with the channel.
	Members *[]ConversationMember `json:"members,omitempty"`

	// The type of the channel. Can be set during creation and can't be changed. The possible values are: standard, private,
	// unknownFutureValue, shared. The default value is standard. Use the Prefer: include-unknown-enum-members request
	// header to get the following value in this evolvable enum: shared.
	MembershipType *ChannelMembershipType `json:"membershipType,omitempty"`

	// A collection of all the messages in the channel. A navigation property. Nullable.
	Messages *[]ChatMessage `json:"messages,omitempty"`

	// A collection of teams with which a channel is shared.
	SharedWithTeams *[]SharedWithChannelTeamInfo `json:"sharedWithTeams,omitempty"`

	// Contains summary information about the channel, including number of owners, members, guests, and an indicator for
	// members from other tenants. The summary property will only be returned if it is specified in the $select clause of
	// the Get channel method.
	Summary *ChannelSummary `json:"summary,omitempty"`

	// A collection of all the tabs in the channel. A navigation property.
	Tabs *[]TeamsTab `json:"tabs,omitempty"`

	// The ID of the Microsoft Entra tenant.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a
	// channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not
	// parsed. Read-only.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

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

func (s Channel) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Channel{}

func (s Channel) MarshalJSON() ([]byte, error) {
	type wrapper Channel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Channel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Channel: %+v", err)
	}

	delete(decoded, "email")
	delete(decoded, "isArchived")
	delete(decoded, "webUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.channel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Channel: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Channel{}

func (s *Channel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime     nullable.Type[string]        `json:"createdDateTime,omitempty"`
		Description         nullable.Type[string]        `json:"description,omitempty"`
		DisplayName         *string                      `json:"displayName,omitempty"`
		Email               nullable.Type[string]        `json:"email,omitempty"`
		FilesFolder         *DriveItem                   `json:"filesFolder,omitempty"`
		IsArchived          nullable.Type[bool]          `json:"isArchived,omitempty"`
		IsFavoriteByDefault nullable.Type[bool]          `json:"isFavoriteByDefault,omitempty"`
		MembershipType      *ChannelMembershipType       `json:"membershipType,omitempty"`
		Messages            *[]ChatMessage               `json:"messages,omitempty"`
		SharedWithTeams     *[]SharedWithChannelTeamInfo `json:"sharedWithTeams,omitempty"`
		Summary             *ChannelSummary              `json:"summary,omitempty"`
		Tabs                *[]TeamsTab                  `json:"tabs,omitempty"`
		TenantId            nullable.Type[string]        `json:"tenantId,omitempty"`
		WebUrl              nullable.Type[string]        `json:"webUrl,omitempty"`
		Id                  *string                      `json:"id,omitempty"`
		ODataId             *string                      `json:"@odata.id,omitempty"`
		ODataType           *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Email = decoded.Email
	s.FilesFolder = decoded.FilesFolder
	s.IsArchived = decoded.IsArchived
	s.IsFavoriteByDefault = decoded.IsFavoriteByDefault
	s.MembershipType = decoded.MembershipType
	s.Messages = decoded.Messages
	s.SharedWithTeams = decoded.SharedWithTeams
	s.Summary = decoded.Summary
	s.Tabs = decoded.Tabs
	s.TenantId = decoded.TenantId
	s.WebUrl = decoded.WebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Channel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allMembers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AllMembers into list []json.RawMessage: %+v", err)
		}

		output := make([]ConversationMember, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalConversationMemberImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AllMembers' for 'Channel': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AllMembers = &output
	}

	if v, ok := temp["members"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Members into list []json.RawMessage: %+v", err)
		}

		output := make([]ConversationMember, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalConversationMemberImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'Channel': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}

	return nil
}
