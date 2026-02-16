package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
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

var _ json.Unmarshaler = &DetectedApp{}

func (s *DetectedApp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeviceCount *int64                   `json:"deviceCount,omitempty"`
		DisplayName nullable.Type[string]    `json:"displayName,omitempty"`
		Platform    *DetectedAppPlatformType `json:"platform,omitempty"`
		Publisher   nullable.Type[string]    `json:"publisher,omitempty"`
		SizeInByte  *int64                   `json:"sizeInByte,omitempty"`
		Version     nullable.Type[string]    `json:"version,omitempty"`
		Id          *string                  `json:"id,omitempty"`
		ODataId     *string                  `json:"@odata.id,omitempty"`
		ODataType   *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DeviceCount = decoded.DeviceCount
	s.DisplayName = decoded.DisplayName
	s.Platform = decoded.Platform
	s.Publisher = decoded.Publisher
	s.SizeInByte = decoded.SizeInByte
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DetectedApp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["managedDevices"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedDevices into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedDevice, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedDeviceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedDevices' for 'DetectedApp': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedDevices = &output
	}

	return nil
}
