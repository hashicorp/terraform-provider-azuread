package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementConfigurationCategory{}

type DeviceManagementConfigurationCategory struct {
	// Description of the category header in policy summary.
	CategoryDescription nullable.Type[string] `json:"categoryDescription,omitempty"`

	// List of child ids of the category.
	ChildCategoryIds *[]string `json:"childCategoryIds,omitempty"`

	// Description of the category. For example: Display
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the category. For example: Device Lock
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Help text of the category. Give more details of the category.
	HelpText nullable.Type[string] `json:"helpText,omitempty"`

	// Name of the item
	Name nullable.Type[string] `json:"name,omitempty"`

	// Direct parent id of the category. If the category is the root, the parent id is same as its id.
	ParentCategoryId nullable.Type[string] `json:"parentCategoryId,omitempty"`

	// Supported platform types.
	Platforms *DeviceManagementConfigurationPlatforms `json:"platforms,omitempty"`

	// Root id of the category.
	RootCategoryId nullable.Type[string] `json:"rootCategoryId,omitempty"`

	// Supported setting types
	SettingUsage *DeviceManagementConfigurationSettingUsage `json:"settingUsage,omitempty"`

	// Describes which technology this setting can be deployed with
	Technologies *DeviceManagementConfigurationTechnologies `json:"technologies,omitempty"`

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

func (s DeviceManagementConfigurationCategory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationCategory{}

func (s DeviceManagementConfigurationCategory) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationCategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationCategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationCategory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationCategory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationCategory: %+v", err)
	}

	return encoded, nil
}
