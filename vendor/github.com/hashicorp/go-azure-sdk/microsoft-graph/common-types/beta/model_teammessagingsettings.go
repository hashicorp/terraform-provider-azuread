package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamMessagingSettings struct {
	// If set to true, @channel mentions are allowed.
	AllowChannelMentions nullable.Type[bool] `json:"allowChannelMentions,omitempty"`

	// If set to true, owners can delete any message.
	AllowOwnerDeleteMessages nullable.Type[bool] `json:"allowOwnerDeleteMessages,omitempty"`

	// If set to true, @team mentions are allowed.
	AllowTeamMentions nullable.Type[bool] `json:"allowTeamMentions,omitempty"`

	// If set to true, users can delete their messages.
	AllowUserDeleteMessages nullable.Type[bool] `json:"allowUserDeleteMessages,omitempty"`

	// If set to true, users can edit their messages.
	AllowUserEditMessages nullable.Type[bool] `json:"allowUserEditMessages,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
