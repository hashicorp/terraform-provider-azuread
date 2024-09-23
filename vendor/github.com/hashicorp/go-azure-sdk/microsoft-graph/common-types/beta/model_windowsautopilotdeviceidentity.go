package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsAutopilotDeviceIdentity{}

type WindowsAutopilotDeviceIdentity struct {
	// Addressable user name.
	AddressableUserName nullable.Type[string] `json:"addressableUserName,omitempty"`

	// AAD Device ID - to be deprecated
	AzureActiveDirectoryDeviceId nullable.Type[string] `json:"azureActiveDirectoryDeviceId,omitempty"`

	// AAD Device ID
	AzureAdDeviceId nullable.Type[string] `json:"azureAdDeviceId,omitempty"`

	// Deployment profile currently assigned to the Windows autopilot device.
	DeploymentProfile *WindowsAutopilotDeploymentProfile `json:"deploymentProfile,omitempty"`

	// Profile set time of the Windows autopilot device.
	DeploymentProfileAssignedDateTime *string `json:"deploymentProfileAssignedDateTime,omitempty"`

	DeploymentProfileAssignmentDetailedStatus *WindowsAutopilotProfileAssignmentDetailedStatus `json:"deploymentProfileAssignmentDetailedStatus,omitempty"`
	DeploymentProfileAssignmentStatus         *WindowsAutopilotProfileAssignmentStatus         `json:"deploymentProfileAssignmentStatus,omitempty"`

	// Surface Hub Device Account Password
	DeviceAccountPassword nullable.Type[string] `json:"deviceAccountPassword,omitempty"`

	// Surface Hub Device Account Upn
	DeviceAccountUpn nullable.Type[string] `json:"deviceAccountUpn,omitempty"`

	// Surface Hub Device Friendly Name
	DeviceFriendlyName nullable.Type[string] `json:"deviceFriendlyName,omitempty"`

	// Display Name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	EnrollmentState *EnrollmentState `json:"enrollmentState,omitempty"`

	// Group Tag of the Windows autopilot device.
	GroupTag nullable.Type[string] `json:"groupTag,omitempty"`

	// Deployment profile intended to be assigned to the Windows autopilot device.
	IntendedDeploymentProfile *WindowsAutopilotDeploymentProfile `json:"intendedDeploymentProfile,omitempty"`

	// Intune Last Contacted Date Time of the Windows autopilot device.
	LastContactedDateTime *string `json:"lastContactedDateTime,omitempty"`

	// Managed Device ID
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// Oem manufacturer of the Windows autopilot device.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Model name of the Windows autopilot device.
	Model nullable.Type[string] `json:"model,omitempty"`

	// Product Key of the Windows autopilot device.
	ProductKey nullable.Type[string] `json:"productKey,omitempty"`

	// Purchase Order Identifier of the Windows autopilot device.
	PurchaseOrderIdentifier nullable.Type[string] `json:"purchaseOrderIdentifier,omitempty"`

	// Device remediation status, indicating whether or not hardware has been changed for an Autopilot-registered device.
	RemediationState *WindowsAutopilotDeviceRemediationState `json:"remediationState,omitempty"`

	// RemediationState set time of Autopilot device.
	RemediationStateLastModifiedDateTime *string `json:"remediationStateLastModifiedDateTime,omitempty"`

	// Resource Name.
	ResourceName nullable.Type[string] `json:"resourceName,omitempty"`

	// Serial number of the Windows autopilot device.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// SKU Number
	SkuNumber nullable.Type[string] `json:"skuNumber,omitempty"`

	// System Family
	SystemFamily nullable.Type[string] `json:"systemFamily,omitempty"`

	// User Principal Name.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Userless enrollment block status, indicating whether the next device enrollment will be blocked.
	UserlessEnrollmentStatus *WindowsAutopilotUserlessEnrollmentStatus `json:"userlessEnrollmentStatus,omitempty"`

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

func (s WindowsAutopilotDeviceIdentity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsAutopilotDeviceIdentity{}

func (s WindowsAutopilotDeviceIdentity) MarshalJSON() ([]byte, error) {
	type wrapper WindowsAutopilotDeviceIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsAutopilotDeviceIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsAutopilotDeviceIdentity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsAutopilotDeviceIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsAutopilotDeviceIdentity: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsAutopilotDeviceIdentity{}

func (s *WindowsAutopilotDeviceIdentity) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AddressableUserName                       nullable.Type[string]                            `json:"addressableUserName,omitempty"`
		AzureActiveDirectoryDeviceId              nullable.Type[string]                            `json:"azureActiveDirectoryDeviceId,omitempty"`
		AzureAdDeviceId                           nullable.Type[string]                            `json:"azureAdDeviceId,omitempty"`
		DeploymentProfileAssignedDateTime         *string                                          `json:"deploymentProfileAssignedDateTime,omitempty"`
		DeploymentProfileAssignmentDetailedStatus *WindowsAutopilotProfileAssignmentDetailedStatus `json:"deploymentProfileAssignmentDetailedStatus,omitempty"`
		DeploymentProfileAssignmentStatus         *WindowsAutopilotProfileAssignmentStatus         `json:"deploymentProfileAssignmentStatus,omitempty"`
		DeviceAccountPassword                     nullable.Type[string]                            `json:"deviceAccountPassword,omitempty"`
		DeviceAccountUpn                          nullable.Type[string]                            `json:"deviceAccountUpn,omitempty"`
		DeviceFriendlyName                        nullable.Type[string]                            `json:"deviceFriendlyName,omitempty"`
		DisplayName                               nullable.Type[string]                            `json:"displayName,omitempty"`
		EnrollmentState                           *EnrollmentState                                 `json:"enrollmentState,omitempty"`
		GroupTag                                  nullable.Type[string]                            `json:"groupTag,omitempty"`
		LastContactedDateTime                     *string                                          `json:"lastContactedDateTime,omitempty"`
		ManagedDeviceId                           nullable.Type[string]                            `json:"managedDeviceId,omitempty"`
		Manufacturer                              nullable.Type[string]                            `json:"manufacturer,omitempty"`
		Model                                     nullable.Type[string]                            `json:"model,omitempty"`
		ProductKey                                nullable.Type[string]                            `json:"productKey,omitempty"`
		PurchaseOrderIdentifier                   nullable.Type[string]                            `json:"purchaseOrderIdentifier,omitempty"`
		RemediationState                          *WindowsAutopilotDeviceRemediationState          `json:"remediationState,omitempty"`
		RemediationStateLastModifiedDateTime      *string                                          `json:"remediationStateLastModifiedDateTime,omitempty"`
		ResourceName                              nullable.Type[string]                            `json:"resourceName,omitempty"`
		SerialNumber                              nullable.Type[string]                            `json:"serialNumber,omitempty"`
		SkuNumber                                 nullable.Type[string]                            `json:"skuNumber,omitempty"`
		SystemFamily                              nullable.Type[string]                            `json:"systemFamily,omitempty"`
		UserPrincipalName                         nullable.Type[string]                            `json:"userPrincipalName,omitempty"`
		UserlessEnrollmentStatus                  *WindowsAutopilotUserlessEnrollmentStatus        `json:"userlessEnrollmentStatus,omitempty"`
		Id                                        *string                                          `json:"id,omitempty"`
		ODataId                                   *string                                          `json:"@odata.id,omitempty"`
		ODataType                                 *string                                          `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AddressableUserName = decoded.AddressableUserName
	s.AzureActiveDirectoryDeviceId = decoded.AzureActiveDirectoryDeviceId
	s.AzureAdDeviceId = decoded.AzureAdDeviceId
	s.DeploymentProfileAssignedDateTime = decoded.DeploymentProfileAssignedDateTime
	s.DeploymentProfileAssignmentDetailedStatus = decoded.DeploymentProfileAssignmentDetailedStatus
	s.DeploymentProfileAssignmentStatus = decoded.DeploymentProfileAssignmentStatus
	s.DeviceAccountPassword = decoded.DeviceAccountPassword
	s.DeviceAccountUpn = decoded.DeviceAccountUpn
	s.DeviceFriendlyName = decoded.DeviceFriendlyName
	s.DisplayName = decoded.DisplayName
	s.EnrollmentState = decoded.EnrollmentState
	s.GroupTag = decoded.GroupTag
	s.LastContactedDateTime = decoded.LastContactedDateTime
	s.ManagedDeviceId = decoded.ManagedDeviceId
	s.Manufacturer = decoded.Manufacturer
	s.Model = decoded.Model
	s.ProductKey = decoded.ProductKey
	s.PurchaseOrderIdentifier = decoded.PurchaseOrderIdentifier
	s.RemediationState = decoded.RemediationState
	s.RemediationStateLastModifiedDateTime = decoded.RemediationStateLastModifiedDateTime
	s.ResourceName = decoded.ResourceName
	s.SerialNumber = decoded.SerialNumber
	s.SkuNumber = decoded.SkuNumber
	s.SystemFamily = decoded.SystemFamily
	s.UserPrincipalName = decoded.UserPrincipalName
	s.UserlessEnrollmentStatus = decoded.UserlessEnrollmentStatus
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsAutopilotDeviceIdentity into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deploymentProfile"]; ok {
		impl, err := UnmarshalWindowsAutopilotDeploymentProfileImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DeploymentProfile' for 'WindowsAutopilotDeviceIdentity': %+v", err)
		}
		s.DeploymentProfile = &impl
	}

	if v, ok := temp["intendedDeploymentProfile"]; ok {
		impl, err := UnmarshalWindowsAutopilotDeploymentProfileImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IntendedDeploymentProfile' for 'WindowsAutopilotDeviceIdentity': %+v", err)
		}
		s.IntendedDeploymentProfile = &impl
	}

	return nil
}
