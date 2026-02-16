package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsInformationProtectionApp = WindowsInformationProtectionStoreApp{}

type WindowsInformationProtectionStoreApp struct {

	// Fields inherited from WindowsInformationProtectionApp

	// If true, app is denied protection or exemption.
	Denied *bool `json:"denied,omitempty"`

	// The app's description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// App display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The product name.
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The publisher name
	PublisherName nullable.Type[string] `json:"publisherName,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsInformationProtectionStoreApp) WindowsInformationProtectionApp() BaseWindowsInformationProtectionAppImpl {
	return BaseWindowsInformationProtectionAppImpl{
		Denied:        s.Denied,
		Description:   s.Description,
		DisplayName:   s.DisplayName,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
		ProductName:   s.ProductName,
		PublisherName: s.PublisherName,
	}
}

var _ json.Marshaler = WindowsInformationProtectionStoreApp{}

func (s WindowsInformationProtectionStoreApp) MarshalJSON() ([]byte, error) {
	type wrapper WindowsInformationProtectionStoreApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsInformationProtectionStoreApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsInformationProtectionStoreApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsInformationProtectionStoreApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsInformationProtectionStoreApp: %+v", err)
	}

	return encoded, nil
}
