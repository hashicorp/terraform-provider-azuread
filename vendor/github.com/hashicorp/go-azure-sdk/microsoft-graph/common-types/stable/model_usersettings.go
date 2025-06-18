package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserSettings{}

type UserSettings struct {
	// Reflects the organization level setting controlling delegate access to the trending API. When set to true, the
	// organization doesn't have access to Office Delve. The relevancy of the content displayed in Microsoft 365, for
	// example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected for
	// the whole organization. This setting is read-only and can only be changed by administrators in the SharePoint admin
	// center.
	ContributionToContentDiscoveryAsOrganizationDisabled *bool `json:"contributionToContentDiscoveryAsOrganizationDisabled,omitempty"`

	// When set to true, the delegate access to the user's trending API is disabled. When set to true, documents in the
	// user's Office Delve are disabled. When set to true, the relevancy of the content displayed in Microsoft 365, for
	// example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected. Users
	// can control this setting in Office Delve.
	ContributionToContentDiscoveryDisabled *bool `json:"contributionToContentDiscoveryDisabled,omitempty"`

	// The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items
	// in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
	ItemInsights *UserInsightsSettings `json:"itemInsights,omitempty"`

	ShiftPreferences *ShiftPreferences `json:"shiftPreferences,omitempty"`
	Storage          *UserStorage      `json:"storage,omitempty"`
	Windows          *[]WindowsSetting `json:"windows,omitempty"`

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

func (s UserSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserSettings{}

func (s UserSettings) MarshalJSON() ([]byte, error) {
	type wrapper UserSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserSettings: %+v", err)
	}

	return encoded, nil
}
