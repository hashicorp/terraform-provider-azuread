package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryState struct {
	AppId      nullable.Type[string] `json:"appId,omitempty"`
	AssignedTo nullable.Type[string] `json:"assignedTo,omitempty"`
	Comments   *[]string             `json:"comments,omitempty"`
	Feedback   *AlertFeedback        `json:"feedback,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Status          *AlertStatus          `json:"status,omitempty"`
	UpdatedDateTime nullable.Type[string] `json:"updatedDateTime,omitempty"`
	User            nullable.Type[string] `json:"user,omitempty"`
}
