package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstitutionData struct {
	// Short description of the institution the user studied at.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the institution the user studied at.
	DisplayName *string `json:"displayName,omitempty"`

	// Address or location of the institute.
	Location *PhysicalAddress `json:"location,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Link to the institution or department homepage.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
