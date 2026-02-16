package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobApplicationParameters struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given
	// synchronization job or template.
	RuleId nullable.Type[string] `json:"ruleId,omitempty"`

	// The identifiers of one or more objects to which a synchronizationJob is to be applied.
	Subjects *[]SynchronizationJobSubject `json:"subjects,omitempty"`
}
