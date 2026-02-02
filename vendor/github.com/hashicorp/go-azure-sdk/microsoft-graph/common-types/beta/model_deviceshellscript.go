package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceShellScript{}

type DeviceShellScript struct {
	// The list of group assignments for the device management script.
	Assignments *[]DeviceManagementScriptAssignment `json:"assignments,omitempty"`

	// Does not notify the user a script is being executed
	BlockExecutionNotifications *bool `json:"blockExecutionNotifications,omitempty"`

	// The date and time the device management script was created. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Optional description for the device management script.
	Description nullable.Type[string] `json:"description,omitempty"`

	// List of run states for this script across all devices.
	DeviceRunStates *[]DeviceManagementScriptDeviceState `json:"deviceRunStates,omitempty"`

	// Name of the device management script.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The interval for script to run. If not defined the script will run once
	ExecutionFrequency *string `json:"executionFrequency,omitempty"`

	// Script file name.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The list of group assignments for the device management script.
	GroupAssignments *[]DeviceManagementScriptGroupAssignment `json:"groupAssignments,omitempty"`

	// The date and time the device management script was last modified. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Number of times for the script to be retried if it fails
	RetryCount *int64 `json:"retryCount,omitempty"`

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

func (s DeviceShellScript) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceShellScript{}

func (s DeviceShellScript) MarshalJSON() ([]byte, error) {
	type wrapper DeviceShellScript
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceShellScript: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceShellScript: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceShellScript"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceShellScript: %+v", err)
	}

	return encoded, nil
}
