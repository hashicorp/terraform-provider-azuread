package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsMyRole struct {
	// A collection of role assignments for the managed tenant.
	Assignments *[]ManagedTenantsRoleAssignment `json:"assignments,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
	TenantId *string `json:"tenantId,omitempty"`
}

var _ json.Marshaler = ManagedTenantsMyRole{}

func (s ManagedTenantsMyRole) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsMyRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsMyRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsMyRole: %+v", err)
	}

	delete(decoded, "tenantId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsMyRole: %+v", err)
	}

	return encoded, nil
}
