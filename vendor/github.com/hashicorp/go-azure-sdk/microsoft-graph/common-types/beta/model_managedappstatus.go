package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppStatus interface {
	Entity
	ManagedAppStatus() BaseManagedAppStatusImpl
}

var _ ManagedAppStatus = BaseManagedAppStatusImpl{}

type BaseManagedAppStatusImpl struct {
	// Friendly name of the status report.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

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

func (s BaseManagedAppStatusImpl) ManagedAppStatus() BaseManagedAppStatusImpl {
	return s
}

func (s BaseManagedAppStatusImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedAppStatus = RawManagedAppStatusImpl{}

// RawManagedAppStatusImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedAppStatusImpl struct {
	managedAppStatus BaseManagedAppStatusImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawManagedAppStatusImpl) ManagedAppStatus() BaseManagedAppStatusImpl {
	return s.managedAppStatus
}

func (s RawManagedAppStatusImpl) Entity() BaseEntityImpl {
	return s.managedAppStatus.Entity()
}

var _ json.Marshaler = BaseManagedAppStatusImpl{}

func (s BaseManagedAppStatusImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedAppStatusImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedAppStatusImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedAppStatusImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAppStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedAppStatusImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalManagedAppStatusImplementation(input []byte) (ManagedAppStatus, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppStatus into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppStatusRaw") {
		var out ManagedAppStatusRaw
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppStatusRaw: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedAppStatusImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedAppStatusImpl: %+v", err)
	}

	return RawManagedAppStatusImpl{
		managedAppStatus: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
