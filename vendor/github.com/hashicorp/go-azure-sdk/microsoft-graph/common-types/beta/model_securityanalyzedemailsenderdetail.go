package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailSenderDetail struct {
	// The sender email address in the mail From header, also known as the envelope sender or the P1 sender.
	FromAddress nullable.Type[string] `json:"fromAddress,omitempty"`

	// The IPv4 address of the last detected mail server that relayed the message.
	IPv4 nullable.Type[string] `json:"ipv4,omitempty"`

	// The sender email address in the From header, which is visible to email recipients on their email clients. Also known
	// as P2 sender.
	MailFromAddress nullable.Type[string] `json:"mailFromAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
