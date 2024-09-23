package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedAppPolicyDeploymentSummary{}

type ManagedAppPolicyDeploymentSummary struct {
	ConfigurationDeployedUserCount       *int64                                     `json:"configurationDeployedUserCount,omitempty"`
	ConfigurationDeploymentSummaryPerApp *[]ManagedAppPolicyDeploymentSummaryPerApp `json:"configurationDeploymentSummaryPerApp,omitempty"`
	DisplayName                          nullable.Type[string]                      `json:"displayName,omitempty"`
	LastRefreshTime                      *string                                    `json:"lastRefreshTime,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s ManagedAppPolicyDeploymentSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedAppPolicyDeploymentSummary{}

func (s ManagedAppPolicyDeploymentSummary) MarshalJSON() ([]byte, error) {
	type wrapper ManagedAppPolicyDeploymentSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedAppPolicyDeploymentSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppPolicyDeploymentSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAppPolicyDeploymentSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedAppPolicyDeploymentSummary: %+v", err)
	}

	return encoded, nil
}
