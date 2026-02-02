package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditResource struct {
	// Audit resource's type.
	AuditResourceType nullable.Type[string] `json:"auditResourceType,omitempty"`

	// Display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// List of modified properties.
	ModifiedProperties *[]AuditProperty `json:"modifiedProperties,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Audit resource's Id.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Audit resource's type.
	Type nullable.Type[string] `json:"type,omitempty"`
}
