package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DelegatedAdminCustomer{}

type DelegatedAdminCustomer struct {
	// The Microsoft Entra ID display name of the customer tenant. Read-only. Supports $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Contains the management details of a service in the customer tenant that's managed by delegated administration.
	ServiceManagementDetails *[]DelegatedAdminServiceManagementDetail `json:"serviceManagementDetails,omitempty"`

	// The Microsoft Entra ID-assigned tenant ID of the customer. Read-only.
	TenantId *string `json:"tenantId,omitempty"`

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

func (s DelegatedAdminCustomer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DelegatedAdminCustomer{}

func (s DelegatedAdminCustomer) MarshalJSON() ([]byte, error) {
	type wrapper DelegatedAdminCustomer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DelegatedAdminCustomer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DelegatedAdminCustomer: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.delegatedAdminCustomer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DelegatedAdminCustomer: %+v", err)
	}

	return encoded, nil
}
