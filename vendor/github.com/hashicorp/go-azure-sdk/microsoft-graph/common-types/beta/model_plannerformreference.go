package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerFormReference struct {
	// The display name of the form.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The unique identifier of the response.
	FormResponse nullable.Type[string] `json:"formResponse,omitempty"`

	// The URL of the form associated with the plannerFormReference object.
	FormWebUrl nullable.Type[string] `json:"formWebUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
