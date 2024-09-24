package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateDetailedInfo struct {
	Category *ManagedTenantsManagementCategory `json:"category,omitempty"`

	// The display name for the management template. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The unique identifier for the management template. Required. Read-only.
	ManagementTemplateId *string `json:"managementTemplateId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Version *int64 `json:"version,omitempty"`
}

var _ json.Marshaler = ManagedTenantsManagementTemplateDetailedInfo{}

func (s ManagedTenantsManagementTemplateDetailedInfo) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementTemplateDetailedInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementTemplateDetailedInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementTemplateDetailedInfo: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "managementTemplateId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementTemplateDetailedInfo: %+v", err)
	}

	return encoded, nil
}
