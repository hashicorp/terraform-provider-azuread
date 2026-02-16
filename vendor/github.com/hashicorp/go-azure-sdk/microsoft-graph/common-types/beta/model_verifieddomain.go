package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedDomain struct {
	// For example, Email, OfficeCommunicationsOnline.
	Capabilities nullable.Type[string] `json:"capabilities,omitempty"`

	// true if this is the default domain associated with the tenant; otherwise, false.
	IsDefault nullable.Type[bool] `json:"isDefault,omitempty"`

	// true if this is the initial domain associated with the tenant; otherwise, false.
	IsInitial nullable.Type[bool] `json:"isInitial,omitempty"`

	// The domain name; for example, contoso.com.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// For example, Managed.
	Type nullable.Type[string] `json:"type,omitempty"`
}
