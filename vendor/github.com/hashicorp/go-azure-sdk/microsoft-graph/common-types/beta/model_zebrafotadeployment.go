package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ZebraFotaDeployment{}

type ZebraFotaDeployment struct {
	// Collection of Android FOTA Assignment
	DeploymentAssignments *[]AndroidFotaDeploymentAssignment `json:"deploymentAssignments,omitempty"`

	// The Zebra FOTA deployment complex type that describes the settings required to create a FOTA deployment.
	DeploymentSettings *ZebraFotaDeploymentSettings `json:"deploymentSettings,omitempty"`

	// Represents the deployment status from Zebra. The status is a high level status of the deployment as opposed being a
	// detailed status per device.
	DeploymentStatus *ZebraFotaDeploymentStatus `json:"deploymentStatus,omitempty"`

	// A human readable description of the deployment.
	Description nullable.Type[string] `json:"description,omitempty"`

	// A human readable name of the deployment.
	DisplayName *string `json:"displayName,omitempty"`

	// List of Scope Tags for this Entity instance
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s ZebraFotaDeployment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ZebraFotaDeployment{}

func (s ZebraFotaDeployment) MarshalJSON() ([]byte, error) {
	type wrapper ZebraFotaDeployment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ZebraFotaDeployment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ZebraFotaDeployment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.zebraFotaDeployment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ZebraFotaDeployment: %+v", err)
	}

	return encoded, nil
}
