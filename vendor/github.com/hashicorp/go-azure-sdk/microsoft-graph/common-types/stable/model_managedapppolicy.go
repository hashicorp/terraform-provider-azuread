package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPolicy interface {
	Entity
	ManagedAppPolicy() BaseManagedAppPolicyImpl
}

var _ ManagedAppPolicy = BaseManagedAppPolicyImpl{}

type BaseManagedAppPolicyImpl struct {
	// The date and time the policy was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The policy's description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Policy display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Last time the policy was modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s BaseManagedAppPolicyImpl) ManagedAppPolicy() BaseManagedAppPolicyImpl {
	return s
}

func (s BaseManagedAppPolicyImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedAppPolicy = RawManagedAppPolicyImpl{}

// RawManagedAppPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedAppPolicyImpl struct {
	managedAppPolicy BaseManagedAppPolicyImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawManagedAppPolicyImpl) ManagedAppPolicy() BaseManagedAppPolicyImpl {
	return s.managedAppPolicy
}

func (s RawManagedAppPolicyImpl) Entity() BaseEntityImpl {
	return s.managedAppPolicy.Entity()
}

var _ json.Marshaler = BaseManagedAppPolicyImpl{}

func (s BaseManagedAppPolicyImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedAppPolicyImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedAppPolicyImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedAppPolicyImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAppPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedAppPolicyImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalManagedAppPolicyImplementation(input []byte) (ManagedAppPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppConfiguration") {
		var out ManagedAppConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppProtection") {
		var out ManagedAppProtection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppProtection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtection") {
		var out WindowsInformationProtection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtection: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedAppPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedAppPolicyImpl: %+v", err)
	}

	return RawManagedAppPolicyImpl{
		managedAppPolicy: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
