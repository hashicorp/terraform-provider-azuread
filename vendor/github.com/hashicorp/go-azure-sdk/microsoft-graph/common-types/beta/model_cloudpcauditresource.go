package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditResource struct {
	// The resource entity display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A list of modified properties.
	ModifiedProperties *[]CloudPCAuditProperty `json:"modifiedProperties,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of the audit resource.
	ResourceId *string `json:"resourceId,omitempty"`

	// The type of the audit resource.
	ResourceType *string `json:"resourceType,omitempty"`
}
