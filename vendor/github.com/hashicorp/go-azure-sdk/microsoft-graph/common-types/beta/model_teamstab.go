package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsTab{}

type TeamsTab struct {
	// Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
	Configuration *TeamsTabConfiguration `json:"configuration,omitempty"`

	// Name of the tab.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	MessageId nullable.Type[string] `json:"messageId,omitempty"`

	// Index of the order used for sorting tabs.
	SortOrderIndex nullable.Type[string] `json:"sortOrderIndex,omitempty"`

	// The application that is linked to the tab.
	TeamsApp *TeamsApp `json:"teamsApp,omitempty"`

	// App definition identifier of the tab. This value can't be changed after tab creation. Because this property is
	// deprecated, we recommend expanding teamsApp to retrieve the application that is linked to the tab.
	TeamsAppId nullable.Type[string] `json:"teamsAppId,omitempty"`

	// Deep link URL of the tab instance. Read only.
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

func (s TeamsTab) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsTab{}

func (s TeamsTab) MarshalJSON() ([]byte, error) {
	type wrapper TeamsTab
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsTab: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsTab: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsTab"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsTab: %+v", err)
	}

	return encoded, nil
}
