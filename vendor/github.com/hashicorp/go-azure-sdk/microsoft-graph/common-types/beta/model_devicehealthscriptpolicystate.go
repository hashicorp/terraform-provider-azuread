package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptPolicyState struct {
	// A list of the assignment filter ids used for health script applicability evaluation
	AssignmentFilterIds *[]string `json:"assignmentFilterIds,omitempty"`

	// Indicates the type of execution status of the device management script.
	DetectionState *RunState `json:"detectionState,omitempty"`

	// The Intune device Id
	DeviceId *string `json:"deviceId,omitempty"`

	// Display name of the device
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The next timestamp of when the device health script is expected to execute
	ExpectedStateUpdateDateTime nullable.Type[string] `json:"expectedStateUpdateDateTime,omitempty"`

	// Key of the device health script policy state is a concatenation of the MT sideCar policy Id and Intune device Id
	Id *string `json:"id,omitempty"`

	// The last timestamp of when the device health script executed
	LastStateUpdateDateTime *string `json:"lastStateUpdateDateTime,omitempty"`

	// The last time that Intune Managment Extension synced with Intune
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value of the OS Version in string
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// The MT sideCar policy Id
	PolicyId *string `json:"policyId,omitempty"`

	// Display name of the device health script
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`

	// Error from the detection script after remediation
	PostRemediationDetectionScriptError nullable.Type[string] `json:"postRemediationDetectionScriptError,omitempty"`

	// Detection script output after remediation
	PostRemediationDetectionScriptOutput nullable.Type[string] `json:"postRemediationDetectionScriptOutput,omitempty"`

	// Error from the detection script before remediation
	PreRemediationDetectionScriptError nullable.Type[string] `json:"preRemediationDetectionScriptError,omitempty"`

	// Output of the detection script before remediation
	PreRemediationDetectionScriptOutput nullable.Type[string] `json:"preRemediationDetectionScriptOutput,omitempty"`

	// Error output of the remediation script
	RemediationScriptError nullable.Type[string] `json:"remediationScriptError,omitempty"`

	// Indicates the type of execution status of the device management script.
	RemediationState *RemediationState `json:"remediationState,omitempty"`

	// Name of the user whom ran the device health script
	UserName nullable.Type[string] `json:"userName,omitempty"`
}
