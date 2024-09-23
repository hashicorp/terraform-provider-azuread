package beta

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

	// ID for the request approver.
	TicketApproverIdentityId nullable.Type[string] `json:"ticketApproverIdentityId,omitempty"`

	// The ticket number.
	TicketNumber nullable.Type[string] `json:"ticketNumber,omitempty"`

	// ID for the request submitter.
	TicketSubmitterIdentityId nullable.Type[string] `json:"ticketSubmitterIdentityId,omitempty"`

	// The description of the ticket system.
	TicketSystem nullable.Type[string] `json:"ticketSystem,omitempty"`
}
