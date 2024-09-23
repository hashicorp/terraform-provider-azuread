package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DetectedApp{}

type DetectedApp struct {
	// The number of devices that have installed this application
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// Name of the discovered application. Read-only
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The devices that have the discovered application installed
	ManagedDevices *[]ManagedDevice `json:"managedDevices,omitempty"`

	// Indicates the operating system / platform of the discovered application. Some possible values are Windows, iOS,
	// macOS. The default value is unknown (0).
	Platform *DetectedAppPlatformType `json:"platform,omitempty"`

	// Indicates the publisher of the discovered application. For example: 'Microsoft'. The default value is an empty
	// string.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Discovered application size in bytes. Read-only
	SizeInByte *int64 `json:"sizeInByte,omitempty"`

	// Version of the discovered application. Read-only
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s DetectedApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DetectedApp{}

func (s DetectedApp) MarshalJSON() ([]byte, error) {
	type wrapper DetectedApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DetectedApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DetectedApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.detectedApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DetectedApp: %+v", err)
	}

	return encoded, nil
}
