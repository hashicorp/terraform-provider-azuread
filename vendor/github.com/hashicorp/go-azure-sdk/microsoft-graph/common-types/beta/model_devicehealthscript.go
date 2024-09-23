package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceHealthScript{}

type DeviceHealthScript struct {
	// The list of group assignments for the device health script
	Assignments *[]DeviceHealthScriptAssignment `json:"assignments,omitempty"`

	// The timestamp of when the device health script was created. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Description of the device health script
	Description nullable.Type[string] `json:"description,omitempty"`

	// The entire content of the detection powershell script
	DetectionScriptContent nullable.Type[string] `json:"detectionScriptContent,omitempty"`

	// List of ComplexType DetectionScriptParameters objects.
	DetectionScriptParameters *[]DeviceHealthScriptParameter `json:"detectionScriptParameters,omitempty"`

	// Indicates the type of device script.
	DeviceHealthScriptType *DeviceHealthScriptType `json:"deviceHealthScriptType,omitempty"`

	// List of run states for the device health script across all devices
	DeviceRunStates *[]DeviceHealthScriptDeviceState `json:"deviceRunStates,omitempty"`

	// Name of the device health script
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicate whether the script signature needs be checked
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`

	// Highest available version for a Microsoft Proprietary script
	HighestAvailableVersion nullable.Type[string] `json:"highestAvailableVersion,omitempty"`

	// Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
	IsGlobalScript *bool `json:"isGlobalScript,omitempty"`

	// The timestamp of when the device health script was modified. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Name of the device health script publisher
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// The entire content of the remediation powershell script
	RemediationScriptContent nullable.Type[string] `json:"remediationScriptContent,omitempty"`

	// List of ComplexType RemediationScriptParameters objects.
	RemediationScriptParameters *[]DeviceHealthScriptParameter `json:"remediationScriptParameters,omitempty"`

	// List of Scope Tag IDs for the device health script
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicate whether PowerShell script(s) should run as 32-bit
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`

	// Indicates the type of execution context the app runs in.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`

	// High level run summary for device health script.
	RunSummary *DeviceHealthScriptRunSummary `json:"runSummary,omitempty"`

	// Version of the device health script
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

func (s DeviceHealthScript) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScript{}

func (s DeviceHealthScript) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScript
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScript: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScript: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScript"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScript: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceHealthScript{}

func (s *DeviceHealthScript) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Assignments              *[]DeviceHealthScriptAssignment  `json:"assignments,omitempty"`
		CreatedDateTime          *string                          `json:"createdDateTime,omitempty"`
		Description              nullable.Type[string]            `json:"description,omitempty"`
		DetectionScriptContent   nullable.Type[string]            `json:"detectionScriptContent,omitempty"`
		DeviceHealthScriptType   *DeviceHealthScriptType          `json:"deviceHealthScriptType,omitempty"`
		DeviceRunStates          *[]DeviceHealthScriptDeviceState `json:"deviceRunStates,omitempty"`
		DisplayName              nullable.Type[string]            `json:"displayName,omitempty"`
		EnforceSignatureCheck    *bool                            `json:"enforceSignatureCheck,omitempty"`
		HighestAvailableVersion  nullable.Type[string]            `json:"highestAvailableVersion,omitempty"`
		IsGlobalScript           *bool                            `json:"isGlobalScript,omitempty"`
		LastModifiedDateTime     *string                          `json:"lastModifiedDateTime,omitempty"`
		Publisher                nullable.Type[string]            `json:"publisher,omitempty"`
		RemediationScriptContent nullable.Type[string]            `json:"remediationScriptContent,omitempty"`
		RoleScopeTagIds          *[]string                        `json:"roleScopeTagIds,omitempty"`
		RunAs32Bit               *bool                            `json:"runAs32Bit,omitempty"`
		RunAsAccount             *RunAsAccountType                `json:"runAsAccount,omitempty"`
		RunSummary               *DeviceHealthScriptRunSummary    `json:"runSummary,omitempty"`
		Version                  nullable.Type[string]            `json:"version,omitempty"`
		Id                       *string                          `json:"id,omitempty"`
		ODataId                  *string                          `json:"@odata.id,omitempty"`
		ODataType                *string                          `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Assignments = decoded.Assignments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DetectionScriptContent = decoded.DetectionScriptContent
	s.DeviceHealthScriptType = decoded.DeviceHealthScriptType
	s.DeviceRunStates = decoded.DeviceRunStates
	s.DisplayName = decoded.DisplayName
	s.EnforceSignatureCheck = decoded.EnforceSignatureCheck
	s.HighestAvailableVersion = decoded.HighestAvailableVersion
	s.IsGlobalScript = decoded.IsGlobalScript
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Publisher = decoded.Publisher
	s.RemediationScriptContent = decoded.RemediationScriptContent
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.RunAs32Bit = decoded.RunAs32Bit
	s.RunAsAccount = decoded.RunAsAccount
	s.RunSummary = decoded.RunSummary
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceHealthScript into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["detectionScriptParameters"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DetectionScriptParameters into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceHealthScriptParameter, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceHealthScriptParameterImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DetectionScriptParameters' for 'DeviceHealthScript': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DetectionScriptParameters = &output
	}

	if v, ok := temp["remediationScriptParameters"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RemediationScriptParameters into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceHealthScriptParameter, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceHealthScriptParameterImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RemediationScriptParameters' for 'DeviceHealthScript': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RemediationScriptParameters = &output
	}

	return nil
}
