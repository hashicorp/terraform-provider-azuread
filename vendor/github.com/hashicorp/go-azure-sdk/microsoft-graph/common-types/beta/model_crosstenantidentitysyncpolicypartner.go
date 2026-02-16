package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantIdentitySyncPolicyPartner struct {
	// Display name for the cross-tenant user synchronization policy. Use the name of the partner Microsoft Entra tenant to
	// easily identify the policy. Optional.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	ExternalCloudAuthorizedApplicationId nullable.Type[string] `json:"externalCloudAuthorizedApplicationId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Tenant identifier for the partner Microsoft Entra organization. Read-only.
	TenantId *string `json:"tenantId,omitempty"`

	// Defines whether users can be synchronized from the partner tenant. Key.
	UserSyncInbound *CrossTenantUserSyncInbound `json:"userSyncInbound,omitempty"`
}

var _ json.Marshaler = CrossTenantIdentitySyncPolicyPartner{}

func (s CrossTenantIdentitySyncPolicyPartner) MarshalJSON() ([]byte, error) {
	type wrapper CrossTenantIdentitySyncPolicyPartner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CrossTenantIdentitySyncPolicyPartner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossTenantIdentitySyncPolicyPartner: %+v", err)
	}

	delete(decoded, "tenantId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CrossTenantIdentitySyncPolicyPartner: %+v", err)
	}

	return encoded, nil
}
