package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthorizationSystem = AwsAuthorizationSystem{}

type AwsAuthorizationSystem struct {
	// List of actions for service in authorization system.
	Actions *[]AwsAuthorizationSystemTypeAction `json:"actions,omitempty"`

	// Identities in the authorization system.
	AssociatedIdentities *AwsAssociatedIdentities `json:"associatedIdentities,omitempty"`

	// Policies associated with the AWS authorization system type.
	Policies *[]AwsPolicy `json:"policies,omitempty"`

	// Resources associated with the authorization system type.
	Resources *[]AwsAuthorizationSystemResource `json:"resources,omitempty"`

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

func (s AwsAuthorizationSystem) AuthorizationSystem() BaseAuthorizationSystemImpl {
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

func (s AwsAuthorizationSystem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsAuthorizationSystem{}

func (s AwsAuthorizationSystem) MarshalJSON() ([]byte, error) {
	type wrapper AwsAuthorizationSystem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsAuthorizationSystem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsAuthorizationSystem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsAuthorizationSystem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsAuthorizationSystem: %+v", err)
	}

	return encoded, nil
}
