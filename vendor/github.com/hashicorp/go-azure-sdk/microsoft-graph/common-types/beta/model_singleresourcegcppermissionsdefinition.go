package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PermissionsDefinition = SingleResourceGcpPermissionsDefinition{}

type SingleResourceGcpPermissionsDefinition struct {
	ActionInfo *GcpPermissionsDefinitionAction `json:"actionInfo,omitempty"`

	// Identifier for the resource.
	ResourceId *string `json:"resourceId,omitempty"`

	// Fields inherited from PermissionsDefinition

	AuthorizationSystemInfo *PermissionsDefinitionAuthorizationSystem         `json:"authorizationSystemInfo,omitempty"`
	IdentityInfo            *PermissionsDefinitionAuthorizationSystemIdentity `json:"identityInfo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SingleResourceGcpPermissionsDefinition) PermissionsDefinition() BasePermissionsDefinitionImpl {
	return BasePermissionsDefinitionImpl{
		AuthorizationSystemInfo: s.AuthorizationSystemInfo,
		IdentityInfo:            s.IdentityInfo,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

var _ json.Marshaler = SingleResourceGcpPermissionsDefinition{}

func (s SingleResourceGcpPermissionsDefinition) MarshalJSON() ([]byte, error) {
	type wrapper SingleResourceGcpPermissionsDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SingleResourceGcpPermissionsDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SingleResourceGcpPermissionsDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.singleResourceGcpPermissionsDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SingleResourceGcpPermissionsDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SingleResourceGcpPermissionsDefinition{}

func (s *SingleResourceGcpPermissionsDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ResourceId              *string                                           `json:"resourceId,omitempty"`
		AuthorizationSystemInfo *PermissionsDefinitionAuthorizationSystem         `json:"authorizationSystemInfo,omitempty"`
		IdentityInfo            *PermissionsDefinitionAuthorizationSystemIdentity `json:"identityInfo,omitempty"`
		ODataId                 *string                                           `json:"@odata.id,omitempty"`
		ODataType               *string                                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ResourceId = decoded.ResourceId
	s.AuthorizationSystemInfo = decoded.AuthorizationSystemInfo
	s.IdentityInfo = decoded.IdentityInfo
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SingleResourceGcpPermissionsDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actionInfo"]; ok {
		impl, err := UnmarshalGcpPermissionsDefinitionActionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ActionInfo' for 'SingleResourceGcpPermissionsDefinition': %+v", err)
		}
		s.ActionInfo = &impl
	}

	return nil
}
