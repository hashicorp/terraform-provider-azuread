package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailSenderDetail struct {
	// Display name of sender from address.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Date and time of creation of the sender domain.
	DomainCreationDateTime nullable.Type[string] `json:"domainCreationDateTime,omitempty"`

	// Registered name of the domain.
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// Owner of the domain.
	DomainOwner nullable.Type[string] `json:"domainOwner,omitempty"`

	// The sender email address in the mail From header, also known as the envelope sender or the P1 sender.
	FromAddress nullable.Type[string] `json:"fromAddress,omitempty"`

	// The IPv4 address of the last detected mail server that relayed the message.
	IPv4 nullable.Type[string] `json:"ipv4,omitempty"`

	// Location of the domain.
	Location nullable.Type[string] `json:"location,omitempty"`

	// The sender email address in the From header, which is visible to email recipients on their email clients. Also known
	// as P2 sender.
	MailFromAddress nullable.Type[string] `json:"mailFromAddress,omitempty"`

	// Domain name of sender mail from address.
	MailFromDomainName nullable.Type[string] `json:"mailFromDomainName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
