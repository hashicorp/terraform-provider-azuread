package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementConfigurationPolicyAssignment{}

type DeviceManagementConfigurationPolicyAssignment struct {
	// Represents source of assignment.
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`

	// The identifier of the source of the assignment.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// The assignment target for the DeviceManagementConfigurationPolicy.
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

func (s DeviceManagementConfigurationPolicyAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationPolicyAssignment{}

func (s DeviceManagementConfigurationPolicyAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationPolicyAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationPolicyAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationPolicyAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationPolicyAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationPolicyAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationPolicyAssignment{}

func (s *DeviceManagementConfigurationPolicyAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Source    *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`
		SourceId  nullable.Type[string]                   `json:"sourceId,omitempty"`
		Id        *string                                 `json:"id,omitempty"`
		ODataId   *string                                 `json:"@odata.id,omitempty"`
		ODataType *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Source = decoded.Source
	s.SourceId = decoded.SourceId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationPolicyAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'DeviceManagementConfigurationPolicyAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
