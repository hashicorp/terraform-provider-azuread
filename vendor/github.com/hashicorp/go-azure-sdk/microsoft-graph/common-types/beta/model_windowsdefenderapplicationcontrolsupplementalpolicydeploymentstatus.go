package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus{}

type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus struct {
	// Enum values for the various WindowsDefenderApplicationControl supplemental policy deployment statuses.
	DeploymentStatus *WindowsDefenderApplicationControlSupplementalPolicyStatuses `json:"deploymentStatus,omitempty"`

	// Device ID.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Device name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Last sync date time.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// Windows OS Version Description.
	OsDescription nullable.Type[string] `json:"osDescription,omitempty"`

	// Windows OS Version.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// The navigation link to the WindowsDefenderApplicationControl supplemental policy.
	Policy *WindowsDefenderApplicationControlSupplementalPolicy `json:"policy,omitempty"`

	// Human readable version of the WindowsDefenderApplicationControl supplemental policy.
	PolicyVersion nullable.Type[string] `json:"policyVersion,omitempty"`

	// The name of the user of this device.
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// User Principal Name.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus{}

func (s WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus: %+v", err)
	}

	return encoded, nil
}
