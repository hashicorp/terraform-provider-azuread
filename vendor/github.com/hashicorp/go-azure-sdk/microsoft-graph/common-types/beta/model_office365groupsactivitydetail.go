package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365GroupsActivityDetail{}

type Office365GroupsActivityDetail struct {
	// The storage used of the group mailbox.
	ExchangeMailboxStorageUsedInBytes nullable.Type[int64] `json:"exchangeMailboxStorageUsedInBytes,omitempty"`

	// The number of items in the group mailbox.
	ExchangeMailboxTotalItemCount nullable.Type[int64] `json:"exchangeMailboxTotalItemCount,omitempty"`

	// The number of emails that the group mailbox received.
	ExchangeReceivedEmailCount nullable.Type[int64] `json:"exchangeReceivedEmailCount,omitempty"`

	// The group external member count.
	ExternalMemberCount nullable.Type[int64] `json:"externalMemberCount,omitempty"`

	// The display name of the group.
	GroupDisplayName nullable.Type[string] `json:"groupDisplayName,omitempty"`

	// The group id.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// The group type. Possible values are: Public or Private.
	GroupType nullable.Type[string] `json:"groupType,omitempty"`

	// Whether this user has been deleted or soft deleted.
	IsDeleted nullable.Type[bool] `json:"isDeleted,omitempty"`

	// The last activity date for the following scenarios: group mailbox received email; user viewed, edited, shared, or
	// synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in
	// Yammer groups.
	LastActivityDate nullable.Type[string] `json:"lastActivityDate,omitempty"`

	// The group member count.
	MemberCount nullable.Type[int64] `json:"memberCount,omitempty"`

	// The group owner principal name.
	OwnerPrincipalName nullable.Type[string] `json:"ownerPrincipalName,omitempty"`

	// The number of days the report covers.
	ReportPeriod nullable.Type[string] `json:"reportPeriod,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The number of active files in SharePoint Group site.
	SharePointActiveFileCount nullable.Type[int64] `json:"sharePointActiveFileCount,omitempty"`

	// The storage used by SharePoint Group site.
	SharePointSiteStorageUsedInBytes nullable.Type[int64] `json:"sharePointSiteStorageUsedInBytes,omitempty"`

	// The total number of files in SharePoint Group site.
	SharePointTotalFileCount nullable.Type[int64] `json:"sharePointTotalFileCount,omitempty"`

	// The number of channel messages in Teams team.
	TeamsChannelMessagesCount nullable.Type[int64] `json:"teamsChannelMessagesCount,omitempty"`

	// The number of meetings organized in Teams team.
	TeamsMeetingsOrganizedCount nullable.Type[int64] `json:"teamsMeetingsOrganizedCount,omitempty"`

	// The number of messages liked in Yammer groups.
	YammerLikedMessageCount nullable.Type[int64] `json:"yammerLikedMessageCount,omitempty"`

	// The number of messages posted to Yammer groups.
	YammerPostedMessageCount nullable.Type[int64] `json:"yammerPostedMessageCount,omitempty"`

	// The number of messages read in Yammer groups.
	YammerReadMessageCount nullable.Type[int64] `json:"yammerReadMessageCount,omitempty"`

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

func (s Office365GroupsActivityDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365GroupsActivityDetail{}

func (s Office365GroupsActivityDetail) MarshalJSON() ([]byte, error) {
	type wrapper Office365GroupsActivityDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365GroupsActivityDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365GroupsActivityDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365GroupsActivityDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365GroupsActivityDetail: %+v", err)
	}

	return encoded, nil
}
