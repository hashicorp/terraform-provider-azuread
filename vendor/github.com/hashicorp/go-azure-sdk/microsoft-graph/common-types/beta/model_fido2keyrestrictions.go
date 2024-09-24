package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2KeyRestrictions struct {
	// A collection of Authenticator Attestation GUIDs. AADGUIDs define key types and manufacturers.
	AaGuids *[]string `json:"aaGuids,omitempty"`

	// Enforcement type. Possible values are: allow, block.
	EnforcementType *Fido2RestrictionEnforcementType `json:"enforcementType,omitempty"`

	// Determines if the configured key enforcement is enabled.
	IsEnforced nullable.Type[bool] `json:"isEnforced,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
