package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementIntent{}

type ManagedTenantsManagementIntent struct {
	// The display name for the management intent. Optional. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A flag indicating whether the management intent is global. Required. Read-only.
	IsGlobal *bool `json:"isGlobal,omitempty"`

	// The collection of management templates associated with the management intent. Optional. Read-only.
	ManagementTemplates *[]ManagedTenantsManagementTemplateDetailedInfo `json:"managementTemplates,omitempty"`

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

func (s ManagedTenantsManagementIntent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementIntent{}

func (s ManagedTenantsManagementIntent) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementIntent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementIntent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementIntent: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "isGlobal")
	delete(decoded, "managementTemplates")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementIntent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementIntent: %+v", err)
	}

	return encoded, nil
}
