package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationPolicyTip struct {
	// The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about
	// what users shouldn't say in chats)
	ComplianceUrl nullable.Type[string] `json:"complianceUrl,omitempty"`

	// Explanatory text shown to the sender of the message.
	GeneralText nullable.Type[string] `json:"generalText,omitempty"`

	// The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its
	// own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
	MatchedConditionDescriptions *[]string `json:"matchedConditionDescriptions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
