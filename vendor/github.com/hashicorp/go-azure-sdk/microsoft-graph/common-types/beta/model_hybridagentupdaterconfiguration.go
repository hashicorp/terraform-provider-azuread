package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridAgentUpdaterConfiguration struct {
	// Indicates if updater configuration will be skipped and the agent will receive an update when the next version of the
	// agent is available.
	AllowUpdateConfigurationOverride *bool `json:"allowUpdateConfigurationOverride,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	DeferUpdateDateTime nullable.Type[string] `json:"deferUpdateDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time window during which the agent can receive updates.
	UpdateWindow *UpdateWindow `json:"updateWindow,omitempty"`
}
