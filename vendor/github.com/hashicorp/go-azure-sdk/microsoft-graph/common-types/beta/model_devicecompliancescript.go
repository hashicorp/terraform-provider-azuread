package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceComplianceScript{}

type DeviceComplianceScript struct {
	// The list of group assignments for the device compliance script
	Assignments *[]DeviceHealthScriptAssignment `json:"assignments,omitempty"`

	// The timestamp of when the device compliance script was created. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Description of the device compliance script
	Description nullable.Type[string] `json:"description,omitempty"`

	// The entire content of the detection powershell script
	DetectionScriptContent nullable.Type[string] `json:"detectionScriptContent,omitempty"`

	// List of run states for the device compliance script across all devices
	DeviceRunStates *[]DeviceComplianceScriptDeviceState `json:"deviceRunStates,omitempty"`

	// Name of the device compliance script
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicate whether the script signature needs be checked
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`

	// The timestamp of when the device compliance script was modified. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Name of the device compliance script publisher
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// List of Scope Tag IDs for the device compliance script
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicate whether PowerShell script(s) should run as 32-bit
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`

	// Indicates the type of execution context the app runs in.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`

	// High level run summary for device compliance script.
	RunSummary *DeviceComplianceScriptRunSummary `json:"runSummary,omitempty"`

	// Version of the device compliance script
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s DeviceComplianceScript) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceComplianceScript{}

func (s DeviceComplianceScript) MarshalJSON() ([]byte, error) {
	type wrapper DeviceComplianceScript
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceComplianceScript: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceComplianceScript: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceComplianceScript"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceComplianceScript: %+v", err)
	}

	return encoded, nil
}
