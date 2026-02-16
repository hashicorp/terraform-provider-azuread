package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetrievalHit struct {
	Extracts *[]RetrievalExtract `json:"extracts,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ResourceMetadata *SearchResourceMetadataDictionary `json:"resourceMetadata,omitempty"`
	ResourceType     *RetrievalEntityType              `json:"resourceType,omitempty"`
	SensitivityLabel *SearchSensitivityLabelInfo       `json:"sensitivityLabel,omitempty"`
	WebUrl           nullable.Type[string]             `json:"webUrl,omitempty"`
}
