package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinition interface {
	PermissionsDefinition() BasePermissionsDefinitionImpl
}

var _ PermissionsDefinition = BasePermissionsDefinitionImpl{}

type BasePermissionsDefinitionImpl struct {
	AuthorizationSystemInfo *PermissionsDefinitionAuthorizationSystem         `json:"authorizationSystemInfo,omitempty"`
	IdentityInfo            *PermissionsDefinitionAuthorizationSystemIdentity `json:"identityInfo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePermissionsDefinitionImpl) PermissionsDefinition() BasePermissionsDefinitionImpl {
	return s
}

var _ PermissionsDefinition = RawPermissionsDefinitionImpl{}

// RawPermissionsDefinitionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPermissionsDefinitionImpl struct {
	permissionsDefinition BasePermissionsDefinitionImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawPermissionsDefinitionImpl) PermissionsDefinition() BasePermissionsDefinitionImpl {
	return s.permissionsDefinition
}

func UnmarshalPermissionsDefinitionImplementation(input []byte) (PermissionsDefinition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsDefinition into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsPermissionsDefinition") {
		var out AwsPermissionsDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsPermissionsDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.singleResourceAzurePermissionsDefinition") {
		var out SingleResourceAzurePermissionsDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleResourceAzurePermissionsDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.singleResourceGcpPermissionsDefinition") {
		var out SingleResourceGcpPermissionsDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleResourceGcpPermissionsDefinition: %+v", err)
		}
		return out, nil
	}

	var parent BasePermissionsDefinitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePermissionsDefinitionImpl: %+v", err)
	}

	return RawPermissionsDefinitionImpl{
		permissionsDefinition: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
