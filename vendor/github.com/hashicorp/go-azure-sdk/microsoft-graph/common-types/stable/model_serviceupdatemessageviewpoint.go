package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageViewpoint struct {
	// Indicates whether the user archived the message.
	IsArchived nullable.Type[bool] `json:"isArchived,omitempty"`

	// Indicates whether the user marked the message as favorite.
	IsFavorited nullable.Type[bool] `json:"isFavorited,omitempty"`

	// Indicates whether the user read the message.
	IsRead nullable.Type[bool] `json:"isRead,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
