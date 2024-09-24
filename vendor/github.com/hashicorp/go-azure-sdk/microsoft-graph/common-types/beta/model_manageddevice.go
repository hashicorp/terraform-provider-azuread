package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDevice interface {
	Entity
	ManagedDevice() BaseManagedDeviceImpl
}

var _ ManagedDevice = BaseManagedDeviceImpl{}

type BaseManagedDeviceImpl struct {
	// Whether the device is Azure Active Directory registered. This property is read-only.
	AadRegistered nullable.Type[bool] `json:"aadRegistered,omitempty"`

	// The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property)
	// for this property when returned as part of managedDevice entity in LIST call. To retrieve actual values GET call
	// needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
	// Read-only. This property is read-only.
	ActivationLockBypassCode nullable.Type[string] `json:"activationLockBypassCode,omitempty"`

	// Android security patch level. This property is read-only.
	AndroidSecurityPatchLevel nullable.Type[string] `json:"androidSecurityPatchLevel,omitempty"`

	// Managed device mobile app configuration states for this device.
	AssignmentFilterEvaluationStatusDetails *[]AssignmentFilterEvaluationStatusDetails `json:"assignmentFilterEvaluationStatusDetails,omitempty"`

	// Reports if the managed device is enrolled via auto-pilot. This property is read-only.
	AutopilotEnrolled *bool `json:"autopilotEnrolled,omitempty"`

	// The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
	AzureADDeviceId nullable.Type[string] `json:"azureADDeviceId,omitempty"`

	// Whether the device is Azure Active Directory registered. This property is read-only.
	AzureADRegistered nullable.Type[bool] `json:"azureADRegistered,omitempty"`

	// The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
	AzureActiveDirectoryDeviceId nullable.Type[string] `json:"azureActiveDirectoryDeviceId,omitempty"`

	// Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include
	// BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If
	// TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
	BootstrapTokenEscrowed *bool `json:"bootstrapTokenEscrowed,omitempty"`

	// Chassis type.
	ChassisType *ChassisType `json:"chassisType,omitempty"`

	// List of properties of the ChromeOS Device. Default is an empty list. To retrieve actual values GET call needs to be
	// made, with device id and included in select parameter.
	ChromeOSDeviceInfo *[]ChromeOSDeviceProperty `json:"chromeOSDeviceInfo,omitempty"`

	CloudPCRemoteActionResults *[]CloudPCRemoteActionResult `json:"cloudPcRemoteActionResults,omitempty"`

	// The DateTime when device compliance grace period expires. This property is read-only.
	ComplianceGracePeriodExpirationDateTime *string `json:"complianceGracePeriodExpirationDateTime,omitempty"`

	// Compliance state.
	ComplianceState *ComplianceState `json:"complianceState,omitempty"`

	// ConfigrMgr client enabled features. This property is read-only.
	ConfigurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures `json:"configurationManagerClientEnabledFeatures,omitempty"`

	// Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
	ConfigurationManagerClientHealthState *ConfigurationManagerClientHealthState `json:"configurationManagerClientHealthState,omitempty"`

	// Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr
	// Agent
	ConfigurationManagerClientInformation *ConfigurationManagerClientInformation `json:"configurationManagerClientInformation,omitempty"`

	// All applications currently installed on the device
	DetectedApps *[]DetectedApp `json:"detectedApps,omitempty"`

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

	// Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not
	// DFCI managed. The default value is FALSE.
	DeviceFirmwareConfigurationInterfaceManaged *bool `json:"deviceFirmwareConfigurationInterfaceManaged,omitempty"`

	// The device health attestation state. This property is read-only.
	DeviceHealthAttestationState *DeviceHealthAttestationState `json:"deviceHealthAttestationState,omitempty"`

	// Results of device health scripts that ran for this device. Default is empty list. This property is read-only.
	DeviceHealthScriptStates *[]DeviceHealthScriptPolicyState `json:"deviceHealthScriptStates,omitempty"`

	// Name of the device. This property is read-only.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Device registration status.
	DeviceRegistrationState *DeviceRegistrationState `json:"deviceRegistrationState,omitempty"`

	// Device type.
	DeviceType *DeviceType `json:"deviceType,omitempty"`

	// Whether the device is Exchange ActiveSync activated. This property is read-only.
	EasActivated *bool `json:"easActivated,omitempty"`

	// Exchange ActivationSync activation time of the device. This property is read-only.
	EasActivationDateTime *string `json:"easActivationDateTime,omitempty"`

	// Exchange ActiveSync Id of the device. This property is read-only.
	EasDeviceId nullable.Type[string] `json:"easDeviceId,omitempty"`

	// Email(s) for the user associated with the device. This property is read-only.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// The Entra (Azure AD) User Principal Name (UPN) of the user responsible for the enrollment of the device. This
	// property is read-only.
	EnrolledByUserPrincipalName nullable.Type[string] `json:"enrolledByUserPrincipalName,omitempty"`

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

	// The hardward details for the device. Includes information such as storage space, manufacturer, serial number, etc. By
	// default most property of this type are set to null/0/false and enum defaults for associated types. To retrieve actual
	// values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not
	// supported. Read-only. This property is read-only.
	HardwareInformation *HardwareInformation `json:"hardwareInformation,omitempty"`

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

	// Device enrollment join type.
	JoinType *JoinType `json:"joinType,omitempty"`

	// The date and time that the device last completed a successful sync with Intune. Supports $filter operator 'lt' and
	// 'gt'. This property is read-only.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// List of log collection requests
	LogCollectionRequests *[]DeviceLogCollectionResponse `json:"logCollectionRequests,omitempty"`

	// State of lost mode, indicating if lost mode is enabled or disabled
	LostModeState *LostModeState `json:"lostModeState,omitempty"`

	// Managed device mobile app configuration states for this device.
	ManagedDeviceMobileAppConfigurationStates *[]ManagedDeviceMobileAppConfigurationState `json:"managedDeviceMobileAppConfigurationStates,omitempty"`

	// Automatically generated name to identify a device. Can be overwritten to a user friendly name.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// Owner type of device.
	ManagedDeviceOwnerType *ManagedDeviceOwnerType `json:"managedDeviceOwnerType,omitempty"`

	// Management agent type.
	ManagementAgent *ManagementAgentType `json:"managementAgent,omitempty"`

	// Reports device management certificate expiration date. This property is read-only.
	ManagementCertificateExpirationDate *string `json:"managementCertificateExpirationDate,omitempty"`

	// Device management features.
	ManagementFeatures *ManagedDeviceManagementFeatures `json:"managementFeatures,omitempty"`

	// Management state of device in Microsoft Intune.
	ManagementState *ManagementState `json:"managementState,omitempty"`

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

	// Owner type of device.
	OwnerType *OwnerType `json:"ownerType,omitempty"`

	// Available health states for the Device Health API
	PartnerReportedThreatState *ManagedDevicePartnerReportedHealthState `json:"partnerReportedThreatState,omitempty"`

	// Phone number of the device. This property is read-only.
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`

	// Total Memory in Bytes. Default is 0. To retrieve actual values GET call needs to be made, with device id and included
	// in select parameter. Supports: $select. Read-only. This property is read-only.
	PhysicalMemoryInBytes *int64 `json:"physicalMemoryInBytes,omitempty"`

	// Reports the DateTime the preferMdmOverGroupPolicy setting was set. When set, the Intune MDM settings will override
	// Group Policy settings if there is a conflict. Read Only. This property is read-only.
	PreferMdmOverGroupPolicyAppliedDateTime *string `json:"preferMdmOverGroupPolicyAppliedDateTime,omitempty"`

	// Processor architecture
	ProcessorArchitecture *ManagedDeviceArchitecture `json:"processorArchitecture,omitempty"`

	// An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
	RemoteAssistanceSessionErrorDetails nullable.Type[string] `json:"remoteAssistanceSessionErrorDetails,omitempty"`

	// Url that allows a Remote Assistance session to be established with the device. Default is an empty string. To
	// retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is
	// read-only.
	RemoteAssistanceSessionUrl nullable.Type[string] `json:"remoteAssistanceSessionUrl,omitempty"`

	// Reports if the managed iOS device is user approval enrollment. This property is read-only.
	RequireUserEnrollmentApproval nullable.Type[bool] `json:"requireUserEnrollmentApproval,omitempty"`

	// Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
	RetireAfterDateTime *string `json:"retireAfterDateTime,omitempty"`

	// List of Scope Tag IDs for this Device instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Security baseline states for this device.
	SecurityBaselineStates *[]SecurityBaselineState `json:"securityBaselineStates,omitempty"`

	// This indicates the security patch level of the operating system. These special updates contain important security
	// fixes. For iOS/MacOS they are in (a) format. For android its in 2017-08-07 format. This property is read-only.
	SecurityPatchLevel nullable.Type[string] `json:"securityPatchLevel,omitempty"`

	// SerialNumber. This property is read-only.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// Device sku family
	SkuFamily nullable.Type[string] `json:"skuFamily,omitempty"`

	// Device sku number, see also: https://learn.microsoft.com/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo.
	// Valid values 0 to 2147483647. This property is read-only.
	SkuNumber *int64 `json:"skuNumber,omitempty"`

	// Specification version. This property is read-only.
	SpecificationVersion nullable.Type[string] `json:"specificationVersion,omitempty"`

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

	// Indicates the last logged on users of a device. This property is read-only.
	UsersLoggedOn *[]LoggedOnUser `json:"usersLoggedOn,omitempty"`

	// Wi-Fi MAC. This property is read-only.
	WiFiMacAddress nullable.Type[string] `json:"wiFiMacAddress,omitempty"`

	// Count of active malware for this windows device. Default is 0. To retrieve actual values GET call needs to be made,
	// with device id and included in select parameter. This property is read-only.
	WindowsActiveMalwareCount *int64 `json:"windowsActiveMalwareCount,omitempty"`

	// The device protection status. This property is read-only.
	WindowsProtectionState *WindowsProtectionState `json:"windowsProtectionState,omitempty"`

	// Count of remediated malware for this windows device. Default is 0. To retrieve actual values GET call needs to be
	// made, with device id and included in select parameter. This property is read-only.
	WindowsRemediatedMalwareCount *int64 `json:"windowsRemediatedMalwareCount,omitempty"`

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

func (s BaseManagedDeviceImpl) ManagedDevice() BaseManagedDeviceImpl {
	return s
}

func (s BaseManagedDeviceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedDevice = RawManagedDeviceImpl{}

// RawManagedDeviceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedDeviceImpl struct {
	managedDevice BaseManagedDeviceImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawManagedDeviceImpl) ManagedDevice() BaseManagedDeviceImpl {
	return s.managedDevice
}

func (s RawManagedDeviceImpl) Entity() BaseEntityImpl {
	return s.managedDevice.Entity()
}

var _ json.Marshaler = BaseManagedDeviceImpl{}

func (s BaseManagedDeviceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedDeviceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedDeviceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedDeviceImpl: %+v", err)
	}

	delete(decoded, "aadRegistered")
	delete(decoded, "activationLockBypassCode")
	delete(decoded, "androidSecurityPatchLevel")
	delete(decoded, "autopilotEnrolled")
	delete(decoded, "azureADDeviceId")
	delete(decoded, "azureADRegistered")
	delete(decoded, "azureActiveDirectoryDeviceId")
	delete(decoded, "bootstrapTokenEscrowed")
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
	delete(decoded, "enrolledByUserPrincipalName")
	delete(decoded, "enrolledDateTime")
	delete(decoded, "enrollmentProfileName")
	delete(decoded, "ethernetMacAddress")
	delete(decoded, "exchangeLastSuccessfulSyncDateTime")
	delete(decoded, "freeStorageSpaceInBytes")
	delete(decoded, "hardwareInformation")
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
	delete(decoded, "preferMdmOverGroupPolicyAppliedDateTime")
	delete(decoded, "remoteAssistanceSessionErrorDetails")
	delete(decoded, "remoteAssistanceSessionUrl")
	delete(decoded, "requireUserEnrollmentApproval")
	delete(decoded, "retireAfterDateTime")
	delete(decoded, "securityPatchLevel")
	delete(decoded, "serialNumber")
	delete(decoded, "skuNumber")
	delete(decoded, "specificationVersion")
	delete(decoded, "subscriberCarrier")
	delete(decoded, "totalStorageSpaceInBytes")
	delete(decoded, "udid")
	delete(decoded, "userDisplayName")
	delete(decoded, "userId")
	delete(decoded, "userPrincipalName")
	delete(decoded, "usersLoggedOn")
	delete(decoded, "wiFiMacAddress")
	delete(decoded, "windowsActiveMalwareCount")
	delete(decoded, "windowsRemediatedMalwareCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedDeviceImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseManagedDeviceImpl{}

func (s *BaseManagedDeviceImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AadRegistered                               nullable.Type[bool]                         `json:"aadRegistered,omitempty"`
		ActivationLockBypassCode                    nullable.Type[string]                       `json:"activationLockBypassCode,omitempty"`
		AndroidSecurityPatchLevel                   nullable.Type[string]                       `json:"androidSecurityPatchLevel,omitempty"`
		AssignmentFilterEvaluationStatusDetails     *[]AssignmentFilterEvaluationStatusDetails  `json:"assignmentFilterEvaluationStatusDetails,omitempty"`
		AutopilotEnrolled                           *bool                                       `json:"autopilotEnrolled,omitempty"`
		AzureADDeviceId                             nullable.Type[string]                       `json:"azureADDeviceId,omitempty"`
		AzureADRegistered                           nullable.Type[bool]                         `json:"azureADRegistered,omitempty"`
		AzureActiveDirectoryDeviceId                nullable.Type[string]                       `json:"azureActiveDirectoryDeviceId,omitempty"`
		BootstrapTokenEscrowed                      *bool                                       `json:"bootstrapTokenEscrowed,omitempty"`
		ChassisType                                 *ChassisType                                `json:"chassisType,omitempty"`
		ChromeOSDeviceInfo                          *[]ChromeOSDeviceProperty                   `json:"chromeOSDeviceInfo,omitempty"`
		CloudPCRemoteActionResults                  *[]CloudPCRemoteActionResult                `json:"cloudPcRemoteActionResults,omitempty"`
		ComplianceGracePeriodExpirationDateTime     *string                                     `json:"complianceGracePeriodExpirationDateTime,omitempty"`
		ComplianceState                             *ComplianceState                            `json:"complianceState,omitempty"`
		ConfigurationManagerClientEnabledFeatures   *ConfigurationManagerClientEnabledFeatures  `json:"configurationManagerClientEnabledFeatures,omitempty"`
		ConfigurationManagerClientHealthState       *ConfigurationManagerClientHealthState      `json:"configurationManagerClientHealthState,omitempty"`
		ConfigurationManagerClientInformation       *ConfigurationManagerClientInformation      `json:"configurationManagerClientInformation,omitempty"`
		DetectedApps                                *[]DetectedApp                              `json:"detectedApps,omitempty"`
		DeviceCategory                              *DeviceCategory                             `json:"deviceCategory,omitempty"`
		DeviceCategoryDisplayName                   nullable.Type[string]                       `json:"deviceCategoryDisplayName,omitempty"`
		DeviceCompliancePolicyStates                *[]DeviceCompliancePolicyState              `json:"deviceCompliancePolicyStates,omitempty"`
		DeviceConfigurationStates                   *[]DeviceConfigurationState                 `json:"deviceConfigurationStates,omitempty"`
		DeviceEnrollmentType                        *DeviceEnrollmentType                       `json:"deviceEnrollmentType,omitempty"`
		DeviceFirmwareConfigurationInterfaceManaged *bool                                       `json:"deviceFirmwareConfigurationInterfaceManaged,omitempty"`
		DeviceHealthAttestationState                *DeviceHealthAttestationState               `json:"deviceHealthAttestationState,omitempty"`
		DeviceHealthScriptStates                    *[]DeviceHealthScriptPolicyState            `json:"deviceHealthScriptStates,omitempty"`
		DeviceName                                  nullable.Type[string]                       `json:"deviceName,omitempty"`
		DeviceRegistrationState                     *DeviceRegistrationState                    `json:"deviceRegistrationState,omitempty"`
		DeviceType                                  *DeviceType                                 `json:"deviceType,omitempty"`
		EasActivated                                *bool                                       `json:"easActivated,omitempty"`
		EasActivationDateTime                       *string                                     `json:"easActivationDateTime,omitempty"`
		EasDeviceId                                 nullable.Type[string]                       `json:"easDeviceId,omitempty"`
		EmailAddress                                nullable.Type[string]                       `json:"emailAddress,omitempty"`
		EnrolledByUserPrincipalName                 nullable.Type[string]                       `json:"enrolledByUserPrincipalName,omitempty"`
		EnrolledDateTime                            *string                                     `json:"enrolledDateTime,omitempty"`
		EnrollmentProfileName                       nullable.Type[string]                       `json:"enrollmentProfileName,omitempty"`
		EthernetMacAddress                          nullable.Type[string]                       `json:"ethernetMacAddress,omitempty"`
		ExchangeAccessState                         *DeviceManagementExchangeAccessState        `json:"exchangeAccessState,omitempty"`
		ExchangeAccessStateReason                   *DeviceManagementExchangeAccessStateReason  `json:"exchangeAccessStateReason,omitempty"`
		ExchangeLastSuccessfulSyncDateTime          *string                                     `json:"exchangeLastSuccessfulSyncDateTime,omitempty"`
		FreeStorageSpaceInBytes                     *int64                                      `json:"freeStorageSpaceInBytes,omitempty"`
		HardwareInformation                         *HardwareInformation                        `json:"hardwareInformation,omitempty"`
		Iccid                                       nullable.Type[string]                       `json:"iccid,omitempty"`
		Imei                                        nullable.Type[string]                       `json:"imei,omitempty"`
		IsEncrypted                                 *bool                                       `json:"isEncrypted,omitempty"`
		IsSupervised                                *bool                                       `json:"isSupervised,omitempty"`
		JailBroken                                  nullable.Type[string]                       `json:"jailBroken,omitempty"`
		JoinType                                    *JoinType                                   `json:"joinType,omitempty"`
		LastSyncDateTime                            *string                                     `json:"lastSyncDateTime,omitempty"`
		LogCollectionRequests                       *[]DeviceLogCollectionResponse              `json:"logCollectionRequests,omitempty"`
		LostModeState                               *LostModeState                              `json:"lostModeState,omitempty"`
		ManagedDeviceMobileAppConfigurationStates   *[]ManagedDeviceMobileAppConfigurationState `json:"managedDeviceMobileAppConfigurationStates,omitempty"`
		ManagedDeviceName                           nullable.Type[string]                       `json:"managedDeviceName,omitempty"`
		ManagedDeviceOwnerType                      *ManagedDeviceOwnerType                     `json:"managedDeviceOwnerType,omitempty"`
		ManagementAgent                             *ManagementAgentType                        `json:"managementAgent,omitempty"`
		ManagementCertificateExpirationDate         *string                                     `json:"managementCertificateExpirationDate,omitempty"`
		ManagementFeatures                          *ManagedDeviceManagementFeatures            `json:"managementFeatures,omitempty"`
		ManagementState                             *ManagementState                            `json:"managementState,omitempty"`
		Manufacturer                                nullable.Type[string]                       `json:"manufacturer,omitempty"`
		Meid                                        nullable.Type[string]                       `json:"meid,omitempty"`
		Model                                       nullable.Type[string]                       `json:"model,omitempty"`
		Notes                                       nullable.Type[string]                       `json:"notes,omitempty"`
		OperatingSystem                             nullable.Type[string]                       `json:"operatingSystem,omitempty"`
		OsVersion                                   nullable.Type[string]                       `json:"osVersion,omitempty"`
		OwnerType                                   *OwnerType                                  `json:"ownerType,omitempty"`
		PartnerReportedThreatState                  *ManagedDevicePartnerReportedHealthState    `json:"partnerReportedThreatState,omitempty"`
		PhoneNumber                                 nullable.Type[string]                       `json:"phoneNumber,omitempty"`
		PhysicalMemoryInBytes                       *int64                                      `json:"physicalMemoryInBytes,omitempty"`
		PreferMdmOverGroupPolicyAppliedDateTime     *string                                     `json:"preferMdmOverGroupPolicyAppliedDateTime,omitempty"`
		ProcessorArchitecture                       *ManagedDeviceArchitecture                  `json:"processorArchitecture,omitempty"`
		RemoteAssistanceSessionErrorDetails         nullable.Type[string]                       `json:"remoteAssistanceSessionErrorDetails,omitempty"`
		RemoteAssistanceSessionUrl                  nullable.Type[string]                       `json:"remoteAssistanceSessionUrl,omitempty"`
		RequireUserEnrollmentApproval               nullable.Type[bool]                         `json:"requireUserEnrollmentApproval,omitempty"`
		RetireAfterDateTime                         *string                                     `json:"retireAfterDateTime,omitempty"`
		RoleScopeTagIds                             *[]string                                   `json:"roleScopeTagIds,omitempty"`
		SecurityBaselineStates                      *[]SecurityBaselineState                    `json:"securityBaselineStates,omitempty"`
		SecurityPatchLevel                          nullable.Type[string]                       `json:"securityPatchLevel,omitempty"`
		SerialNumber                                nullable.Type[string]                       `json:"serialNumber,omitempty"`
		SkuFamily                                   nullable.Type[string]                       `json:"skuFamily,omitempty"`
		SkuNumber                                   *int64                                      `json:"skuNumber,omitempty"`
		SpecificationVersion                        nullable.Type[string]                       `json:"specificationVersion,omitempty"`
		SubscriberCarrier                           nullable.Type[string]                       `json:"subscriberCarrier,omitempty"`
		TotalStorageSpaceInBytes                    *int64                                      `json:"totalStorageSpaceInBytes,omitempty"`
		Udid                                        nullable.Type[string]                       `json:"udid,omitempty"`
		UserDisplayName                             nullable.Type[string]                       `json:"userDisplayName,omitempty"`
		UserId                                      nullable.Type[string]                       `json:"userId,omitempty"`
		UserPrincipalName                           nullable.Type[string]                       `json:"userPrincipalName,omitempty"`
		Users                                       *[]User                                     `json:"users,omitempty"`
		UsersLoggedOn                               *[]LoggedOnUser                             `json:"usersLoggedOn,omitempty"`
		WiFiMacAddress                              nullable.Type[string]                       `json:"wiFiMacAddress,omitempty"`
		WindowsActiveMalwareCount                   *int64                                      `json:"windowsActiveMalwareCount,omitempty"`
		WindowsProtectionState                      *WindowsProtectionState                     `json:"windowsProtectionState,omitempty"`
		WindowsRemediatedMalwareCount               *int64                                      `json:"windowsRemediatedMalwareCount,omitempty"`
		Id                                          *string                                     `json:"id,omitempty"`
		ODataId                                     *string                                     `json:"@odata.id,omitempty"`
		ODataType                                   *string                                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AadRegistered = decoded.AadRegistered
	s.ActivationLockBypassCode = decoded.ActivationLockBypassCode
	s.AndroidSecurityPatchLevel = decoded.AndroidSecurityPatchLevel
	s.AssignmentFilterEvaluationStatusDetails = decoded.AssignmentFilterEvaluationStatusDetails
	s.AutopilotEnrolled = decoded.AutopilotEnrolled
	s.AzureADDeviceId = decoded.AzureADDeviceId
	s.AzureADRegistered = decoded.AzureADRegistered
	s.AzureActiveDirectoryDeviceId = decoded.AzureActiveDirectoryDeviceId
	s.BootstrapTokenEscrowed = decoded.BootstrapTokenEscrowed
	s.ChassisType = decoded.ChassisType
	s.ChromeOSDeviceInfo = decoded.ChromeOSDeviceInfo
	s.CloudPCRemoteActionResults = decoded.CloudPCRemoteActionResults
	s.ComplianceGracePeriodExpirationDateTime = decoded.ComplianceGracePeriodExpirationDateTime
	s.ComplianceState = decoded.ComplianceState
	s.ConfigurationManagerClientEnabledFeatures = decoded.ConfigurationManagerClientEnabledFeatures
	s.ConfigurationManagerClientHealthState = decoded.ConfigurationManagerClientHealthState
	s.ConfigurationManagerClientInformation = decoded.ConfigurationManagerClientInformation
	s.DetectedApps = decoded.DetectedApps
	s.DeviceCategory = decoded.DeviceCategory
	s.DeviceCategoryDisplayName = decoded.DeviceCategoryDisplayName
	s.DeviceCompliancePolicyStates = decoded.DeviceCompliancePolicyStates
	s.DeviceConfigurationStates = decoded.DeviceConfigurationStates
	s.DeviceEnrollmentType = decoded.DeviceEnrollmentType
	s.DeviceFirmwareConfigurationInterfaceManaged = decoded.DeviceFirmwareConfigurationInterfaceManaged
	s.DeviceHealthAttestationState = decoded.DeviceHealthAttestationState
	s.DeviceHealthScriptStates = decoded.DeviceHealthScriptStates
	s.DeviceName = decoded.DeviceName
	s.DeviceRegistrationState = decoded.DeviceRegistrationState
	s.DeviceType = decoded.DeviceType
	s.EasActivated = decoded.EasActivated
	s.EasActivationDateTime = decoded.EasActivationDateTime
	s.EasDeviceId = decoded.EasDeviceId
	s.EmailAddress = decoded.EmailAddress
	s.EnrolledByUserPrincipalName = decoded.EnrolledByUserPrincipalName
	s.EnrolledDateTime = decoded.EnrolledDateTime
	s.EnrollmentProfileName = decoded.EnrollmentProfileName
	s.EthernetMacAddress = decoded.EthernetMacAddress
	s.ExchangeAccessState = decoded.ExchangeAccessState
	s.ExchangeAccessStateReason = decoded.ExchangeAccessStateReason
	s.ExchangeLastSuccessfulSyncDateTime = decoded.ExchangeLastSuccessfulSyncDateTime
	s.FreeStorageSpaceInBytes = decoded.FreeStorageSpaceInBytes
	s.HardwareInformation = decoded.HardwareInformation
	s.Iccid = decoded.Iccid
	s.Imei = decoded.Imei
	s.IsEncrypted = decoded.IsEncrypted
	s.IsSupervised = decoded.IsSupervised
	s.JailBroken = decoded.JailBroken
	s.JoinType = decoded.JoinType
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.LogCollectionRequests = decoded.LogCollectionRequests
	s.LostModeState = decoded.LostModeState
	s.ManagedDeviceMobileAppConfigurationStates = decoded.ManagedDeviceMobileAppConfigurationStates
	s.ManagedDeviceName = decoded.ManagedDeviceName
	s.ManagedDeviceOwnerType = decoded.ManagedDeviceOwnerType
	s.ManagementAgent = decoded.ManagementAgent
	s.ManagementCertificateExpirationDate = decoded.ManagementCertificateExpirationDate
	s.ManagementFeatures = decoded.ManagementFeatures
	s.ManagementState = decoded.ManagementState
	s.Manufacturer = decoded.Manufacturer
	s.Meid = decoded.Meid
	s.Model = decoded.Model
	s.Notes = decoded.Notes
	s.OperatingSystem = decoded.OperatingSystem
	s.OsVersion = decoded.OsVersion
	s.OwnerType = decoded.OwnerType
	s.PartnerReportedThreatState = decoded.PartnerReportedThreatState
	s.PhoneNumber = decoded.PhoneNumber
	s.PhysicalMemoryInBytes = decoded.PhysicalMemoryInBytes
	s.PreferMdmOverGroupPolicyAppliedDateTime = decoded.PreferMdmOverGroupPolicyAppliedDateTime
	s.ProcessorArchitecture = decoded.ProcessorArchitecture
	s.RemoteAssistanceSessionErrorDetails = decoded.RemoteAssistanceSessionErrorDetails
	s.RemoteAssistanceSessionUrl = decoded.RemoteAssistanceSessionUrl
	s.RequireUserEnrollmentApproval = decoded.RequireUserEnrollmentApproval
	s.RetireAfterDateTime = decoded.RetireAfterDateTime
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.SecurityBaselineStates = decoded.SecurityBaselineStates
	s.SecurityPatchLevel = decoded.SecurityPatchLevel
	s.SerialNumber = decoded.SerialNumber
	s.SkuFamily = decoded.SkuFamily
	s.SkuNumber = decoded.SkuNumber
	s.SpecificationVersion = decoded.SpecificationVersion
	s.SubscriberCarrier = decoded.SubscriberCarrier
	s.TotalStorageSpaceInBytes = decoded.TotalStorageSpaceInBytes
	s.Udid = decoded.Udid
	s.UserDisplayName = decoded.UserDisplayName
	s.UserId = decoded.UserId
	s.UserPrincipalName = decoded.UserPrincipalName
	s.Users = decoded.Users
	s.UsersLoggedOn = decoded.UsersLoggedOn
	s.WiFiMacAddress = decoded.WiFiMacAddress
	s.WindowsActiveMalwareCount = decoded.WindowsActiveMalwareCount
	s.WindowsProtectionState = decoded.WindowsProtectionState
	s.WindowsRemediatedMalwareCount = decoded.WindowsRemediatedMalwareCount
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseManagedDeviceImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'DeviceActionResults' for 'BaseManagedDeviceImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeviceActionResults = &output
	}

	return nil
}

func UnmarshalManagedDeviceImplementation(input []byte) (ManagedDevice, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDevice into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsManagedDevice") {
		var out WindowsManagedDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsManagedDevice: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedDeviceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedDeviceImpl: %+v", err)
	}

	return RawManagedDeviceImpl{
		managedDevice: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
