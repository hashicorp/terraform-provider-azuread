package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Team{}

type Team struct {
	// List of channels either hosted in or shared with the team (incoming channels).
	AllChannels *[]Channel `json:"allChannels,omitempty"`

	// The collection of channels and messages associated with the team.
	Channels *[]Channel `json:"channels,omitempty"`

	// An optional label. Typically describes the data or business sensitivity of the team. Must match one of a
	// pre-configured set in the tenant's directory.
	Classification nullable.Type[string] `json:"classification,omitempty"`

	// Timestamp at which the team was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// An optional description for the team. Maximum length: 1,024 characters.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Settings to configure team discoverability by others.
	DiscoverySettings *TeamDiscoverySettings `json:"discoverySettings,omitempty"`

	// The name of the team.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The name of the first channel in the team. This is an optional property, only used during team creation and isn't
	// returned in methods to get and list teams.
	FirstChannelName nullable.Type[string] `json:"firstChannelName,omitempty"`

	// Settings to configure the use of Giphy, memes, and stickers in the team.
	FunSettings *TeamFunSettings `json:"funSettings,omitempty"`

	Group *Group `json:"group,omitempty"`

	// Settings to configure whether guests can create, update, or delete channels in the team.
	GuestSettings *TeamGuestSettings `json:"guestSettings,omitempty"`

	// List of channels shared with the team.
	IncomingChannels *[]Channel `json:"incomingChannels,omitempty"`

	// The apps installed in this team.
	InstalledApps *[]TeamsAppInstallation `json:"installedApps,omitempty"`

	// A unique ID for the team used in a few places such as the audit log/Office 365 Management Activity API.
	InternalId nullable.Type[string] `json:"internalId,omitempty"`

	// Whether this team is in read-only mode.
	IsArchived nullable.Type[bool] `json:"isArchived,omitempty"`

	// If set to true, the team is currently in the owner-only team membership state and inaccessible by other team members,
	// such as students.
	IsMembershipLimitedToOwners nullable.Type[bool] `json:"isMembershipLimitedToOwners,omitempty"`

	// Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the
	// team.
	MemberSettings *TeamMemberSettings `json:"memberSettings,omitempty"`

	// Members and owners of the team.
	Members *[]ConversationMember `json:"members,omitempty"`

	// Settings to configure messaging and mentions in the team.
	MessagingSettings *TeamMessagingSettings `json:"messagingSettings,omitempty"`

	// The async operations that ran or are running on this team.
	Operations *[]TeamsAsyncOperation `json:"operations,omitempty"`

	// The list of this team's owners. Currently, when creating a team using application permissions, exactly one owner must
	// be specified. When using user-delegated permissions, no owner can be specified (the current user is the owner). The
	// owner must be specified as an object ID (GUID), not a UPN.
	Owners *[]User `json:"owners,omitempty"`

	// A collection of permissions granted to apps to access the team.
	PermissionGrants *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`

	// The team photo.
	Photo *ProfilePhoto `json:"photo,omitempty"`

	// The general channel for the team.
	PrimaryChannel *Channel `json:"primaryChannel,omitempty"`

	// The schedule of shifts for this team.
	Schedule *Schedule `json:"schedule,omitempty"`

	// Optional. Indicates whether the team is intended for a particular use case. Each team specialization has access to
	// unique behaviors and experiences targeted to its use case.
	Specialization *TeamSpecialization `json:"specialization,omitempty"`

	// Contains summary information about the team, including the number of owners, members, and guests.
	Summary *TeamSummary `json:"summary,omitempty"`

	// The tags associated with the team.
	Tags *[]TeamworkTag `json:"tags,omitempty"`

	// The template this team was created from. See available templates.
	Template *TeamsTemplate `json:"template,omitempty"`

	// Generic representation of a team template definition for a team with a specific structure and configuration.
	TemplateDefinition *TeamTemplateDefinition `json:"templateDefinition,omitempty"`

	// The ID of the Microsoft Entra tenant.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The visibility of the group and team. Defaults to Public.
	Visibility *TeamVisibilityType `json:"visibility,omitempty"`

	// A hyperlink that goes to the team in the Microsoft Teams client. It's the URL you get when you right-click a team in
	// the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
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

func (s Team) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Team{}

func (s Team) MarshalJSON() ([]byte, error) {
	type wrapper Team
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Team: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Team: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.team"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Team: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Team{}

func (s *Team) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllChannels                 *[]Channel                         `json:"allChannels,omitempty"`
		Channels                    *[]Channel                         `json:"channels,omitempty"`
		Classification              nullable.Type[string]              `json:"classification,omitempty"`
		CreatedDateTime             nullable.Type[string]              `json:"createdDateTime,omitempty"`
		Description                 nullable.Type[string]              `json:"description,omitempty"`
		DiscoverySettings           *TeamDiscoverySettings             `json:"discoverySettings,omitempty"`
		DisplayName                 nullable.Type[string]              `json:"displayName,omitempty"`
		FirstChannelName            nullable.Type[string]              `json:"firstChannelName,omitempty"`
		FunSettings                 *TeamFunSettings                   `json:"funSettings,omitempty"`
		Group                       *Group                             `json:"group,omitempty"`
		GuestSettings               *TeamGuestSettings                 `json:"guestSettings,omitempty"`
		IncomingChannels            *[]Channel                         `json:"incomingChannels,omitempty"`
		InternalId                  nullable.Type[string]              `json:"internalId,omitempty"`
		IsArchived                  nullable.Type[bool]                `json:"isArchived,omitempty"`
		IsMembershipLimitedToOwners nullable.Type[bool]                `json:"isMembershipLimitedToOwners,omitempty"`
		MemberSettings              *TeamMemberSettings                `json:"memberSettings,omitempty"`
		MessagingSettings           *TeamMessagingSettings             `json:"messagingSettings,omitempty"`
		Operations                  *[]TeamsAsyncOperation             `json:"operations,omitempty"`
		Owners                      *[]User                            `json:"owners,omitempty"`
		PermissionGrants            *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
		Photo                       *ProfilePhoto                      `json:"photo,omitempty"`
		PrimaryChannel              *Channel                           `json:"primaryChannel,omitempty"`
		Schedule                    *Schedule                          `json:"schedule,omitempty"`
		Specialization              *TeamSpecialization                `json:"specialization,omitempty"`
		Summary                     *TeamSummary                       `json:"summary,omitempty"`
		Tags                        *[]TeamworkTag                     `json:"tags,omitempty"`
		Template                    *TeamsTemplate                     `json:"template,omitempty"`
		TemplateDefinition          *TeamTemplateDefinition            `json:"templateDefinition,omitempty"`
		TenantId                    nullable.Type[string]              `json:"tenantId,omitempty"`
		Visibility                  *TeamVisibilityType                `json:"visibility,omitempty"`
		WebUrl                      nullable.Type[string]              `json:"webUrl,omitempty"`
		Id                          *string                            `json:"id,omitempty"`
		ODataId                     *string                            `json:"@odata.id,omitempty"`
		ODataType                   *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllChannels = decoded.AllChannels
	s.Channels = decoded.Channels
	s.Classification = decoded.Classification
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DiscoverySettings = decoded.DiscoverySettings
	s.DisplayName = decoded.DisplayName
	s.FirstChannelName = decoded.FirstChannelName
	s.FunSettings = decoded.FunSettings
	s.Group = decoded.Group
	s.GuestSettings = decoded.GuestSettings
	s.IncomingChannels = decoded.IncomingChannels
	s.InternalId = decoded.InternalId
	s.IsArchived = decoded.IsArchived
	s.IsMembershipLimitedToOwners = decoded.IsMembershipLimitedToOwners
	s.MemberSettings = decoded.MemberSettings
	s.MessagingSettings = decoded.MessagingSettings
	s.Operations = decoded.Operations
	s.Owners = decoded.Owners
	s.PermissionGrants = decoded.PermissionGrants
	s.Photo = decoded.Photo
	s.PrimaryChannel = decoded.PrimaryChannel
	s.Schedule = decoded.Schedule
	s.Specialization = decoded.Specialization
	s.Summary = decoded.Summary
	s.Tags = decoded.Tags
	s.Template = decoded.Template
	s.TemplateDefinition = decoded.TemplateDefinition
	s.TenantId = decoded.TenantId
	s.Visibility = decoded.Visibility
	s.WebUrl = decoded.WebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Team into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["installedApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling InstalledApps into list []json.RawMessage: %+v", err)
		}

		output := make([]TeamsAppInstallation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalTeamsAppInstallationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'InstalledApps' for 'Team': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.InstalledApps = &output
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
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'Team': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}

	return nil
}
