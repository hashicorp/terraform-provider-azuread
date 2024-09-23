package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DowngradeJustification struct {
	// Indicates whether the downgrade is or isn't justified.
	IsDowngradeJustified *bool `json:"isDowngradeJustified,omitempty"`

	// Message that indicates why a downgrade is justified. The message appears in administrative logs.
	JustificationMessage nullable.Type[string] `json:"justificationMessage,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
