package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AndroidDeviceOwnerKioskModeFolderItem = AndroidDeviceOwnerKioskModeWeblink{}

type AndroidDeviceOwnerKioskModeWeblink struct {
	// Display name for weblink
	Label nullable.Type[string] `json:"label,omitempty"`

	// Link for weblink
	Link nullable.Type[string] `json:"link,omitempty"`

	// Fields inherited from AndroidDeviceOwnerKioskModeHomeScreenItem

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AndroidDeviceOwnerKioskModeWeblink) AndroidDeviceOwnerKioskModeFolderItem() BaseAndroidDeviceOwnerKioskModeFolderItemImpl {
	return BaseAndroidDeviceOwnerKioskModeFolderItemImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AndroidDeviceOwnerKioskModeWeblink) AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl {
	return BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerKioskModeWeblink{}

func (s AndroidDeviceOwnerKioskModeWeblink) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerKioskModeWeblink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerKioskModeWeblink: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerKioskModeWeblink: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerKioskModeWeblink"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerKioskModeWeblink: %+v", err)
	}

	return encoded, nil
}
