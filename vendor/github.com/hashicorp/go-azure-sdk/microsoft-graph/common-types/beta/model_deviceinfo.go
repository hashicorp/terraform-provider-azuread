package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceInfo struct {
	DeviceId              nullable.Type[string] `json:"deviceId,omitempty"`
	DisplayName           nullable.Type[string] `json:"displayName,omitempty"`
	EnrollmentProfileName nullable.Type[string] `json:"enrollmentProfileName,omitempty"`
	ExtensionAttribute1   nullable.Type[string] `json:"extensionAttribute1,omitempty"`
	ExtensionAttribute10  nullable.Type[string] `json:"extensionAttribute10,omitempty"`
	ExtensionAttribute11  nullable.Type[string] `json:"extensionAttribute11,omitempty"`
	ExtensionAttribute12  nullable.Type[string] `json:"extensionAttribute12,omitempty"`
	ExtensionAttribute13  nullable.Type[string] `json:"extensionAttribute13,omitempty"`
	ExtensionAttribute14  nullable.Type[string] `json:"extensionAttribute14,omitempty"`
	ExtensionAttribute15  nullable.Type[string] `json:"extensionAttribute15,omitempty"`
	ExtensionAttribute2   nullable.Type[string] `json:"extensionAttribute2,omitempty"`
	ExtensionAttribute3   nullable.Type[string] `json:"extensionAttribute3,omitempty"`
	ExtensionAttribute4   nullable.Type[string] `json:"extensionAttribute4,omitempty"`
	ExtensionAttribute5   nullable.Type[string] `json:"extensionAttribute5,omitempty"`
	ExtensionAttribute6   nullable.Type[string] `json:"extensionAttribute6,omitempty"`
	ExtensionAttribute7   nullable.Type[string] `json:"extensionAttribute7,omitempty"`
	ExtensionAttribute8   nullable.Type[string] `json:"extensionAttribute8,omitempty"`
	ExtensionAttribute9   nullable.Type[string] `json:"extensionAttribute9,omitempty"`
	IsCompliant           nullable.Type[bool]   `json:"isCompliant,omitempty"`
	Manufacturer          nullable.Type[string] `json:"manufacturer,omitempty"`
	MdmAppId              nullable.Type[string] `json:"mdmAppId,omitempty"`
	Model                 nullable.Type[string] `json:"model,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	OperatingSystem        nullable.Type[string] `json:"operatingSystem,omitempty"`
	OperatingSystemVersion nullable.Type[string] `json:"operatingSystemVersion,omitempty"`
	Ownership              nullable.Type[string] `json:"ownership,omitempty"`
	PhysicalIds            *[]string             `json:"physicalIds,omitempty"`
	ProfileType            nullable.Type[string] `json:"profileType,omitempty"`
	SystemLabels           *[]string             `json:"systemLabels,omitempty"`
	TrustType              nullable.Type[string] `json:"trustType,omitempty"`
}
