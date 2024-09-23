package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionAuthorizationSystemIdentity struct {
	// Unique ID of the identity within the external system. Prefixed with rsn: if this is a SAML or ED user in AWS.
	// Alternate key.
	ExternalId *string `json:"externalId,omitempty"`

	IdentityType *PermissionsDefinitionIdentityType `json:"identityType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Source PermissionsDefinitionIdentitySource `json:"source"`
}

var _ json.Unmarshaler = &PermissionsDefinitionAuthorizationSystemIdentity{}

func (s *PermissionsDefinitionAuthorizationSystemIdentity) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExternalId   *string                            `json:"externalId,omitempty"`
		IdentityType *PermissionsDefinitionIdentityType `json:"identityType,omitempty"`
		ODataId      *string                            `json:"@odata.id,omitempty"`
		ODataType    *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExternalId = decoded.ExternalId
	s.IdentityType = decoded.IdentityType
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PermissionsDefinitionAuthorizationSystemIdentity into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalPermissionsDefinitionIdentitySourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'PermissionsDefinitionAuthorizationSystemIdentity': %+v", err)
		}
		s.Source = impl
	}

	return nil
}
