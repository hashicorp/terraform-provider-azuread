package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMappingMetadataEntry struct {
	// Possible values are: EscrowBehavior, DisableMonitoringForChanges, OriginalJoiningProperty, Disposition,
	// IsCustomerDefined, ExcludeFromReporting, Unsynchronized.
	Key *ObjectMappingMetadata `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value of the metadata property.
	Value nullable.Type[string] `json:"value,omitempty"`
}
