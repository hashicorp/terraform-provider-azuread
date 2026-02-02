package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceCompliancePolicyAssignment{}

type DeviceCompliancePolicyAssignment struct {
	// Represents source of assignment.
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`

	// The identifier of the source of the assignment.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// Target for the compliance policy assignment.
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

func (s DeviceCompliancePolicyAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceCompliancePolicyAssignment{}

func (s DeviceCompliancePolicyAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceCompliancePolicyAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceCompliancePolicyAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceCompliancePolicyAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceCompliancePolicyAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceCompliancePolicyAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceCompliancePolicyAssignment{}

func (s *DeviceCompliancePolicyAssignment) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling DeviceCompliancePolicyAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'DeviceCompliancePolicyAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
