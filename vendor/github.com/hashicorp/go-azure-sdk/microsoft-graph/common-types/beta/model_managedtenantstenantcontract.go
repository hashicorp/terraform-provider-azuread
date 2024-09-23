package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantContract struct {
	// The type of relationship that exists between the managing entity and tenant. Optional. Read-only.
	ContractType nullable.Type[int64] `json:"contractType,omitempty"`

	// The default domain name for the tenant. Required. Read-only.
	DefaultDomainName nullable.Type[string] `json:"defaultDomainName,omitempty"`

	// The display name for the tenant. Optional. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = ManagedTenantsTenantContract{}

func (s ManagedTenantsTenantContract) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTenantContract
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTenantContract: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTenantContract: %+v", err)
	}

	delete(decoded, "contractType")
	delete(decoded, "defaultDomainName")
	delete(decoded, "displayName")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTenantContract: %+v", err)
	}

	return encoded, nil
}
