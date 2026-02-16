package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemInfo struct {
	// The type of authorization system.The possible values are: azure, gcp, aws, unknownFutureValue.
	AuthorizationSystemType *AuthorizationSystemType `json:"authorizationSystemType,omitempty"`

	// Display name for the authorization system.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique identifier for the authorization system.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
