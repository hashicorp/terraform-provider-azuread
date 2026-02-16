package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkOnPremisesCalendarSyncConfiguration struct {
	// The fully qualified domain name (FQDN) of the Skype for Business Server. Use the Exchange domain if the Skype for
	// Business SIP domain is different from the Exchange domain of the user.
	Domain nullable.Type[string] `json:"domain,omitempty"`

	// The domain and username of the console device, for example, Seattle/RanierConf.
	DomainUserName nullable.Type[string] `json:"domainUserName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The Simple Mail Transfer Protocol (SMTP) address of the user account. This is only required if a different user
	// principal name (UPN) is used to sign in to Exchange other than Microsoft Teams and Skype for Business. This is a
	// common scenario in a hybrid environment where an on-premises Exchange server is used.
	SmtpAddress nullable.Type[string] `json:"smtpAddress,omitempty"`
}
