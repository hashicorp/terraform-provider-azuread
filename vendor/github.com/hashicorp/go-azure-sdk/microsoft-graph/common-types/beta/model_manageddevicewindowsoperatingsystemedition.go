package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceWindowsOperatingSystemEdition struct {
	// Windows Operating System is available in different editions, which have a specific set of features available. This
	// enum type defines the corresponding edition.
	EditionType *ManagedDeviceWindowsOperatingSystemEditionType `json:"editionType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the Date until which this Operating System edition type is officially supported. The Timestamp type
	// represents date and time information using ISO 8601 format and is always in Pacific Time Zone (PT). For example,
	// 2014-01-01 would mean '2014-01-01T07:00:00Z' in UTC time. Returned by default. Read-only.
	SupportEndDate *string `json:"supportEndDate,omitempty"`
}

var _ json.Marshaler = ManagedDeviceWindowsOperatingSystemEdition{}

func (s ManagedDeviceWindowsOperatingSystemEdition) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceWindowsOperatingSystemEdition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceWindowsOperatingSystemEdition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceWindowsOperatingSystemEdition: %+v", err)
	}

	delete(decoded, "supportEndDate")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceWindowsOperatingSystemEdition: %+v", err)
	}

	return encoded, nil
}
