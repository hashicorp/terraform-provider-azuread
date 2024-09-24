package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySource = CrossCloudAzureActiveDirectoryTenant{}

type CrossCloudAzureActiveDirectoryTenant struct {
	// The ID of the cloud where the tenant is located, one of microsoftonline.com, microsoftonline.us or
	// partner.microsoftonline.cn. Read only.
	CloudInstance *string `json:"cloudInstance,omitempty"`

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

func (s CrossCloudAzureActiveDirectoryTenant) IdentitySource() BaseIdentitySourceImpl {
	return BaseIdentitySourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CrossCloudAzureActiveDirectoryTenant{}

func (s CrossCloudAzureActiveDirectoryTenant) MarshalJSON() ([]byte, error) {
	type wrapper CrossCloudAzureActiveDirectoryTenant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CrossCloudAzureActiveDirectoryTenant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossCloudAzureActiveDirectoryTenant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.crossCloudAzureActiveDirectoryTenant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CrossCloudAzureActiveDirectoryTenant: %+v", err)
	}

	return encoded, nil
}
