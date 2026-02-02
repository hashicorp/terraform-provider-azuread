package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceCompliancePolicyState{}

type DeviceCompliancePolicyState struct {
	// The name of the policy for this policyBase
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Supported platform types for policies.
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`

	// Count of how many setting a policy holds
	SettingCount *int64 `json:"settingCount,omitempty"`

	SettingStates *[]DeviceCompliancePolicySettingState `json:"settingStates,omitempty"`
	State         *ComplianceStatus                     `json:"state,omitempty"`

	// The version of the policy
	Version *int64 `json:"version,omitempty"`

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

func (s DeviceCompliancePolicyState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceCompliancePolicyState{}

func (s DeviceCompliancePolicyState) MarshalJSON() ([]byte, error) {
	type wrapper DeviceCompliancePolicyState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceCompliancePolicyState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceCompliancePolicyState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceCompliancePolicyState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceCompliancePolicyState: %+v", err)
	}

	return encoded, nil
}
