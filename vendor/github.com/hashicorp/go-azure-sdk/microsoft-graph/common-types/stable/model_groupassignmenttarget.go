package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAssignmentTarget interface {
	DeviceAndAppManagementAssignmentTarget
	GroupAssignmentTarget() BaseGroupAssignmentTargetImpl
}

var _ GroupAssignmentTarget = BaseGroupAssignmentTargetImpl{}

type BaseGroupAssignmentTargetImpl struct {
	// The group Id that is the target of the assignment.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Fields inherited from DeviceAndAppManagementAssignmentTarget

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseGroupAssignmentTargetImpl) GroupAssignmentTarget() BaseGroupAssignmentTargetImpl {
	return s
}

func (s BaseGroupAssignmentTargetImpl) DeviceAndAppManagementAssignmentTarget() BaseDeviceAndAppManagementAssignmentTargetImpl {
	return BaseDeviceAndAppManagementAssignmentTargetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GroupAssignmentTarget = RawGroupAssignmentTargetImpl{}

// RawGroupAssignmentTargetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGroupAssignmentTargetImpl struct {
	groupAssignmentTarget BaseGroupAssignmentTargetImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawGroupAssignmentTargetImpl) GroupAssignmentTarget() BaseGroupAssignmentTargetImpl {
	return s.groupAssignmentTarget
}

func (s RawGroupAssignmentTargetImpl) DeviceAndAppManagementAssignmentTarget() BaseDeviceAndAppManagementAssignmentTargetImpl {
	return s.groupAssignmentTarget.DeviceAndAppManagementAssignmentTarget()
}

var _ json.Marshaler = BaseGroupAssignmentTargetImpl{}

func (s BaseGroupAssignmentTargetImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGroupAssignmentTargetImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGroupAssignmentTargetImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGroupAssignmentTargetImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupAssignmentTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGroupAssignmentTargetImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalGroupAssignmentTargetImplementation(input []byte) (GroupAssignmentTarget, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupAssignmentTarget into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.exclusionGroupAssignmentTarget") {
		var out ExclusionGroupAssignmentTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExclusionGroupAssignmentTarget: %+v", err)
		}
		return out, nil
	}

	var parent BaseGroupAssignmentTargetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGroupAssignmentTargetImpl: %+v", err)
	}

	return RawGroupAssignmentTargetImpl{
		groupAssignmentTarget: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
