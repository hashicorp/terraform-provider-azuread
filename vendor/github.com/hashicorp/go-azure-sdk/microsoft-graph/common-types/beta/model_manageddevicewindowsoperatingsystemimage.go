package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceWindowsOperatingSystemImage{}

type ManagedDeviceWindowsOperatingSystemImage struct {
	// Indicates the available Quality/Security updates for a specific Windows product version (example: Windows 11 22H1),
	// for upto last 3 Patch Tuesdays . This value in the API response would be updated 2-3 days after every Patch Tuesday.
	// Supports: $filter, $select, $top, $skip. Read-only.
	AvailableUpdates *[]ManagedDeviceWindowsOperatingSystemUpdate `json:"availableUpdates,omitempty"`

	// Indicates the list of architectures supported by the image. E.g. ['ARM64','X86']. Supports: $filter, $select, $top,
	// $skip. Read-only.
	SupportedArchitectures *[]ManagedDeviceArchitecture `json:"supportedArchitectures,omitempty"`

	// Indicates the list of editions supported by the image along with their support dates. Supports: $filter, $select,
	// $top, $skip. Read-only.
	SupportedEditions *[]ManagedDeviceWindowsOperatingSystemEdition `json:"supportedEditions,omitempty"`

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

func (s ManagedDeviceWindowsOperatingSystemImage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceWindowsOperatingSystemImage{}

func (s ManagedDeviceWindowsOperatingSystemImage) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceWindowsOperatingSystemImage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceWindowsOperatingSystemImage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceWindowsOperatingSystemImage: %+v", err)
	}

	delete(decoded, "availableUpdates")
	delete(decoded, "supportedArchitectures")
	delete(decoded, "supportedEditions")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceWindowsOperatingSystemImage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceWindowsOperatingSystemImage: %+v", err)
	}

	return encoded, nil
}
