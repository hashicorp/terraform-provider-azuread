package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonNamePronounciation struct {
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`
	First       nullable.Type[string] `json:"first,omitempty"`
	Last        nullable.Type[string] `json:"last,omitempty"`
	Maiden      nullable.Type[string] `json:"maiden,omitempty"`
	Middle      nullable.Type[string] `json:"middle,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
