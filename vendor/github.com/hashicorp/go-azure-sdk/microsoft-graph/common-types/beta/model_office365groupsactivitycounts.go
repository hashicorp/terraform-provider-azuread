package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365GroupsActivityCounts{}

type Office365GroupsActivityCounts struct {
	// The number of emails received by Group mailboxes.
	ExchangeEmailsReceived nullable.Type[int64] `json:"exchangeEmailsReceived,omitempty"`

	// The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked
	// in a Yammer group
	ReportDate nullable.Type[string] `json:"reportDate,omitempty"`

	// The number of days the report covers.
	ReportPeriod nullable.Type[string] `json:"reportPeriod,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The number of channel messages in Teams team.
	TeamsChannelMessages nullable.Type[int64] `json:"teamsChannelMessages,omitempty"`

	// The number of meetings organized in Teams team.
	TeamsMeetingsOrganized nullable.Type[int64] `json:"teamsMeetingsOrganized,omitempty"`

	// The number of messages liked in Yammer groups.
	YammerMessagesLiked nullable.Type[int64] `json:"yammerMessagesLiked,omitempty"`

	// The number of messages posted to Yammer groups.
	YammerMessagesPosted nullable.Type[int64] `json:"yammerMessagesPosted,omitempty"`

	// The number of messages read in Yammer groups.
	YammerMessagesRead nullable.Type[int64] `json:"yammerMessagesRead,omitempty"`

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

func (s Office365GroupsActivityCounts) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365GroupsActivityCounts{}

func (s Office365GroupsActivityCounts) MarshalJSON() ([]byte, error) {
	type wrapper Office365GroupsActivityCounts
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365GroupsActivityCounts: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365GroupsActivityCounts: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365GroupsActivityCounts"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365GroupsActivityCounts: %+v", err)
	}

	return encoded, nil
}
