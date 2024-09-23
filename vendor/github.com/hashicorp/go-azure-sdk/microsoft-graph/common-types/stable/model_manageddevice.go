package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDevice{}

type ManagedDevice struct {
	// The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property)
	// for this property when returned as part of managedDevice entity in LIST call. To retrieve actual values GET call
	// needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
	// Read-only. This property is read-only.
	ActivationLockBypassCode nullable.Type[string] `json:"activationLockBypassCode,omitempty"`

	// Android security patch level. This property is read-only.
	AndroidSecurityPatchLevel nullable.Type[string] `json:"androidSecurityPatchLevel,omitempty"`

	// The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
	AzureADDeviceId nullable.Type[string] `json:"azureADDeviceId,omitempty"`

	// Whether the device is Azure Active Directory registered. This property is read-only.
	AzureADRegistered nullable.Type[bool] `json:"azureADRegistered,omitempty"`

	// The DateTime when device compliance grace period expires. This property is read-only.
	ComplianceGracePeriodExpirationDateTime *string `json:"complianceGracePeriodExpirationDateTime,omitempty"`

	// Compliance state.
	ComplianceState *ComplianceState `json:"complianceState,omitempty"`

	// ConfigrMgr client enabled features. This property is read-only.
	ConfigurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures `json:"configurationManagerClientEnabledFeatures,omitempty"`

	// List of ComplexType deviceActionResult objects. This property is read-only.
	DeviceActionResults *[]DeviceActionResult `json:"deviceActionResults,omitempty"`

	// Device category
	DeviceCategory *DeviceCategory `json:"deviceCategory,omitempty"`

	// Device category display name. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is
	// read-only.
	DeviceCategoryDisplayName nullable.Type[string] `json:"deviceCategoryDisplayName,omitempty"`

	// Device compliance policy states for this device.
	DeviceCompliancePolicyStates *[]DeviceCompliancePolicyState `json:"deviceCompliancePolicyStates,omitempty"`

	// Device configuration states for this device.
	DeviceConfigurationStates *[]DeviceConfigurationState `json:"deviceConfigurationStates,omitempty"`

	// Possible ways of adding a mobile device to management.
	DeviceEnrollmentType *DeviceEnrollmentType `json:"deviceEnrollmentType,omitempty"`

	// The device health attestation state. This property is read-only.
	DeviceHealthAttestationState *DeviceHealthAttestationState `json:"deviceHealthAttestationState,omitempty"`

	// Name of the device. This property is read-only.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Device registration status.
	DeviceRegistrationState *DeviceRegistrationState `json:"deviceRegistrationState,omitempty"`

	// Whether the device is Exchange ActiveSync activated. This property is read-only.
	EasActivated *bool `json:"easActivated,omitempty"`

	// Exchange ActivationSync activation time of the device. This property is read-only.
	EasActivationDateTime *string `json:"easActivationDateTime,omitempty"`

	// Exchange ActiveSync Id of the device. This property is read-only.
	EasDeviceId nullable.Type[string] `json:"easDeviceId,omitempty"`

	// Email(s) for the user associated with the device. This property is read-only.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Enrollment time of the device. Supports $filter operator 'lt' and 'gt'. This property is read-only.
	EnrolledDateTime *string `json:"enrolledDateTime,omitempty"`

	// Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment
	// profile was assgined. This property is read-only.
	EnrollmentProfileName nullable.Type[string] `json:"enrollmentProfileName,omitempty"`

	// Indicates Ethernet MAC Address of the device. Default, is Null (Non-Default property) for this property when returned
	// as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values.
	// Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is
	// not supported. Read-only. This property is read-only.
	EthernetMacAddress nullable.Type[string] `json:"ethernetMacAddress,omitempty"`

	// Device Exchange Access State.
	ExchangeAccessState *DeviceManagementExchangeAccessState `json:"exchangeAccessState,omitempty"`

	// Device Exchange Access State Reason.
	ExchangeAccessStateReason *DeviceManagementExchangeAccessStateReason `json:"exchangeAccessStateReason,omitempty"`

	// Last time the device contacted Exchange. This property is read-only.
	ExchangeLastSuccessfulSyncDateTime *string `json:"exchangeLastSuccessfulSyncDateTime,omitempty"`

	// Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
	FreeStorageSpaceInBytes *int64 `json:"freeStorageSpaceInBytes,omitempty"`

	// Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Default is an empty string. To
	// retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select.
	// $Search is not supported. Read-only. This property is read-only.
	Iccid nullable.Type[string] `json:"iccid,omitempty"`

	// IMEI. This property is read-only.
	Imei nullable.Type[string] `json:"imei,omitempty"`

	// Device encryption status. This property is read-only.
	IsEncrypted *bool `json:"isEncrypted,omitempty"`

	// Device supervised status. This property is read-only.
	IsSupervised *bool `json:"isSupervised,omitempty"`

	// Whether the device is jail broken or rooted. Default is an empty string. Supports $filter operator 'eq' and 'or'.
	// This property is read-only.
	JailBroken nullable.Type[string] `json:"jailBroken,omitempty"`

	// The date and time that the device last completed a successful sync with Intune. Supports $filter operator 'lt' and
	// 'gt'. This property is read-only.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// List of log collection requests
	LogCollectionRequests *[]DeviceLogCollectionResponse `json:"logCollectionRequests,omitempty"`

	// Automatically generated name to identify a device. Can be overwritten to a user friendly name.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// Owner type of device.
	ManagedDeviceOwnerType *ManagedDeviceOwnerType `json:"managedDeviceOwnerType,omitempty"`

	ManagementAgent *ManagementAgentType `json:"managementAgent,omitempty"`

	// Reports device management certificate expiration date. This property is read-only.
	ManagementCertificateExpirationDate *string `json:"managementCertificateExpirationDate,omitempty"`

	// Manufacturer of the device. This property is read-only.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// MEID. This property is read-only.
	Meid nullable.Type[string] `json:"meid,omitempty"`

	// Model of the device. This property is read-only.
	Model nullable.Type[string] `json:"model,omitempty"`

	// Notes on the device created by IT Admin. Default is null. To retrieve actual values GET call needs to be made, with
	// device id and included in select parameter. Supports: $select. $Search is not supported.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// Operating system of the device. Windows, iOS, etc. This property is read-only.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// Operating system version of the device. This property is read-only.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Available health states for the Device Health API
	PartnerReportedThreatState *ManagedDevicePartnerReportedHealthState `json:"partnerReportedThreatState,omitempty"`

	// Phone number of the device. This property is read-only.
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`

	// Total Memory in Bytes. Default is 0. To retrieve actual values GET call needs to be made, with device id and included
	// in select parameter. Supports: $select. Read-only. This property is read-only.
	PhysicalMemoryInBytes *int64 `json:"physicalMemoryInBytes,omitempty"`

	// An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
	RemoteAssistanceSessionErrorDetails nullable.Type[string] `json:"remoteAssistanceSessionErrorDetails,omitempty"`

	// Url that allows a Remote Assistance session to be established with the device. Default is an empty string. To
	// retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is
	// read-only.
	RemoteAssistanceSessionUrl nullable.Type[string] `json:"remoteAssistanceSessionUrl,omitempty"`

	// Reports if the managed iOS device is user approval enrollment. This property is read-only.
	RequireUserEnrollmentApproval nullable.Type[bool] `json:"requireUserEnrollmentApproval,omitempty"`

	// SerialNumber. This property is read-only.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// Subscriber Carrier. This property is read-only.
	SubscriberCarrier nullable.Type[string] `json:"subscriberCarrier,omitempty"`

	// Total Storage in Bytes. This property is read-only.
	TotalStorageSpaceInBytes *int64 `json:"totalStorageSpaceInBytes,omitempty"`

	// Unique Device Identifier for iOS and macOS devices. Default is an empty string. To retrieve actual values GET call
	// needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
	// Read-only. This property is read-only.
	Udid nullable.Type[string] `json:"udid,omitempty"`

	// User display name. This property is read-only.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// Unique Identifier for the user associated with the device. This property is read-only.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Device user principal name. This property is read-only.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The primary users associated with the managed device.
	Users *[]User `json:"users,omitempty"`

	// Wi-Fi MAC. This property is read-only.
	WiFiMacAddress nullable.Type[string] `json:"wiFiMacAddress,omitempty"`

	// The device protection status. This property is read-only.
	WindowsProtectionState *WindowsProtectionState `json:"windowsProtectionState,omitempty"`

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

func (s ManagedDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDevice{}

func (s ManagedDevice) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDevice: %+v", err)
	}

	delete(decoded, "activationLockBypassCode")
	delete(decoded, "androidSecurityPatchLevel")
	delete(decoded, "azureADDeviceId")
	delete(decoded, "azureADRegistered")
	delete(decoded, "complianceGracePeriodExpirationDateTime")
	delete(decoded, "configurationManagerClientEnabledFeatures")
	delete(decoded, "deviceActionResults")
	delete(decoded, "deviceCategoryDisplayName")
	delete(decoded, "deviceHealthAttestationState")
	delete(decoded, "deviceName")
	delete(decoded, "easActivated")
	delete(decoded, "easActivationDateTime")
	delete(decoded, "easDeviceId")
	delete(decoded, "emailAddress")
	delete(decoded, "enrolledDateTime")
	delete(decoded, "enrollmentProfileName")
	delete(decoded, "ethernetMacAddress")
	delete(decoded, "exchangeLastSuccessfulSyncDateTime")
	delete(decoded, "freeStorageSpaceInBytes")
	delete(decoded, "iccid")
	delete(decoded, "imei")
	delete(decoded, "isEncrypted")
	delete(decoded, "isSupervised")
	delete(decoded, "jailBroken")
	delete(decoded, "lastSyncDateTime")
	delete(decoded, "managementCertificateExpirationDate")
	delete(decoded, "manufacturer")
	delete(decoded, "meid")
	delete(decoded, "model")
	delete(decoded, "operatingSystem")
	delete(decoded, "osVersion")
	delete(decoded, "phoneNumber")
	delete(decoded, "physicalMemoryInBytes")
	delete(decoded, "remoteAssistanceSessionErrorDetails")
	delete(decoded, "remoteAssistanceSessionUrl")
	delete(decoded, "requireUserEnrollmentApproval")
	delete(decoded, "serialNumber")
	delete(decoded, "subscriberCarrier")
	delete(decoded, "totalStorageSpaceInBytes")
	delete(decoded, "udid")
	delete(decoded, "userDisplayName")
	delete(decoded, "userId")
	delete(decoded, "userPrincipalName")
	delete(decoded, "wiFiMacAddress")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDevice: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ManagedDevice{}

func (s *ManagedDevice) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActivationLockBypassCode                  nullable.Type[string]                      `json:"activationLockBypassCode,omitempty"`
		AndroidSecurityPatchLevel                 nullable.Type[string]                      `json:"androidSecurityPatchLevel,omitempty"`
		AzureADDeviceId                           nullable.Type[string]                      `json:"azureADDeviceId,omitempty"`
		AzureADRegistered                         nullable.Type[bool]                        `json:"azureADRegistered,omitempty"`
		ComplianceGracePeriodExpirationDateTime   *string                                    `json:"complianceGracePeriodExpirationDateTime,omitempty"`
		ComplianceState                           *ComplianceState                           `json:"complianceState,omitempty"`
		ConfigurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures `json:"configurationManagerClientEnabledFeatures,omitempty"`
		DeviceCategory                            *DeviceCategory                            `json:"deviceCategory,omitempty"`
		DeviceCategoryDisplayName                 nullable.Type[string]                      `json:"deviceCategoryDisplayName,omitempty"`
		DeviceCompliancePolicyStates              *[]DeviceCompliancePolicyState             `json:"deviceCompliancePolicyStates,omitempty"`
		DeviceConfigurationStates                 *[]DeviceConfigurationState                `json:"deviceConfigurationStates,omitempty"`
		DeviceEnrollmentType                      *DeviceEnrollmentType                      `json:"deviceEnrollmentType,omitempty"`
		DeviceHealthAttestationState              *DeviceHealthAttestationState              `json:"deviceHealthAttestationState,omitempty"`
		DeviceName                                nullable.Type[string]                      `json:"deviceName,omitempty"`
		DeviceRegistrationState                   *DeviceRegistrationState                   `json:"deviceRegistrationState,omitempty"`
		EasActivated                              *bool                                      `json:"easActivated,omitempty"`
		EasActivationDateTime                     *string                                    `json:"easActivationDateTime,omitempty"`
		EasDeviceId                               nullable.Type[string]                      `json:"easDeviceId,omitempty"`
		EmailAddress                              nullable.Type[string]                      `json:"emailAddress,omitempty"`
		EnrolledDateTime                          *string                                    `json:"enrolledDateTime,omitempty"`
		EnrollmentProfileName                     nullable.Type[string]                      `json:"enrollmentProfileName,omitempty"`
		EthernetMacAddress                        nullable.Type[string]                      `json:"ethernetMacAddress,omitempty"`
		ExchangeAccessState                       *DeviceManagementExchangeAccessState       `json:"exchangeAccessState,omitempty"`
		ExchangeAccessStateReason                 *DeviceManagementExchangeAccessStateReason `json:"exchangeAccessStateReason,omitempty"`
		ExchangeLastSuccessfulSyncDateTime        *string                                    `json:"exchangeLastSuccessfulSyncDateTime,omitempty"`
		FreeStorageSpaceInBytes                   *int64                                     `json:"freeStorageSpaceInBytes,omitempty"`
		Iccid                                     nullable.Type[string]                      `json:"iccid,omitempty"`
		Imei                                      nullable.Type[string]                      `json:"imei,omitempty"`
		IsEncrypted                               *bool                                      `json:"isEncrypted,omitempty"`
		IsSupervised                              *bool                                      `json:"isSupervised,omitempty"`
		JailBroken                                nullable.Type[string]                      `json:"jailBroken,omitempty"`
		LastSyncDateTime                          *string                                    `json:"lastSyncDateTime,omitempty"`
		LogCollectionRequests                     *[]DeviceLogCollectionResponse             `json:"logCollectionRequests,omitempty"`
		ManagedDeviceName                         nullable.Type[string]                      `json:"managedDeviceName,omitempty"`
		ManagedDeviceOwnerType                    *ManagedDeviceOwnerType                    `json:"managedDeviceOwnerType,omitempty"`
		ManagementAgent                           *ManagementAgentType                       `json:"managementAgent,omitempty"`
		ManagementCertificateExpirationDate       *string                                    `json:"managementCertificateExpirationDate,omitempty"`
		Manufacturer                              nullable.Type[string]                      `json:"manufacturer,omitempty"`
		Meid                                      nullable.Type[string]                      `json:"meid,omitempty"`
		Model                                     nullable.Type[string]                      `json:"model,omitempty"`
		Notes                                     nullable.Type[string]                      `json:"notes,omitempty"`
		OperatingSystem                           nullable.Type[string]                      `json:"operatingSystem,omitempty"`
		OsVersion                                 nullable.Type[string]                      `json:"osVersion,omitempty"`
		PartnerReportedThreatState                *ManagedDevicePartnerReportedHealthState   `json:"partnerReportedThreatState,omitempty"`
		PhoneNumber                               nullable.Type[string]                      `json:"phoneNumber,omitempty"`
		PhysicalMemoryInBytes                     *int64                                     `json:"physicalMemoryInBytes,omitempty"`
		RemoteAssistanceSessionErrorDetails       nullable.Type[string]                      `json:"remoteAssistanceSessionErrorDetails,omitempty"`
		RemoteAssistanceSessionUrl                nullable.Type[string]                      `json:"remoteAssistanceSessionUrl,omitempty"`
		RequireUserEnrollmentApproval             nullable.Type[bool]                        `json:"requireUserEnrollmentApproval,omitempty"`
		SerialNumber                              nullable.Type[string]                      `json:"serialNumber,omitempty"`
		SubscriberCarrier                         nullable.Type[string]                      `json:"subscriberCarrier,omitempty"`
		TotalStorageSpaceInBytes                  *int64                                     `json:"totalStorageSpaceInBytes,omitempty"`
		Udid                                      nullable.Type[string]                      `json:"udid,omitempty"`
		UserDisplayName                           nullable.Type[string]                      `json:"userDisplayName,omitempty"`
		UserId                                    nullable.Type[string]                      `json:"userId,omitempty"`
		UserPrincipalName                         nullable.Type[string]                      `json:"userPrincipalName,omitempty"`
		Users                                     *[]User                                    `json:"users,omitempty"`
		WiFiMacAddress                            nullable.Type[string]                      `json:"wiFiMacAddress,omitempty"`
		WindowsProtectionState                    *WindowsProtectionState                    `json:"windowsProtectionState,omitempty"`
		Id                                        *string                                    `json:"id,omitempty"`
		ODataId                                   *string                                    `json:"@odata.id,omitempty"`
		ODataType                                 *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActivationLockBypassCode = decoded.ActivationLockBypassCode
	s.AndroidSecurityPatchLevel = decoded.AndroidSecurityPatchLevel
	s.AzureADDeviceId = decoded.AzureADDeviceId
	s.AzureADRegistered = decoded.AzureADRegistered
	s.ComplianceGracePeriodExpirationDateTime = decoded.ComplianceGracePeriodExpirationDateTime
	s.ComplianceState = decoded.ComplianceState
	s.ConfigurationManagerClientEnabledFeatures = decoded.ConfigurationManagerClientEnabledFeatures
	s.DeviceCategory = decoded.DeviceCategory
	s.DeviceCategoryDisplayName = decoded.DeviceCategoryDisplayName
	s.DeviceCompliancePolicyStates = decoded.DeviceCompliancePolicyStates
	s.DeviceConfigurationStates = decoded.DeviceConfigurationStates
	s.DeviceEnrollmentType = decoded.DeviceEnrollmentType
	s.DeviceHealthAttestationState = decoded.DeviceHealthAttestationState
	s.DeviceName = decoded.DeviceName
	s.DeviceRegistrationState = decoded.DeviceRegistrationState
	s.EasActivated = decoded.EasActivated
	s.EasActivationDateTime = decoded.EasActivationDateTime
	s.EasDeviceId = decoded.EasDeviceId
	s.EmailAddress = decoded.EmailAddress
	s.EnrolledDateTime = decoded.EnrolledDateTime
	s.EnrollmentProfileName = decoded.EnrollmentProfileName
	s.EthernetMacAddress = decoded.EthernetMacAddress
	s.ExchangeAccessState = decoded.ExchangeAccessState
	s.ExchangeAccessStateReason = decoded.ExchangeAccessStateReason
	s.ExchangeLastSuccessfulSyncDateTime = decoded.ExchangeLastSuccessfulSyncDateTime
	s.FreeStorageSpaceInBytes = decoded.FreeStorageSpaceInBytes
	s.Iccid = decoded.Iccid
	s.Imei = decoded.Imei
	s.IsEncrypted = decoded.IsEncrypted
	s.IsSupervised = decoded.IsSupervised
	s.JailBroken = decoded.JailBroken
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.LogCollectionRequests = decoded.LogCollectionRequests
	s.ManagedDeviceName = decoded.ManagedDeviceName
	s.ManagedDeviceOwnerType = decoded.ManagedDeviceOwnerType
	s.ManagementAgent = decoded.ManagementAgent
	s.ManagementCertificateExpirationDate = decoded.ManagementCertificateExpirationDate
	s.Manufacturer = decoded.Manufacturer
	s.Meid = decoded.Meid
	s.Model = decoded.Model
	s.Notes = decoded.Notes
	s.OperatingSystem = decoded.OperatingSystem
	s.OsVersion = decoded.OsVersion
	s.PartnerReportedThreatState = decoded.PartnerReportedThreatState
	s.PhoneNumber = decoded.PhoneNumber
	s.PhysicalMemoryInBytes = decoded.PhysicalMemoryInBytes
	s.RemoteAssistanceSessionErrorDetails = decoded.RemoteAssistanceSessionErrorDetails
	s.RemoteAssistanceSessionUrl = decoded.RemoteAssistanceSessionUrl
	s.RequireUserEnrollmentApproval = decoded.RequireUserEnrollmentApproval
	s.SerialNumber = decoded.SerialNumber
	s.SubscriberCarrier = decoded.SubscriberCarrier
	s.TotalStorageSpaceInBytes = decoded.TotalStorageSpaceInBytes
	s.Udid = decoded.Udid
	s.UserDisplayName = decoded.UserDisplayName
	s.UserId = decoded.UserId
	s.UserPrincipalName = decoded.UserPrincipalName
	s.Users = decoded.Users
	s.WiFiMacAddress = decoded.WiFiMacAddress
	s.WindowsProtectionState = decoded.WindowsProtectionState
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ManagedDevice into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deviceActionResults"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DeviceActionResults into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceActionResult, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceActionResultImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DeviceActionResults' for 'ManagedDevice': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeviceActionResults = &output
	}

	return nil
}
