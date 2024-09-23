package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessDevices struct {
	// Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties
	// (such as extension attributes) to include/exclude them. Cannot be set if includeDevices or excludeDevices is set.
	DeviceFilter *ConditionalAccessFilter `json:"deviceFilter,omitempty"`

	// States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
	ExcludeDeviceStates *[]string `json:"excludeDeviceStates,omitempty"`

	// States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter
	// is set.
	ExcludeDevices *[]string `json:"excludeDevices,omitempty"`

	// States in the scope of the policy. All is the only allowed value.
	IncludeDeviceStates *[]string `json:"includeDeviceStates,omitempty"`

	// States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFilter is set.
	IncludeDevices *[]string `json:"includeDevices,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
