package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineContributingPolicy struct {
	// Name of the policy
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique identifier of the policy
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// Authoring source of a policy
	SourceType *SecurityBaselinePolicySourceType `json:"sourceType,omitempty"`
}
