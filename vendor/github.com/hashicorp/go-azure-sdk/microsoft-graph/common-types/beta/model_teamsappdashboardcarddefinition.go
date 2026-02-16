package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAppDashboardCardDefinition{}

type TeamsAppDashboardCardDefinition struct {
	// The configuration for the source of the card content. Required.
	ContentSource TeamsAppDashboardCardContentSource `json:"contentSource"`

	// The size of the card. The possible values are: medium, large, unknownFutureValue. Required.
	DefaultSize TeamsAppDashboardCardSize `json:"defaultSize"`

	// The description for the card. Required.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the card. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Configuration for the display of the icon in the card picker. If neither this nor any of its properties (iconUrl and
	// officeUIFabricIconName) are specified, the color icon of the app is used. Optional.
	Icon *TeamsAppDashboardCardIcon `json:"icon,omitempty"`

	// ID for the group in the card picker. Required.
	PickerGroupId nullable.Type[string] `json:"pickerGroupId,omitempty"`

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

func (s TeamsAppDashboardCardDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAppDashboardCardDefinition{}

func (s TeamsAppDashboardCardDefinition) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAppDashboardCardDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAppDashboardCardDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppDashboardCardDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAppDashboardCardDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAppDashboardCardDefinition: %+v", err)
	}

	return encoded, nil
}
