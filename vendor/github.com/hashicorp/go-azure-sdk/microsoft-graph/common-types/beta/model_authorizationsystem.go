package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystem interface {
	Entity
	AuthorizationSystem() BaseAuthorizationSystemImpl
}

var _ AuthorizationSystem = BaseAuthorizationSystemImpl{}

type BaseAuthorizationSystemImpl struct {
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

func (s BaseAuthorizationSystemImpl) AuthorizationSystem() BaseAuthorizationSystemImpl {
	return s
}

func (s BaseAuthorizationSystemImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthorizationSystem = RawAuthorizationSystemImpl{}

// RawAuthorizationSystemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthorizationSystemImpl struct {
	authorizationSystem BaseAuthorizationSystemImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawAuthorizationSystemImpl) AuthorizationSystem() BaseAuthorizationSystemImpl {
	return s.authorizationSystem
}

func (s RawAuthorizationSystemImpl) Entity() BaseEntityImpl {
	return s.authorizationSystem.Entity()
}

var _ json.Marshaler = BaseAuthorizationSystemImpl{}

func (s BaseAuthorizationSystemImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthorizationSystemImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthorizationSystemImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthorizationSystemImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authorizationSystem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthorizationSystemImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthorizationSystemImplementation(input []byte) (AuthorizationSystem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthorizationSystem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsAuthorizationSystem") {
		var out AwsAuthorizationSystem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAuthorizationSystem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAuthorizationSystem") {
		var out AzureAuthorizationSystem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAuthorizationSystem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpAuthorizationSystem") {
		var out GcpAuthorizationSystem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpAuthorizationSystem: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthorizationSystemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthorizationSystemImpl: %+v", err)
	}

	return RawAuthorizationSystemImpl{
		authorizationSystem: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
