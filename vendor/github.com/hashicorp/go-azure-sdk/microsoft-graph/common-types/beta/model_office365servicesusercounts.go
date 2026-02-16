package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365ServicesUserCounts{}

type Office365ServicesUserCounts struct {
	// The number of active users on Exchange. Any user who can read and send email is considered an active user.
	ExchangeActive nullable.Type[int64] `json:"exchangeActive,omitempty"`

	// The number of inactive users on Exchange.
	ExchangeInactive nullable.Type[int64] `json:"exchangeInactive,omitempty"`

	// The number of active users on Microsoft 365.
	Office365Active nullable.Type[int64] `json:"office365Active,omitempty"`

	// The number of inactive users on Microsoft 365.
	Office365Inactive nullable.Type[int64] `json:"office365Inactive,omitempty"`

	// The number of active users on OneDrive. Any user who viewed or edited files, shared files internally or externally,
	// or synced files is considered an active user.
	OneDriveActive nullable.Type[int64] `json:"oneDriveActive,omitempty"`

	// The number of inactive users on OneDrive.
	OneDriveInactive nullable.Type[int64] `json:"oneDriveInactive,omitempty"`

	// The number of days the report covers.
	ReportPeriod nullable.Type[string] `json:"reportPeriod,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The number of active users on SharePoint. Any user who viewed or edited files, shared files internally or externally,
	// synced files, or viewed SharePoint pages is considered an active user.
	SharePointActive nullable.Type[int64] `json:"sharePointActive,omitempty"`

	// The number of inactive users on SharePoint.
	SharePointInactive nullable.Type[int64] `json:"sharePointInactive,omitempty"`

	// The number of active users on Skype For Business. Any user who organized or participated in conferences, or joined
	// peer-to-peer sessions is considered an active user.
	SkypeForBusinessActive nullable.Type[int64] `json:"skypeForBusinessActive,omitempty"`

	// The number of inactive users on Skype For Business.
	SkypeForBusinessInactive nullable.Type[int64] `json:"skypeForBusinessInactive,omitempty"`

	// The number of active users on Microsoft Teams. Any user who posted messages in team channels, sent messages in
	// private chat sessions, or participated in meetings or calls is considered an active user.
	TeamsActive nullable.Type[int64] `json:"teamsActive,omitempty"`

	// The number of inactive users on Microsoft Teams.
	TeamsInactive nullable.Type[int64] `json:"teamsInactive,omitempty"`

	// The number of active users on Yammer. Any user who can post, read, or like messages is considered an active user.
	YammerActive nullable.Type[int64] `json:"yammerActive,omitempty"`

	// The number of inactive users on Yammer.
	YammerInactive nullable.Type[int64] `json:"yammerInactive,omitempty"`

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

func (s Office365ServicesUserCounts) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365ServicesUserCounts{}

func (s Office365ServicesUserCounts) MarshalJSON() ([]byte, error) {
	type wrapper Office365ServicesUserCounts
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365ServicesUserCounts: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365ServicesUserCounts: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365ServicesUserCounts"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365ServicesUserCounts: %+v", err)
	}

	return encoded, nil
}
