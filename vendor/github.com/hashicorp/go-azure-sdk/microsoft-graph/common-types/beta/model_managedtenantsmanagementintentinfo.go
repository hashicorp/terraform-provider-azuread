package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementIntentInfo struct {
	// The display name for the management intent. Optional. Read-only.
	ManagementIntentDisplayName nullable.Type[string] `json:"managementIntentDisplayName,omitempty"`

	// The identifier for the management intent. Required. Read-only.
	ManagementIntentId *string `json:"managementIntentId,omitempty"`

	// The collection of management template information associated with the management intent. Optional. Read-only.
	ManagementTemplates *[]ManagedTenantsManagementTemplateDetailedInfo `json:"managementTemplates,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = ManagedTenantsManagementIntentInfo{}

func (s ManagedTenantsManagementIntentInfo) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementIntentInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementIntentInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementIntentInfo: %+v", err)
	}

	delete(decoded, "managementIntentDisplayName")
	delete(decoded, "managementIntentId")
	delete(decoded, "managementTemplates")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementIntentInfo: %+v", err)
	}

	return encoded, nil
}
