package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesUpdatableAsset = WindowsUpdatesAzureADDevice{}

type WindowsUpdatesAzureADDevice struct {
	Enrollment *WindowsUpdatesUpdateManagementEnrollment `json:"enrollment,omitempty"`

	// Specifies any errors that prevent the device from being enrolled in update management or receving deployed content.
	// Read-only. Returned by default.
	Errors *[]WindowsUpdatesUpdatableAssetError `json:"errors,omitempty"`

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

func (s WindowsUpdatesAzureADDevice) WindowsUpdatesUpdatableAsset() BaseWindowsUpdatesUpdatableAssetImpl {
	return BaseWindowsUpdatesUpdatableAssetImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesAzureADDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesAzureADDevice{}

func (s WindowsUpdatesAzureADDevice) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesAzureADDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesAzureADDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesAzureADDevice: %+v", err)
	}

	delete(decoded, "errors")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.azureADDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesAzureADDevice: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesAzureADDevice{}

func (s *WindowsUpdatesAzureADDevice) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Enrollment *WindowsUpdatesUpdateManagementEnrollment `json:"enrollment,omitempty"`
		Id         *string                                   `json:"id,omitempty"`
		ODataId    *string                                   `json:"@odata.id,omitempty"`
		ODataType  *string                                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Enrollment = decoded.Enrollment
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesAzureADDevice into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["errors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Errors into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesUpdatableAssetError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesUpdatableAssetErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Errors' for 'WindowsUpdatesAzureADDevice': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Errors = &output
	}

	return nil
}
