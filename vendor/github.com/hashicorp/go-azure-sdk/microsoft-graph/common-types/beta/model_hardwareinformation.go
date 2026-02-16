package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformation struct {
	// The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
	BatteryChargeCycles *int64 `json:"batteryChargeCycles,omitempty"`

	// The device’s current battery’s health percentage. Valid values 0 to 100
	BatteryHealthPercentage *int64 `json:"batteryHealthPercentage,omitempty"`

	// The serial number of the device’s current battery
	BatterySerialNumber nullable.Type[string] `json:"batterySerialNumber,omitempty"`

	// Cellular technology of the device
	CellularTechnology nullable.Type[string] `json:"cellularTechnology,omitempty"`

	// Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an
	// empty string.
	DeviceFullQualifiedDomainName nullable.Type[string] `json:"deviceFullQualifiedDomainName,omitempty"`

	DeviceGuardLocalSystemAuthorityCredentialGuardState            *DeviceGuardLocalSystemAuthorityCredentialGuardState            `json:"deviceGuardLocalSystemAuthorityCredentialGuardState,omitempty"`
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState `json:"deviceGuardVirtualizationBasedSecurityHardwareRequirementState,omitempty"`
	DeviceGuardVirtualizationBasedSecurityState                    *DeviceGuardVirtualizationBasedSecurityState                    `json:"deviceGuardVirtualizationBasedSecurityState,omitempty"`

	// A standard error code indicating the last error, or 0 indicating no error (default). The update frequency of this
	// property is daily. Note this property is currently supported only for Windows based Device based subscription
	// licensing. Valid values 0 to 2147483647
	DeviceLicensingLastErrorCode *int64 `json:"deviceLicensingLastErrorCode,omitempty"`

	// Error text message as a descripition for deviceLicensingLastErrorCode. The update frequency of this property is
	// daily. Note this property is currently supported only for Windows based Device based subscription licensing.
	DeviceLicensingLastErrorDescription nullable.Type[string] `json:"deviceLicensingLastErrorDescription,omitempty"`

	// Indicates the device licensing status after Windows device based subscription has been enabled.
	DeviceLicensingStatus *DeviceLicensingStatus `json:"deviceLicensingStatus,omitempty"`

	// eSIM identifier
	EsimIdentifier nullable.Type[string] `json:"esimIdentifier,omitempty"`

	// Free storage space of the device.
	FreeStorageSpace *int64 `json:"freeStorageSpace,omitempty"`

	// IPAddressV4
	IPAddressV4 nullable.Type[string] `json:"ipAddressV4,omitempty"`

	// IMEI
	Imei nullable.Type[string] `json:"imei,omitempty"`

	// Encryption status of the device
	IsEncrypted *bool `json:"isEncrypted,omitempty"`

	// Shared iPad
	IsSharedDevice *bool `json:"isSharedDevice,omitempty"`

	// Supervised mode of the device
	IsSupervised *bool `json:"isSupervised,omitempty"`

	// Manufacturer of the device
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// MEID
	Meid nullable.Type[string] `json:"meid,omitempty"`

	// Model of the device
	Model nullable.Type[string] `json:"model,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// String that specifies the OS edition.
	OperatingSystemEdition nullable.Type[string] `json:"operatingSystemEdition,omitempty"`

	// Operating system language of the device
	OperatingSystemLanguage nullable.Type[string] `json:"operatingSystemLanguage,omitempty"`

	// Int that specifies the Windows Operating System ProductType. More details here
	// https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
	OperatingSystemProductType *int64 `json:"operatingSystemProductType,omitempty"`

	// Operating System Build Number on Android device
	OsBuildNumber nullable.Type[string] `json:"osBuildNumber,omitempty"`

	// Phone number of the device
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`

	// The product name, e.g. iPad8,12 etc. The update frequency of this property is weekly. Note this property is currently
	// supported only on iOS/MacOS devices, and is available only when Device Information access right is obtained.
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The number of users currently on this device, or null (default) if the value of this property cannot be determined.
	// The update frequency of this property is per-checkin. Note this property is currently supported only on devices
	// running iOS 13.4 and later, and is available only when Device Information access right is obtained. Valid values 0 to
	// 2147483647
	ResidentUsersCount nullable.Type[int64] `json:"residentUsersCount,omitempty"`

	// Serial number.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// All users on the shared Apple device
	SharedDeviceCachedUsers *[]SharedAppleDeviceUser `json:"sharedDeviceCachedUsers,omitempty"`

	// SubnetAddress
	SubnetAddress nullable.Type[string] `json:"subnetAddress,omitempty"`

	// Subscriber carrier of the device
	SubscriberCarrier nullable.Type[string] `json:"subscriberCarrier,omitempty"`

	// BIOS version as reported by SMBIOS
	SystemManagementBIOSVersion nullable.Type[string] `json:"systemManagementBIOSVersion,omitempty"`

	// Total storage space of the device.
	TotalStorageSpace *int64 `json:"totalStorageSpace,omitempty"`

	// The identifying information that uniquely names the TPM manufacturer
	TpmManufacturer nullable.Type[string] `json:"tpmManufacturer,omitempty"`

	// String that specifies the specification version.
	TpmSpecificationVersion nullable.Type[string] `json:"tpmSpecificationVersion,omitempty"`

	// The version of the TPM, as specified by the manufacturer
	TpmVersion nullable.Type[string] `json:"tpmVersion,omitempty"`

	// WiFi MAC address of the device
	WifiMac nullable.Type[string] `json:"wifiMac,omitempty"`

	// A list of wired IPv4 addresses. The update frequency (the maximum delay for the change of property value to be
	// synchronized from the device to the cloud storage) of this property is daily. Note this property is currently
	// supported only on devices running on Windows.
	WiredIPv4Addresses *[]string `json:"wiredIPv4Addresses,omitempty"`
}
