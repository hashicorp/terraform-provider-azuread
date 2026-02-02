package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceWindowsOperatingSystemUpdate struct {
	// Indicates the build version for associated windows update. Windows Operating System updates are usually released on
	// the Patch Tuesday or B-week of each month. Read-only.
	BuildVersion *string `json:"buildVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the Month in which this B-week update was released. Read-only.
	ReleaseMonth *int64 `json:"releaseMonth,omitempty"`

	// Indicates the Year in which this B-week update was released. Read-only.
	ReleaseYear *int64 `json:"releaseYear,omitempty"`
}

var _ json.Marshaler = ManagedDeviceWindowsOperatingSystemUpdate{}

func (s ManagedDeviceWindowsOperatingSystemUpdate) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceWindowsOperatingSystemUpdate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceWindowsOperatingSystemUpdate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceWindowsOperatingSystemUpdate: %+v", err)
	}

	delete(decoded, "buildVersion")
	delete(decoded, "releaseMonth")
	delete(decoded, "releaseYear")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceWindowsOperatingSystemUpdate: %+v", err)
	}

	return encoded, nil
}
