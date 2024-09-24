package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantSecureScore struct {
	// When this Secure Score was created.
	CreateDateTime *string `json:"createDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The maximum historical Secure Score for the tenant.
	TenantMaxScore *int64 `json:"tenantMaxScore,omitempty"`

	// The Secure Score.
	TenantScore *int64 `json:"tenantScore,omitempty"`
}
