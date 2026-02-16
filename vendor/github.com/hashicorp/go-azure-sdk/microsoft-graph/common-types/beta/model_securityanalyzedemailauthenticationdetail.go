package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailAuthenticationDetail struct {
	// A value used by Microsoft 365 to combine email authentication such as SPF, DKIM, and DMARC, to determine whether the
	// message is authentic.
	CompositeAuthentication nullable.Type[string] `json:"compositeAuthentication,omitempty"`

	// DomainKeys identified mail (DKIM). Indicates whether it was pass/fail/soft fail.
	Dkim nullable.Type[string] `json:"dkim,omitempty"`

	// Domain-based Message Authentication. Indicates whether it was pass/fail/soft fail.
	Dmarc nullable.Type[string] `json:"dmarc,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Sender Policy Framework (SPF). Indicates whether it was pass/fail/soft fail.
	SenderPolicyFramework nullable.Type[string] `json:"senderPolicyFramework,omitempty"`
}
