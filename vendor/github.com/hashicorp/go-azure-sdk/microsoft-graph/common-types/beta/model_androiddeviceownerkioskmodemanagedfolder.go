package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeManagedFolder struct {
	// Unique identifier for the folder
	FolderIdentifier nullable.Type[string] `json:"folderIdentifier,omitempty"`

	// Display name for the folder
	FolderName *string `json:"folderName,omitempty"`

	// Items to be added to managed folder. This collection can contain a maximum of 500 elements.
	Items *[]AndroidDeviceOwnerKioskModeFolderItem `json:"items,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AndroidDeviceOwnerKioskModeManagedFolder{}

func (s *AndroidDeviceOwnerKioskModeManagedFolder) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FolderIdentifier nullable.Type[string] `json:"folderIdentifier,omitempty"`
		FolderName       *string               `json:"folderName,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FolderIdentifier = decoded.FolderIdentifier
	s.FolderName = decoded.FolderName
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidDeviceOwnerKioskModeManagedFolder into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["items"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Items into list []json.RawMessage: %+v", err)
		}

		output := make([]AndroidDeviceOwnerKioskModeFolderItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAndroidDeviceOwnerKioskModeFolderItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Items' for 'AndroidDeviceOwnerKioskModeManagedFolder': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Items = &output
	}

	return nil
}
