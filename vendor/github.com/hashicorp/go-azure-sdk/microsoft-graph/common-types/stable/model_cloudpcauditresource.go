package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditResource struct {
	// The display name of the modified resource entity.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The list of modified properties.
	ModifiedProperties *[]CloudPCAuditProperty `json:"modifiedProperties,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier of the modified resource entity.
	ResourceId *string `json:"resourceId,omitempty"`
}
