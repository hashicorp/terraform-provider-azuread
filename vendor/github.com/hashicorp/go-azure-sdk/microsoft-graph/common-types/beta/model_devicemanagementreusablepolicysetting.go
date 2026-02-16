package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementReusablePolicySetting{}

type DeviceManagementReusablePolicySetting struct {
	// reusable setting creation date and time. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// reusable setting description supplied by user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// reusable setting display name supplied by user.
	DisplayName *string `json:"displayName,omitempty"`

	// date and time when reusable setting was last modified. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// configuration policies referencing the current reusable setting. This property is read-only.
	ReferencingConfigurationPolicies *[]DeviceManagementConfigurationPolicy `json:"referencingConfigurationPolicies,omitempty"`

	// count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property
	// is read-only.
	ReferencingConfigurationPolicyCount *int64 `json:"referencingConfigurationPolicyCount,omitempty"`

	// setting definition id associated with this reusable setting.
	SettingDefinitionId nullable.Type[string] `json:"settingDefinitionId,omitempty"`

	// reusable setting configuration instance
	SettingInstance DeviceManagementConfigurationSettingInstance `json:"settingInstance"`

	// version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
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

func (s DeviceManagementReusablePolicySetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementReusablePolicySetting{}

func (s DeviceManagementReusablePolicySetting) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementReusablePolicySetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementReusablePolicySetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementReusablePolicySetting: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "referencingConfigurationPolicyCount")
	delete(decoded, "version")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementReusablePolicySetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementReusablePolicySetting: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementReusablePolicySetting{}

func (s *DeviceManagementReusablePolicySetting) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime                     *string                                `json:"createdDateTime,omitempty"`
		Description                         nullable.Type[string]                  `json:"description,omitempty"`
		DisplayName                         *string                                `json:"displayName,omitempty"`
		LastModifiedDateTime                *string                                `json:"lastModifiedDateTime,omitempty"`
		ReferencingConfigurationPolicies    *[]DeviceManagementConfigurationPolicy `json:"referencingConfigurationPolicies,omitempty"`
		ReferencingConfigurationPolicyCount *int64                                 `json:"referencingConfigurationPolicyCount,omitempty"`
		SettingDefinitionId                 nullable.Type[string]                  `json:"settingDefinitionId,omitempty"`
		Version                             *int64                                 `json:"version,omitempty"`
		Id                                  *string                                `json:"id,omitempty"`
		ODataId                             *string                                `json:"@odata.id,omitempty"`
		ODataType                           *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ReferencingConfigurationPolicies = decoded.ReferencingConfigurationPolicies
	s.ReferencingConfigurationPolicyCount = decoded.ReferencingConfigurationPolicyCount
	s.SettingDefinitionId = decoded.SettingDefinitionId
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementReusablePolicySetting into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settingInstance"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSettingInstanceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SettingInstance' for 'DeviceManagementReusablePolicySetting': %+v", err)
		}
		s.SettingInstance = impl
	}

	return nil
}
