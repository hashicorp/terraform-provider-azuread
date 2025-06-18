package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppScope interface {
	Entity
	AppScope() BaseAppScopeImpl
}

var _ AppScope = BaseAppScopeImpl{}

type BaseAppScopeImpl struct {
	// Provides the display name of the app-specific resource represented by the app scope. Read only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Describes the type of app-specific resource represented by the app scope. Read-only.
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

func (s BaseAppScopeImpl) AppScope() BaseAppScopeImpl {
	return s
}

func (s BaseAppScopeImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AppScope = RawAppScopeImpl{}

// RawAppScopeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAppScopeImpl struct {
	appScope BaseAppScopeImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawAppScopeImpl) AppScope() BaseAppScopeImpl {
	return s.appScope
}

func (s RawAppScopeImpl) Entity() BaseEntityImpl {
	return s.appScope.Entity()
}

var _ json.Marshaler = BaseAppScopeImpl{}

func (s BaseAppScopeImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAppScopeImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAppScopeImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAppScopeImpl: %+v", err)
	}

	delete(decoded, "type")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAppScopeImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAppScopeImplementation(input []byte) (AppScope, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AppScope into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customAppScope") {
		var out CustomAppScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomAppScope: %+v", err)
		}
		return out, nil
	}

	var parent BaseAppScopeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAppScopeImpl: %+v", err)
	}

	return RawAppScopeImpl{
		appScope: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
