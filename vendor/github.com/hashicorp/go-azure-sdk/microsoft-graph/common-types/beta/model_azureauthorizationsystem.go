package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthorizationSystem = AzureAuthorizationSystem{}

type AzureAuthorizationSystem struct {
	// List of actions for service in authorization system.
	Actions *[]AzureAuthorizationSystemTypeAction `json:"actions,omitempty"`

	// Identities in the authorization system.
	AssociatedIdentities *AzureAssociatedIdentities `json:"associatedIdentities,omitempty"`

	// Resources associated with the authorization system type.
	Resources *[]AzureAuthorizationSystemResource `json:"resources,omitempty"`

	// Roles associated with the authorization system type.
	RoleDefinitions *[]AzureRoleDefinition `json:"roleDefinitions,omitempty"`

	// Services associated with the authorization system type.
	Services *[]AuthorizationSystemTypeService `json:"services,omitempty"`

	// Fields inherited from AuthorizationSystem

	// ID of the authorization system retrieved from the customer cloud environment. Supports $filter(eq, contains) and
	// $orderBy.
	AuthorizationSystemId *string `json:"authorizationSystemId,omitempty"`

	// Name of the authorization system detected after onboarding. Supports $filter(eq,contains) and $orderBy.
	AuthorizationSystemName *string `json:"authorizationSystemName,omitempty"`

	// The type of authorization system. Can be gcp, azure, or aws. Supports $filter(eq).
	AuthorizationSystemType *string `json:"authorizationSystemType,omitempty"`

	// Defines how and whether Permissions Management collects data from the onboarded authorization system. Supports
	// $filter (eq) as follows: $filter=dataCollectionInfo/entitlements/permissionsModificationCapability and
	// $filter=dataCollectionInfo/entitlements/status.
	DataCollectionInfo *DataCollectionInfo `json:"dataCollectionInfo,omitempty"`

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

func (s AzureAuthorizationSystem) AuthorizationSystem() BaseAuthorizationSystemImpl {
	return BaseAuthorizationSystemImpl{
		AuthorizationSystemId:   s.AuthorizationSystemId,
		AuthorizationSystemName: s.AuthorizationSystemName,
		AuthorizationSystemType: s.AuthorizationSystemType,
		DataCollectionInfo:      s.DataCollectionInfo,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s AzureAuthorizationSystem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AzureAuthorizationSystem{}

func (s AzureAuthorizationSystem) MarshalJSON() ([]byte, error) {
	type wrapper AzureAuthorizationSystem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureAuthorizationSystem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureAuthorizationSystem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azureAuthorizationSystem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureAuthorizationSystem: %+v", err)
	}

	return encoded, nil
}
