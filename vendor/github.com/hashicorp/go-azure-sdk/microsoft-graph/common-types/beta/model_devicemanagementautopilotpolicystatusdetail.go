package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementAutopilotPolicyStatusDetail{}

type DeviceManagementAutopilotPolicyStatusDetail struct {
	ComplianceStatus *DeviceManagementAutopilotPolicyComplianceStatus `json:"complianceStatus,omitempty"`

	// The friendly name of the policy.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The errorode associated with the compliance or enforcement status of the policy. Error code for enforcement status
	// takes precedence if it exists.
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// Timestamp of the reported policy status
	LastReportedDateTime *string `json:"lastReportedDateTime,omitempty"`

	PolicyType *DeviceManagementAutopilotPolicyType `json:"policyType,omitempty"`

	// Indicates if this policy was tracked as part of the autopilot bootstrap enrollment sync session
	TrackedOnEnrollmentStatus *bool `json:"trackedOnEnrollmentStatus,omitempty"`

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

func (s DeviceManagementAutopilotPolicyStatusDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementAutopilotPolicyStatusDetail{}

func (s DeviceManagementAutopilotPolicyStatusDetail) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementAutopilotPolicyStatusDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementAutopilotPolicyStatusDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementAutopilotPolicyStatusDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementAutopilotPolicyStatusDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementAutopilotPolicyStatusDetail: %+v", err)
	}

	return encoded, nil
}
