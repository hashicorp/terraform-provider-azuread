package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GovernanceResource{}

type GovernanceResource struct {
	// The display name of the resource.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The external id of the resource, representing its original id in the external system. For example, a subscription
	// resource's external id can be '/subscriptions/c14ae696-5e0c-4e5d-88cc-bef6637737ac'.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Read-only. The parent resource. for pimforazurerbac scenario, it can represent the subscription the resource belongs
	// to.
	Parent *GovernanceResource `json:"parent,omitempty"`

	// Represents the date time when the resource is registered in PIM.
	RegisteredDateTime nullable.Type[string] `json:"registeredDateTime,omitempty"`

	// The externalId of the resource's root scope that is registered in PIM. The root scope can be the parent, grandparent,
	// or higher ancestor resources.
	RegisteredRoot nullable.Type[string] `json:"registeredRoot,omitempty"`

	// The collection of role assignment requests for the resource.
	RoleAssignmentRequests *[]GovernanceRoleAssignmentRequest `json:"roleAssignmentRequests,omitempty"`

	// The collection of role assignments for the resource.
	RoleAssignments *[]GovernanceRoleAssignment `json:"roleAssignments,omitempty"`

	// The collection of role definitions for the resource.
	RoleDefinitions *[]GovernanceRoleDefinition `json:"roleDefinitions,omitempty"`

	// The collection of role settings for the resource.
	RoleSettings *[]GovernanceRoleSetting `json:"roleSettings,omitempty"`

	// The status of a given resource. For example, it could represent whether the resource is locked or not (values:
	// Active/Locked). Note: This property may be extended in the future to support more scenarios.
	Status nullable.Type[string] `json:"status,omitempty"`

	// Required. Resource type. For example, for Azure resources, the type could be 'Subscription', 'ResourceGroup',
	// 'Microsoft.Sql/server', etc.
	Type nullable.Type[string] `json:"type,omitempty"`

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

func (s GovernanceResource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GovernanceResource{}

func (s GovernanceResource) MarshalJSON() ([]byte, error) {
	type wrapper GovernanceResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GovernanceResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceResource: %+v", err)
	}

	delete(decoded, "parent")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.governanceResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GovernanceResource: %+v", err)
	}

	return encoded, nil
}
