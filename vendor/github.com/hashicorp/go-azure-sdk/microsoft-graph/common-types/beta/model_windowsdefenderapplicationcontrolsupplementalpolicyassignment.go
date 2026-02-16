package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsDefenderApplicationControlSupplementalPolicyAssignment{}

type WindowsDefenderApplicationControlSupplementalPolicyAssignment struct {
	// The target group assignment defined by the admin.
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

func (s WindowsDefenderApplicationControlSupplementalPolicyAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDefenderApplicationControlSupplementalPolicyAssignment{}

func (s WindowsDefenderApplicationControlSupplementalPolicyAssignment) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDefenderApplicationControlSupplementalPolicyAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDefenderApplicationControlSupplementalPolicyAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDefenderApplicationControlSupplementalPolicyAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDefenderApplicationControlSupplementalPolicyAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsDefenderApplicationControlSupplementalPolicyAssignment{}

func (s *WindowsDefenderApplicationControlSupplementalPolicyAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsDefenderApplicationControlSupplementalPolicyAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'WindowsDefenderApplicationControlSupplementalPolicyAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
