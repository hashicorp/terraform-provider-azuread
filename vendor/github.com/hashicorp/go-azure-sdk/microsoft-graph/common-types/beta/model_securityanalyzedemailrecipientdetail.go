package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailRecipientDetail struct {
	// Recipient address in the cc field.
	CcRecipients *[]string `json:"ccRecipients,omitempty"`

	// Domain name of the recipient.
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
