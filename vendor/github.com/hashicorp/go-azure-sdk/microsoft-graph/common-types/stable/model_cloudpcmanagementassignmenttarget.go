package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCManagementAssignmentTarget interface {
	CloudPCManagementAssignmentTarget() BaseCloudPCManagementAssignmentTargetImpl
}

var _ CloudPCManagementAssignmentTarget = BaseCloudPCManagementAssignmentTargetImpl{}

type BaseCloudPCManagementAssignmentTargetImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCloudPCManagementAssignmentTargetImpl) CloudPCManagementAssignmentTarget() BaseCloudPCManagementAssignmentTargetImpl {
	return s
}

var _ CloudPCManagementAssignmentTarget = RawCloudPCManagementAssignmentTargetImpl{}

// RawCloudPCManagementAssignmentTargetImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCloudPCManagementAssignmentTargetImpl struct {
	cloudPCManagementAssignmentTarget BaseCloudPCManagementAssignmentTargetImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawCloudPCManagementAssignmentTargetImpl) CloudPCManagementAssignmentTarget() BaseCloudPCManagementAssignmentTargetImpl {
	return s.cloudPCManagementAssignmentTarget
}

func (s RawCloudPCManagementAssignmentTargetImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCloudPCManagementAssignmentTargetImplementation(input []byte) (CloudPCManagementAssignmentTarget, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCManagementAssignmentTarget into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcManagementGroupAssignmentTarget") {
		var out CloudPCManagementGroupAssignmentTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCManagementGroupAssignmentTarget: %+v", err)
		}
		return out, nil
	}

	var parent BaseCloudPCManagementAssignmentTargetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCloudPCManagementAssignmentTargetImpl: %+v", err)
	}

	return RawCloudPCManagementAssignmentTargetImpl{
		cloudPCManagementAssignmentTarget: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
