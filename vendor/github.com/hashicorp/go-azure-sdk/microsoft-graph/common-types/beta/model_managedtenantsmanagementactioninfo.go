package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionInfo struct {
	// The identifier for the management action. Required. Read-only.
	ManagementActionId *string `json:"managementActionId,omitempty"`

	// The identifier for the management template. Required. Read-only.
	ManagementTemplateId *string `json:"managementTemplateId,omitempty"`

	ManagementTemplateVersion *int64 `json:"managementTemplateVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = ManagedTenantsManagementActionInfo{}

func (s ManagedTenantsManagementActionInfo) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementActionInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementActionInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementActionInfo: %+v", err)
	}

	delete(decoded, "managementActionId")
	delete(decoded, "managementTemplateId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementActionInfo: %+v", err)
	}

	return encoded, nil
}
