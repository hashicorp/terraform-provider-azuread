package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionIdentitySource interface {
	PermissionsDefinitionIdentitySource() BasePermissionsDefinitionIdentitySourceImpl
}

var _ PermissionsDefinitionIdentitySource = BasePermissionsDefinitionIdentitySourceImpl{}

type BasePermissionsDefinitionIdentitySourceImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePermissionsDefinitionIdentitySourceImpl) PermissionsDefinitionIdentitySource() BasePermissionsDefinitionIdentitySourceImpl {
	return s
}

var _ PermissionsDefinitionIdentitySource = RawPermissionsDefinitionIdentitySourceImpl{}

// RawPermissionsDefinitionIdentitySourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPermissionsDefinitionIdentitySourceImpl struct {
	permissionsDefinitionIdentitySource BasePermissionsDefinitionIdentitySourceImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawPermissionsDefinitionIdentitySourceImpl) PermissionsDefinitionIdentitySource() BasePermissionsDefinitionIdentitySourceImpl {
	return s.permissionsDefinitionIdentitySource
}

func UnmarshalPermissionsDefinitionIdentitySourceImplementation(input []byte) (PermissionsDefinitionIdentitySource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsDefinitionIdentitySource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsIdentitySource") {
		var out AwsIdentitySource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsIdentitySource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.edIdentitySource") {
		var out EdIdentitySource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdIdentitySource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.localIdentitySource") {
		var out LocalIdentitySource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocalIdentitySource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.samlIdentitySource") {
		var out SamlIdentitySource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SamlIdentitySource: %+v", err)
		}
		return out, nil
	}

	var parent BasePermissionsDefinitionIdentitySourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePermissionsDefinitionIdentitySourceImpl: %+v", err)
	}

	return RawPermissionsDefinitionIdentitySourceImpl{
		permissionsDefinitionIdentitySource: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
