package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementIntent{}

type DeviceManagementIntent struct {
	// Collection of assignments
	Assignments *[]DeviceManagementIntentAssignment `json:"assignments,omitempty"`

	// Collection of setting categories within the intent
	Categories *[]DeviceManagementIntentSettingCategory `json:"categories,omitempty"`

	// The user given description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Collection of settings and their states and counts of devices that belong to corresponding state for all settings
	// within the intent
	DeviceSettingStateSummaries *[]DeviceManagementIntentDeviceSettingStateSummary `json:"deviceSettingStateSummaries,omitempty"`

	// A summary of device states and counts of devices that belong to corresponding state for all devices that the intent
	// is applied to
	DeviceStateSummary *DeviceManagementIntentDeviceStateSummary `json:"deviceStateSummary,omitempty"`

	// Collection of states of all devices that the intent is applied to
	DeviceStates *[]DeviceManagementIntentDeviceState `json:"deviceStates,omitempty"`

	// The user given display name
	DisplayName *string `json:"displayName,omitempty"`

	// Signifies whether or not the intent is assigned to users
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// Signifies whether or not the intent is being migrated to the configurationPolicies endpoint
	IsMigratingToConfigurationPolicy nullable.Type[bool] `json:"isMigratingToConfigurationPolicy,omitempty"`

	// When the intent was last modified
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Collection of all settings to be applied
	Settings *[]DeviceManagementSettingInstance `json:"settings,omitempty"`

	// The ID of the template this intent was created from (if any)
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`

	// A summary of user states and counts of users that belong to corresponding state for all users that the intent is
	// applied to
	UserStateSummary *DeviceManagementIntentUserStateSummary `json:"userStateSummary,omitempty"`

	// Collection of states of all users that the intent is applied to
	UserStates *[]DeviceManagementIntentUserState `json:"userStates,omitempty"`

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

func (s DeviceManagementIntent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementIntent{}

func (s DeviceManagementIntent) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementIntent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementIntent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementIntent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementIntent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementIntent: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementIntent{}

func (s *DeviceManagementIntent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Assignments                      *[]DeviceManagementIntentAssignment                `json:"assignments,omitempty"`
		Categories                       *[]DeviceManagementIntentSettingCategory           `json:"categories,omitempty"`
		Description                      nullable.Type[string]                              `json:"description,omitempty"`
		DeviceSettingStateSummaries      *[]DeviceManagementIntentDeviceSettingStateSummary `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStateSummary               *DeviceManagementIntentDeviceStateSummary          `json:"deviceStateSummary,omitempty"`
		DeviceStates                     *[]DeviceManagementIntentDeviceState               `json:"deviceStates,omitempty"`
		DisplayName                      *string                                            `json:"displayName,omitempty"`
		IsAssigned                       *bool                                              `json:"isAssigned,omitempty"`
		IsMigratingToConfigurationPolicy nullable.Type[bool]                                `json:"isMigratingToConfigurationPolicy,omitempty"`
		LastModifiedDateTime             *string                                            `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                  *[]string                                          `json:"roleScopeTagIds,omitempty"`
		TemplateId                       nullable.Type[string]                              `json:"templateId,omitempty"`
		UserStateSummary                 *DeviceManagementIntentUserStateSummary            `json:"userStateSummary,omitempty"`
		UserStates                       *[]DeviceManagementIntentUserState                 `json:"userStates,omitempty"`
		Id                               *string                                            `json:"id,omitempty"`
		ODataId                          *string                                            `json:"@odata.id,omitempty"`
		ODataType                        *string                                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Assignments = decoded.Assignments
	s.Categories = decoded.Categories
	s.Description = decoded.Description
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStateSummary = decoded.DeviceStateSummary
	s.DeviceStates = decoded.DeviceStates
	s.DisplayName = decoded.DisplayName
	s.IsAssigned = decoded.IsAssigned
	s.IsMigratingToConfigurationPolicy = decoded.IsMigratingToConfigurationPolicy
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.TemplateId = decoded.TemplateId
	s.UserStateSummary = decoded.UserStateSummary
	s.UserStates = decoded.UserStates
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementIntent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Settings into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingInstance, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingInstanceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Settings' for 'DeviceManagementIntent': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Settings = &output
	}

	return nil
}
