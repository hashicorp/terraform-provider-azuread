package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CampaignSchedule struct {
	// The date and time at which the campaign completed.
	CompletionDateTime nullable.Type[string] `json:"completionDateTime,omitempty"`

	// The date and time at which the campaign was launched.
	LaunchDateTime nullable.Type[string] `json:"launchDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The current state of the campaign. The possible values are: unknown, draft, inProgress, scheduled, completed, failed,
	// cancelled, excluded, deleted, unknownFutureValue.
	Status *CampaignStatus `json:"status,omitempty"`
}
