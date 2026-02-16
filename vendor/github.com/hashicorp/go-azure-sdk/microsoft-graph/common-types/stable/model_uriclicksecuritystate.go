package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UriClickSecurityState struct {
	ClickAction   nullable.Type[string] `json:"clickAction,omitempty"`
	ClickDateTime nullable.Type[string] `json:"clickDateTime,omitempty"`
	Id            nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SourceId  nullable.Type[string] `json:"sourceId,omitempty"`
	UriDomain nullable.Type[string] `json:"uriDomain,omitempty"`
	Verdict   nullable.Type[string] `json:"verdict,omitempty"`
}
