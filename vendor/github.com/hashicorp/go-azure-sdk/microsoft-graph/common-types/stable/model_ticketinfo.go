package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TicketInfo struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ticket number.
	TicketNumber nullable.Type[string] `json:"ticketNumber,omitempty"`

	// The description of the ticket system.
	TicketSystem nullable.Type[string] `json:"ticketSystem,omitempty"`
}
