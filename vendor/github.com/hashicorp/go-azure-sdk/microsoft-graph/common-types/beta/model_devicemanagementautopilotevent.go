package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementAutopilotEvent{}

type DeviceManagementAutopilotEvent struct {
	// Time spent in user ESP.
	AccountSetupDuration *string `json:"accountSetupDuration,omitempty"`

	// Deployment states for Autopilot devices
	AccountSetupStatus *WindowsAutopilotDeploymentState `json:"accountSetupStatus,omitempty"`

	// Autopilot deployment duration including enrollment.
	DeploymentDuration *string `json:"deploymentDuration,omitempty"`

	// Deployment end time.
	DeploymentEndDateTime *string `json:"deploymentEndDateTime,omitempty"`

	// Deployment start time.
	DeploymentStartDateTime *string `json:"deploymentStartDateTime,omitempty"`

	// Deployment states for Autopilot devices
	DeploymentState *WindowsAutopilotDeploymentState `json:"deploymentState,omitempty"`

	// Total deployment duration from enrollment to Desktop screen.
	DeploymentTotalDuration *string `json:"deploymentTotalDuration,omitempty"`

	// Device id associated with the object
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Time spent in device enrollment.
	DevicePreparationDuration *string `json:"devicePreparationDuration,omitempty"`

	// Device registration date.
	DeviceRegisteredDateTime *string `json:"deviceRegisteredDateTime,omitempty"`

	// Device serial number.
	DeviceSerialNumber nullable.Type[string] `json:"deviceSerialNumber,omitempty"`

	// Time spent in device ESP.
	DeviceSetupDuration *string `json:"deviceSetupDuration,omitempty"`

	// Deployment states for Autopilot devices
	DeviceSetupStatus *WindowsAutopilotDeploymentState `json:"deviceSetupStatus,omitempty"`

	// Enrollment failure details.
	EnrollmentFailureDetails nullable.Type[string] `json:"enrollmentFailureDetails,omitempty"`

	// Device enrollment start date.
	EnrollmentStartDateTime *string `json:"enrollmentStartDateTime,omitempty"`

	EnrollmentState *EnrollmentState                `json:"enrollmentState,omitempty"`
	EnrollmentType  *WindowsAutopilotEnrollmentType `json:"enrollmentType,omitempty"`

	// Time when the event occurred .
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// Managed device name.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// Device operating system version.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Policy and application status details for this device.
	PolicyStatusDetails *[]DeviceManagementAutopilotPolicyStatusDetail `json:"policyStatusDetails,omitempty"`

	// Count of applications targeted.
	TargetedAppCount *int64 `json:"targetedAppCount,omitempty"`

	// Count of policies targeted.
	TargetedPolicyCount *int64 `json:"targetedPolicyCount,omitempty"`

	// User principal name used to enroll the device.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Enrollment Status Page profile name
	Windows10EnrollmentCompletionPageConfigurationDisplayName nullable.Type[string] `json:"windows10EnrollmentCompletionPageConfigurationDisplayName,omitempty"`

	// Enrollment Status Page profile ID
	Windows10EnrollmentCompletionPageConfigurationId nullable.Type[string] `json:"windows10EnrollmentCompletionPageConfigurationId,omitempty"`

	// Autopilot profile name.
	WindowsAutopilotDeploymentProfileDisplayName nullable.Type[string] `json:"windowsAutopilotDeploymentProfileDisplayName,omitempty"`

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

func (s DeviceManagementAutopilotEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementAutopilotEvent{}

func (s DeviceManagementAutopilotEvent) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementAutopilotEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementAutopilotEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementAutopilotEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementAutopilotEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementAutopilotEvent: %+v", err)
	}

	return encoded, nil
}
