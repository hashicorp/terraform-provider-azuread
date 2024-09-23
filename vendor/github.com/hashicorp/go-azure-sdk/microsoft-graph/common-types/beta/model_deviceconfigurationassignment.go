package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceConfigurationAssignment{}

type DeviceConfigurationAssignment struct {
	// The admin intent to apply or remove the profile. Possible values are: apply, remove.
	Intent *DeviceConfigAssignmentIntent `json:"intent,omitempty"`

	// Represents source of assignment.
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`

	// The identifier of the source of the assignment. This property is read-only.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// The assignment target for the device configuration.
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

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

func (s DeviceConfigurationAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceConfigurationAssignment{}

func (s DeviceConfigurationAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceConfigurationAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceConfigurationAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfigurationAssignment: %+v", err)
	}

	delete(decoded, "sourceId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfigurationAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceConfigurationAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceConfigurationAssignment{}

func (s *DeviceConfigurationAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Intent    *DeviceConfigAssignmentIntent           `json:"intent,omitempty"`
		Source    *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`
		SourceId  nullable.Type[string]                   `json:"sourceId,omitempty"`
		Id        *string                                 `json:"id,omitempty"`
		ODataId   *string                                 `json:"@odata.id,omitempty"`
		ODataType *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Intent = decoded.Intent
	s.Source = decoded.Source
	s.SourceId = decoded.SourceId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceConfigurationAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'DeviceConfigurationAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
