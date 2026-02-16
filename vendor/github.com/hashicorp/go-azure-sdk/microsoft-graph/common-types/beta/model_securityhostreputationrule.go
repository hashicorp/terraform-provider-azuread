package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostReputationRule struct {
	// The description of the rule that gives more context.
	Description *string `json:"description,omitempty"`

	// The name of the rule.
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Link to a web page with details related to this rule.
	RelatedDetailsUrl nullable.Type[string] `json:"relatedDetailsUrl,omitempty"`

	Severity *SecurityHostReputationRuleSeverity `json:"severity,omitempty"`
}
