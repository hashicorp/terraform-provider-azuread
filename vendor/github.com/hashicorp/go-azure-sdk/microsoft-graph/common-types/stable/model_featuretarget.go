package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureTarget struct {
	// The ID of the entity that's targeted in the include or exclude rule, or all_users to target all users.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The kind of entity that's targeted. The possible values are: group, administrativeUnit, role, unknownFutureValue.
	TargetType *FeatureTargetType `json:"targetType,omitempty"`
}
