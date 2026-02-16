package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfOfficeSettings struct {
	// True if either:It is currently in the out of office time window configured on the Outlook or Teams client.There is
	// currently an event on the user's calendar that's marked as Show as Out of OfficeOtherwise, false.
	IsOutOfOffice nullable.Type[bool] `json:"isOutOfOffice,omitempty"`

	// The out of office message that the user configured on Outlook client (Automatic Replies (Out of Office)) or the Teams
	// client (Schedule out of office).
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
