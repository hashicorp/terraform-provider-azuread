package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySource = AzureActiveDirectoryTenant{}

type AzureActiveDirectoryTenant struct {
	// The name of the Microsoft Entra tenant. Read only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The ID of the Microsoft Entra tenant. Read only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// Fields inherited from IdentitySource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AzureActiveDirectoryTenant) IdentitySource() BaseIdentitySourceImpl {
	return BaseIdentitySourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AzureActiveDirectoryTenant{}

func (s AzureActiveDirectoryTenant) MarshalJSON() ([]byte, error) {
	type wrapper AzureActiveDirectoryTenant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureActiveDirectoryTenant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureActiveDirectoryTenant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azureActiveDirectoryTenant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureActiveDirectoryTenant: %+v", err)
	}

	return encoded, nil
}
