package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDashboardCardBotConfiguration struct {
	// The ID (usually a GUID) of the bot associated with the specific teamsAppDefinition. This is a unique app ID for the
	// bot as registered with the Bot Framework.
	BotId nullable.Type[string] `json:"botId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
