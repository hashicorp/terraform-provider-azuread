package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary{}

type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary struct {
	// Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
	DeployedDeviceCount *int64 `json:"deployedDeviceCount,omitempty"`

	// Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

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

func (s WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary{}

func (s WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyDeploymentSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary: %+v", err)
	}

	return encoded, nil
}
