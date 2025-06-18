package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScopeBase = TenantScope{}

type TenantScope struct {

	// Fields inherited from ScopeBase

	// The identifier for the scope. This could be a user ID, group ID, or a keyword like 'All' for tenant scope.
	Identity nullable.Type[string] `json:"identity,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TenantScope) ScopeBase() BaseScopeBaseImpl {
	return BaseScopeBaseImpl{
		Identity:  s.Identity,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TenantScope{}

func (s TenantScope) MarshalJSON() ([]byte, error) {
	type wrapper TenantScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TenantScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TenantScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tenantScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TenantScope: %+v", err)
	}

	return encoded, nil
}
