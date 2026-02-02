package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudPCManagementAssignmentTarget = CloudPCManagementGroupAssignmentTarget{}

type CloudPCManagementGroupAssignmentTarget struct {
	AllotmentDisplayName   nullable.Type[string] `json:"allotmentDisplayName,omitempty"`
	AllotmentLicensesCount nullable.Type[int64]  `json:"allotmentLicensesCount,omitempty"`

	// The ID of the target group for the assignment.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// The unique identifier for the service plan that indicates which size of the Cloud PC to provision for the user. Use a
	// null value, when the provisioningType is dedicated.
	ServicePlanId nullable.Type[string] `json:"servicePlanId,omitempty"`

	// Fields inherited from CloudPCManagementAssignmentTarget

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CloudPCManagementGroupAssignmentTarget) CloudPCManagementAssignmentTarget() BaseCloudPCManagementAssignmentTargetImpl {
	return BaseCloudPCManagementAssignmentTargetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCManagementGroupAssignmentTarget{}

func (s CloudPCManagementGroupAssignmentTarget) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCManagementGroupAssignmentTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCManagementGroupAssignmentTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCManagementGroupAssignmentTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcManagementGroupAssignmentTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCManagementGroupAssignmentTarget: %+v", err)
	}

	return encoded, nil
}
