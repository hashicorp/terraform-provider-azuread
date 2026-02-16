package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementConfigurationPolicyTemplate{}

type DeviceManagementConfigurationPolicyTemplate struct {
	// Allow unmanaged setting templates
	AllowUnmanagedSettings *bool `json:"allowUnmanagedSettings,omitempty"`

	// Template base identifier
	BaseId nullable.Type[string] `json:"baseId,omitempty"`

	// Template description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Template display name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Description of template version
	DisplayVersion nullable.Type[string] `json:"displayVersion,omitempty"`

	// Describes current lifecycle state of a template
	LifecycleState *DeviceManagementTemplateLifecycleState `json:"lifecycleState,omitempty"`

	// Supported platform types.
	Platforms *DeviceManagementConfigurationPlatforms `json:"platforms,omitempty"`

	// Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
	SettingTemplateCount *int64 `json:"settingTemplateCount,omitempty"`

	// Setting templates
	SettingTemplates *[]DeviceManagementConfigurationSettingTemplate `json:"settingTemplates,omitempty"`

	// Describes which technology this setting can be deployed with
	Technologies *DeviceManagementConfigurationTechnologies `json:"technologies,omitempty"`

	// Describes the TemplateFamily for the Template entity
	TemplateFamily *DeviceManagementConfigurationTemplateFamily `json:"templateFamily,omitempty"`

	// Template version. Valid values 1 to 2147483647. This property is read-only.
	Version *int64 `json:"version,omitempty"`

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

func (s DeviceManagementConfigurationPolicyTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationPolicyTemplate{}

func (s DeviceManagementConfigurationPolicyTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationPolicyTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationPolicyTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationPolicyTemplate: %+v", err)
	}

	delete(decoded, "settingTemplateCount")
	delete(decoded, "version")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationPolicyTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationPolicyTemplate: %+v", err)
	}

	return encoded, nil
}
