package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceInfo struct {
	// Unique identifier set by Azure Device Registration Service at the time of registration.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The display name for the device.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Enrollment profile applied to the device.
	EnrollmentProfileName nullable.Type[string] `json:"enrollmentProfileName,omitempty"`

	// Extension attribute.
	ExtensionAttribute1 nullable.Type[string] `json:"extensionAttribute1,omitempty"`

	// Extension attribute.
	ExtensionAttribute10 nullable.Type[string] `json:"extensionAttribute10,omitempty"`

	// Extension attribute.
	ExtensionAttribute11 nullable.Type[string] `json:"extensionAttribute11,omitempty"`

	// Extension attribute.
	ExtensionAttribute12 nullable.Type[string] `json:"extensionAttribute12,omitempty"`

	// Extension attribute.
	ExtensionAttribute13 nullable.Type[string] `json:"extensionAttribute13,omitempty"`

	// Extension attribute.
	ExtensionAttribute14 nullable.Type[string] `json:"extensionAttribute14,omitempty"`

	// Extension attribute.
	ExtensionAttribute15 nullable.Type[string] `json:"extensionAttribute15,omitempty"`

	// Extension attribute.
	ExtensionAttribute2 nullable.Type[string] `json:"extensionAttribute2,omitempty"`

	// Extension attribute.
	ExtensionAttribute3 nullable.Type[string] `json:"extensionAttribute3,omitempty"`

	// Extension attribute.
	ExtensionAttribute4 nullable.Type[string] `json:"extensionAttribute4,omitempty"`

	// Extension attribute.
	ExtensionAttribute5 nullable.Type[string] `json:"extensionAttribute5,omitempty"`

	// Extension attribute.
	ExtensionAttribute6 nullable.Type[string] `json:"extensionAttribute6,omitempty"`

	// Extension attribute.
	ExtensionAttribute7 nullable.Type[string] `json:"extensionAttribute7,omitempty"`

	// Extension attribute.
	ExtensionAttribute8 nullable.Type[string] `json:"extensionAttribute8,omitempty"`

	// Extension attribute.
	ExtensionAttribute9 nullable.Type[string] `json:"extensionAttribute9,omitempty"`

	// Indicates the device compliance status with Mobile Management Device (MDM) policies. Default is false.
	IsCompliant nullable.Type[bool] `json:"isCompliant,omitempty"`

	// Manufacturer of the device.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Application identifier used to register device into MDM.
	MdmAppId nullable.Type[string] `json:"mdmAppId,omitempty"`

	// Model of the device.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of operating system on the device.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// The version of the operating system on the device.
	OperatingSystemVersion nullable.Type[string] `json:"operatingSystemVersion,omitempty"`

	// Ownership of the device. This property is set by Intune.
	Ownership nullable.Type[string] `json:"ownership,omitempty"`

	// A collection of physical identifiers for the device.
	PhysicalIds *[]string `json:"physicalIds,omitempty"`

	// The profile type of the device.
	ProfileType nullable.Type[string] `json:"profileType,omitempty"`

	// List of labels applied to the device by the system.
	SystemLabels *[]string `json:"systemLabels,omitempty"`

	// Type of trust for the joined device.
	TrustType nullable.Type[string] `json:"trustType,omitempty"`
}
