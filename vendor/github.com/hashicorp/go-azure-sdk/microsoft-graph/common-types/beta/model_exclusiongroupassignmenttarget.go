package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupAssignmentTarget = ExclusionGroupAssignmentTarget{}

type ExclusionGroupAssignmentTarget struct {

	// Fields inherited from GroupAssignmentTarget

	// The group Id that is the target of the assignment.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Fields inherited from DeviceAndAppManagementAssignmentTarget

	// The Id of the filter for the target assignment.
	DeviceAndAppManagementAssignmentFilterId nullable.Type[string] `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`

	// Represents type of the assignment filter.
	DeviceAndAppManagementAssignmentFilterType *DeviceAndAppManagementAssignmentFilterType `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ExclusionGroupAssignmentTarget) GroupAssignmentTarget() BaseGroupAssignmentTargetImpl {
	return BaseGroupAssignmentTargetImpl{
		GroupId:                                  s.GroupId,
		DeviceAndAppManagementAssignmentFilterId: s.DeviceAndAppManagementAssignmentFilterId,
		DeviceAndAppManagementAssignmentFilterType: s.DeviceAndAppManagementAssignmentFilterType,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s ExclusionGroupAssignmentTarget) DeviceAndAppManagementAssignmentTarget() BaseDeviceAndAppManagementAssignmentTargetImpl {
	return BaseDeviceAndAppManagementAssignmentTargetImpl{
		DeviceAndAppManagementAssignmentFilterId:   s.DeviceAndAppManagementAssignmentFilterId,
		DeviceAndAppManagementAssignmentFilterType: s.DeviceAndAppManagementAssignmentFilterType,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExclusionGroupAssignmentTarget{}

func (s ExclusionGroupAssignmentTarget) MarshalJSON() ([]byte, error) {
	type wrapper ExclusionGroupAssignmentTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExclusionGroupAssignmentTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExclusionGroupAssignmentTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.exclusionGroupAssignmentTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExclusionGroupAssignmentTarget: %+v", err)
	}

	return encoded, nil
}
