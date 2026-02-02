package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelIdentity struct {
	// The identity of the channel in which the message was posted.
	ChannelId nullable.Type[string] `json:"channelId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identity of the team in which the message was posted.
	TeamId nullable.Type[string] `json:"teamId,omitempty"`
}
