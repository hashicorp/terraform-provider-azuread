package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MultiTenantOrganizationIdentitySyncPolicyTemplate{}

type MultiTenantOrganizationIdentitySyncPolicyTemplate struct {
	TemplateApplicationLevel *TemplateApplicationLevel `json:"templateApplicationLevel,omitempty"`

	// Defines whether users can be synchronized from the partner tenant.
	UserSyncInbound *CrossTenantUserSyncInbound `json:"userSyncInbound,omitempty"`

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

func (s MultiTenantOrganizationIdentitySyncPolicyTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MultiTenantOrganizationIdentitySyncPolicyTemplate{}

func (s MultiTenantOrganizationIdentitySyncPolicyTemplate) MarshalJSON() ([]byte, error) {
	type wrapper MultiTenantOrganizationIdentitySyncPolicyTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiTenantOrganizationIdentitySyncPolicyTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiTenantOrganizationIdentitySyncPolicyTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.multiTenantOrganizationIdentitySyncPolicyTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiTenantOrganizationIdentitySyncPolicyTemplate: %+v", err)
	}

	return encoded, nil
}
