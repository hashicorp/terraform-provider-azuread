package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementConfigurationPolicy{}

type DeviceManagementConfigurationPolicy struct {
	// Policy assignments
	Assignments *[]DeviceManagementConfigurationPolicyAssignment `json:"assignments,omitempty"`

	// Policy creation date and time
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Policy creation source
	CreationSource nullable.Type[string] `json:"creationSource,omitempty"`

	// Policy description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Policy assignment status. This property is read-only.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// Policy last modification date and time
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Policy name
	Name nullable.Type[string] `json:"name,omitempty"`

	// Supported platform types.
	Platforms *DeviceManagementConfigurationPlatforms `json:"platforms,omitempty"`

	// Indicates the priority of each policies that are selected by the admin during enrollment process
	PriorityMetaData *DeviceManagementPriorityMetaData `json:"priorityMetaData,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Number of settings
	SettingCount *int64 `json:"settingCount,omitempty"`

	// Policy settings
	Settings *[]DeviceManagementConfigurationSetting `json:"settings,omitempty"`

	// Describes which technology this setting can be deployed with
	Technologies *DeviceManagementConfigurationTechnologies `json:"technologies,omitempty"`

	// Template reference information
	TemplateReference *DeviceManagementConfigurationPolicyTemplateReference `json:"templateReference,omitempty"`

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

func (s DeviceManagementConfigurationPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationPolicy{}

func (s DeviceManagementConfigurationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationPolicy: %+v", err)
	}

	delete(decoded, "isAssigned")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationPolicy: %+v", err)
	}

	return encoded, nil
}
