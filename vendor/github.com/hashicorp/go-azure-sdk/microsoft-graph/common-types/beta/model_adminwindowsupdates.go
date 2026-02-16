package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AdminWindowsUpdates{}

type AdminWindowsUpdates struct {
	// Catalog of content that can be approved for deployment by Windows Autopatch. Read-only.
	Catalog *WindowsUpdatesCatalog `json:"catalog,omitempty"`

	// The set of updatableAsset resources to which a deployment can apply.
	DeploymentAudiences *[]WindowsUpdatesDeploymentAudience `json:"deploymentAudiences,omitempty"`

	// Deployments created using Windows Autopatch.
	Deployments *[]WindowsUpdatesDeployment `json:"deployments,omitempty"`

	// A collection of Windows products.
	Products *[]WindowsUpdatesProduct `json:"products,omitempty"`

	// Service connections to external resources such as analytics workspaces.
	ResourceConnections *[]WindowsUpdatesResourceConnection `json:"resourceConnections,omitempty"`

	// Assets registered with Windows Autopatch that can receive updates.
	UpdatableAssets *[]WindowsUpdatesUpdatableAsset `json:"updatableAssets,omitempty"`

	// A collection of policies for approving the deployment of different content to an audience over time.
	UpdatePolicies *[]WindowsUpdatesUpdatePolicy `json:"updatePolicies,omitempty"`

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

func (s AdminWindowsUpdates) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AdminWindowsUpdates{}

func (s AdminWindowsUpdates) MarshalJSON() ([]byte, error) {
	type wrapper AdminWindowsUpdates
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AdminWindowsUpdates: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AdminWindowsUpdates: %+v", err)
	}

	delete(decoded, "catalog")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.adminWindowsUpdates"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AdminWindowsUpdates: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AdminWindowsUpdates{}

func (s *AdminWindowsUpdates) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Catalog             *WindowsUpdatesCatalog              `json:"catalog,omitempty"`
		DeploymentAudiences *[]WindowsUpdatesDeploymentAudience `json:"deploymentAudiences,omitempty"`
		Deployments         *[]WindowsUpdatesDeployment         `json:"deployments,omitempty"`
		Products            *[]WindowsUpdatesProduct            `json:"products,omitempty"`
		UpdatePolicies      *[]WindowsUpdatesUpdatePolicy       `json:"updatePolicies,omitempty"`
		Id                  *string                             `json:"id,omitempty"`
		ODataId             *string                             `json:"@odata.id,omitempty"`
		ODataType           *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Catalog = decoded.Catalog
	s.DeploymentAudiences = decoded.DeploymentAudiences
	s.Deployments = decoded.Deployments
	s.Products = decoded.Products
	s.UpdatePolicies = decoded.UpdatePolicies
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AdminWindowsUpdates into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resourceConnections"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ResourceConnections into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesResourceConnection, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesResourceConnectionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ResourceConnections' for 'AdminWindowsUpdates': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ResourceConnections = &output
	}

	if v, ok := temp["updatableAssets"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling UpdatableAssets into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesUpdatableAsset, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesUpdatableAssetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'UpdatableAssets' for 'AdminWindowsUpdates': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.UpdatableAssets = &output
	}

	return nil
}
