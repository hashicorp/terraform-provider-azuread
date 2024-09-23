package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlStateUpdate struct {
	// Assign the control to the user who will take the action
	AssignedTo nullable.Type[string] `json:"assignedTo,omitempty"`

	// Provides optional comment about the control
	Comment nullable.Type[string] `json:"comment,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// State of the control can be modified using PATCH command(Ex: ignored, thirdParty etc)
	State nullable.Type[string] `json:"state,omitempty"`

	// ID of the user who updated tenant state
	UpdatedBy nullable.Type[string] `json:"updatedBy,omitempty"`

	// Time at which control state was updated
	UpdatedDateTime nullable.Type[string] `json:"updatedDateTime,omitempty"`
}
