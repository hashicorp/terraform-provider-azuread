package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeBase interface {
	ScopeBase() BaseScopeBaseImpl
}

var _ ScopeBase = BaseScopeBaseImpl{}

type BaseScopeBaseImpl struct {
	// The identifier for the scope. This could be a user ID, group ID, or a keyword like 'All' for tenant scope.
	Identity nullable.Type[string] `json:"identity,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseScopeBaseImpl) ScopeBase() BaseScopeBaseImpl {
	return s
}

var _ ScopeBase = RawScopeBaseImpl{}

// RawScopeBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScopeBaseImpl struct {
	scopeBase BaseScopeBaseImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawScopeBaseImpl) ScopeBase() BaseScopeBaseImpl {
	return s.scopeBase
}

func UnmarshalScopeBaseImplementation(input []byte) (ScopeBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScopeBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupScope") {
		var out GroupScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tenantScope") {
		var out TenantScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TenantScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userScope") {
		var out UserScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserScope: %+v", err)
		}
		return out, nil
	}

	var parent BaseScopeBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScopeBaseImpl: %+v", err)
	}

	return RawScopeBaseImpl{
		scopeBase: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
