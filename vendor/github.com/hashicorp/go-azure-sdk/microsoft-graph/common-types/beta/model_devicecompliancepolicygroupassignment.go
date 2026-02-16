package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceCompliancePolicyGroupAssignment{}

type DeviceCompliancePolicyGroupAssignment struct {
	// The navigation link to the device compliance polic targeted.
	DeviceCompliancePolicy *DeviceCompliancePolicy `json:"deviceCompliancePolicy,omitempty"`

	// Indicates if this group is should be excluded. Defaults that the group should be included
	ExcludeGroup *bool `json:"excludeGroup,omitempty"`

	// The Id of the AAD group we are targeting the device compliance policy to.
	TargetGroupId nullable.Type[string] `json:"targetGroupId,omitempty"`

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

func (s DeviceCompliancePolicyGroupAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceCompliancePolicyGroupAssignment{}

func (s DeviceCompliancePolicyGroupAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceCompliancePolicyGroupAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceCompliancePolicyGroupAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceCompliancePolicyGroupAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceCompliancePolicyGroupAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceCompliancePolicyGroupAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceCompliancePolicyGroupAssignment{}

func (s *DeviceCompliancePolicyGroupAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExcludeGroup  *bool                 `json:"excludeGroup,omitempty"`
		TargetGroupId nullable.Type[string] `json:"targetGroupId,omitempty"`
		Id            *string               `json:"id,omitempty"`
		ODataId       *string               `json:"@odata.id,omitempty"`
		ODataType     *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExcludeGroup = decoded.ExcludeGroup
	s.TargetGroupId = decoded.TargetGroupId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceCompliancePolicyGroupAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deviceCompliancePolicy"]; ok {
		impl, err := UnmarshalDeviceCompliancePolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DeviceCompliancePolicy' for 'DeviceCompliancePolicyGroupAssignment': %+v", err)
		}
		s.DeviceCompliancePolicy = &impl
	}

	return nil
}
