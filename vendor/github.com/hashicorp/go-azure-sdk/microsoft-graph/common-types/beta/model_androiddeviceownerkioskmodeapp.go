package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AndroidDeviceOwnerKioskModeFolderItem = AndroidDeviceOwnerKioskModeApp{}

type AndroidDeviceOwnerKioskModeApp struct {
	// Class name of application
	ClassName nullable.Type[string] `json:"className,omitempty"`

	// Package name of application
	Package *string `json:"package,omitempty"`

	// Fields inherited from AndroidDeviceOwnerKioskModeHomeScreenItem

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AndroidDeviceOwnerKioskModeApp) AndroidDeviceOwnerKioskModeFolderItem() BaseAndroidDeviceOwnerKioskModeFolderItemImpl {
	return BaseAndroidDeviceOwnerKioskModeFolderItemImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AndroidDeviceOwnerKioskModeApp) AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl {
	return BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerKioskModeApp{}

func (s AndroidDeviceOwnerKioskModeApp) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerKioskModeApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerKioskModeApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerKioskModeApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerKioskModeApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerKioskModeApp: %+v", err)
	}

	return encoded, nil
}
