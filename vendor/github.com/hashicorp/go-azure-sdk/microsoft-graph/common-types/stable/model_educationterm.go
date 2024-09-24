package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationTerm struct {
	// Display name of the term.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// End of the term.
	EndDate nullable.Type[string] `json:"endDate,omitempty"`

	// ID of term in the syncing system.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Start of the term.
	StartDate nullable.Type[string] `json:"startDate,omitempty"`
}
