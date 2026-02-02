package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppConfiguration interface {
	Entity
	ManagedAppPolicy
	ManagedAppConfiguration() BaseManagedAppConfigurationImpl
}

var _ ManagedAppConfiguration = BaseManagedAppConfigurationImpl{}

type BaseManagedAppConfigurationImpl struct {
	// A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped,
	// unalterned by this service
	CustomSettings *[]KeyValuePair `json:"customSettings,omitempty"`

	// Fields inherited from ManagedAppPolicy

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

func (s BaseManagedAppConfigurationImpl) ManagedAppConfiguration() BaseManagedAppConfigurationImpl {
	return s
}

func (s BaseManagedAppConfigurationImpl) ManagedAppPolicy() BaseManagedAppPolicyImpl {
	return BaseManagedAppPolicyImpl{
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s BaseManagedAppConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedAppConfiguration = RawManagedAppConfigurationImpl{}

// RawManagedAppConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedAppConfigurationImpl struct {
	managedAppConfiguration BaseManagedAppConfigurationImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawManagedAppConfigurationImpl) ManagedAppConfiguration() BaseManagedAppConfigurationImpl {
	return s.managedAppConfiguration
}

func (s RawManagedAppConfigurationImpl) ManagedAppPolicy() BaseManagedAppPolicyImpl {
	return s.managedAppConfiguration.ManagedAppPolicy()
}

func (s RawManagedAppConfigurationImpl) Entity() BaseEntityImpl {
	return s.managedAppConfiguration.Entity()
}

var _ json.Marshaler = BaseManagedAppConfigurationImpl{}

func (s BaseManagedAppConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedAppConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedAppConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedAppConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAppConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedAppConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalManagedAppConfigurationImplementation(input []byte) (ManagedAppConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.targetedManagedAppConfiguration") {
		var out TargetedManagedAppConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetedManagedAppConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedAppConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedAppConfigurationImpl: %+v", err)
	}

	return RawManagedAppConfigurationImpl{
		managedAppConfiguration: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
