package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsTenantTag{}

type ManagedTenantsTenantTag struct {
	// The identifier for the account that created the tenant tag. Required. Read-only.
	CreatedByUserId nullable.Type[string] `json:"createdByUserId,omitempty"`

	// The date and time when the tenant tag was created. Required. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date and time when the tenant tag was deleted. Required. Read-only.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

	// The description for the tenant tag. Optional. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the tenant tag. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identifier for the account that lasted on the tenant tag. Optional. Read-only.
	LastActionByUserId nullable.Type[string] `json:"lastActionByUserId,omitempty"`

	// The date and time the last action was performed against the tenant tag. Optional. Read-only.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// The collection of managed tenants associated with the tenant tag. Optional.
	Tenants *[]ManagedTenantsTenantInfo `json:"tenants,omitempty"`

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

func (s ManagedTenantsTenantTag) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsTenantTag{}

func (s ManagedTenantsTenantTag) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTenantTag
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTenantTag: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTenantTag: %+v", err)
	}

	delete(decoded, "createdByUserId")
	delete(decoded, "createdDateTime")
	delete(decoded, "deletedDateTime")
	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "lastActionByUserId")
	delete(decoded, "lastActionDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.tenantTag"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTenantTag: %+v", err)
	}

	return encoded, nil
}
