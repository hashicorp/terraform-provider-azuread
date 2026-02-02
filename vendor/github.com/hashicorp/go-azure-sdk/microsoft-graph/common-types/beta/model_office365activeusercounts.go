package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365ActiveUserCounts{}

type Office365ActiveUserCounts struct {
	// The number of active users in Exchange. Any user who can read and send email is considered an active user.
	Exchange nullable.Type[int64] `json:"exchange,omitempty"`

	// The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive,
	// SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each
	// product in the respective property description.
	Office365 nullable.Type[int64] `json:"office365,omitempty"`

	// The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally,
	// or synced files is considered an active user.
	OneDrive nullable.Type[int64] `json:"oneDrive,omitempty"`

	// The date on which a number of users were active.
	ReportDate nullable.Type[string] `json:"reportDate,omitempty"`

	// The number of days the report covers.
	ReportPeriod nullable.Type[string] `json:"reportPeriod,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally,
	// synced files, or viewed SharePoint pages is considered an active user.
	SharePoint nullable.Type[int64] `json:"sharePoint,omitempty"`

	// The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined
	// peer-to-peer sessions is considered an active user.
	SkypeForBusiness nullable.Type[int64] `json:"skypeForBusiness,omitempty"`

	// The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in
	// private chat sessions, or participated in meetings or calls is considered an active user.
	Teams nullable.Type[int64] `json:"teams,omitempty"`

	// The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
	Yammer nullable.Type[int64] `json:"yammer,omitempty"`

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

func (s Office365ActiveUserCounts) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365ActiveUserCounts{}

func (s Office365ActiveUserCounts) MarshalJSON() ([]byte, error) {
	type wrapper Office365ActiveUserCounts
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365ActiveUserCounts: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365ActiveUserCounts: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365ActiveUserCounts"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365ActiveUserCounts: %+v", err)
	}

	return encoded, nil
}
