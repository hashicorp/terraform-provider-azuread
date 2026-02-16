package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceCustomAttributeShellScript{}

type DeviceCustomAttributeShellScript struct {
	// The list of group assignments for the device management script.
	Assignments *[]DeviceManagementScriptAssignment `json:"assignments,omitempty"`

	// The date and time the device management script was created. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The name of the custom attribute.
	CustomAttributeName nullable.Type[string] `json:"customAttributeName,omitempty"`

	// Represents the expected type for a macOS custom attribute script value.
	CustomAttributeType *DeviceCustomAttributeValueType `json:"customAttributeType,omitempty"`

	// Optional description for the device management script.
	Description nullable.Type[string] `json:"description,omitempty"`

	// List of run states for this script across all devices.
	DeviceRunStates *[]DeviceManagementScriptDeviceState `json:"deviceRunStates,omitempty"`

	// Name of the device management script.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Script file name.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The list of group assignments for the device management script.
	GroupAssignments *[]DeviceManagementScriptGroupAssignment `json:"groupAssignments,omitempty"`

	// The date and time the device management script was last modified. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tag IDs for this PowerShellScript instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates the type of execution context the app runs in.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`

	// Run summary for device management script.
	RunSummary *DeviceManagementScriptRunSummary `json:"runSummary,omitempty"`

	// The script content.
	ScriptContent nullable.Type[string] `json:"scriptContent,omitempty"`

	// List of run states for this script across all users.
	UserRunStates *[]DeviceManagementScriptUserState `json:"userRunStates,omitempty"`

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

func (s DeviceCustomAttributeShellScript) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceCustomAttributeShellScript{}

func (s DeviceCustomAttributeShellScript) MarshalJSON() ([]byte, error) {
	type wrapper DeviceCustomAttributeShellScript
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceCustomAttributeShellScript: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceCustomAttributeShellScript: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceCustomAttributeShellScript"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceCustomAttributeShellScript: %+v", err)
	}

	return encoded, nil
}
