package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileSourceAnnotation struct {
	// Indicates whether the source is the default one.
	IsDefaultSource nullable.Type[bool] `json:"isDefaultSource,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Names of properties that have data from this source.
	Properties *[]string `json:"properties,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`
}
