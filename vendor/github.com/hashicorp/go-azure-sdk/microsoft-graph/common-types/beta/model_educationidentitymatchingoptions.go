package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityMatchingOptions struct {
	AppliesTo *EducationUserRole `json:"appliesTo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The name of the source property, which should be a field name in the source data. This property is case-sensitive.
	SourcePropertyName *string `json:"sourcePropertyName,omitempty"`

	// The domain to suffix with the source property to match on the target. If provided as null, the source property will
	// be used to match with the target property.
	TargetDomain nullable.Type[string] `json:"targetDomain,omitempty"`

	// The name of the target property, which should be a valid property in Microsoft Entra ID. This property is
	// case-sensitive.
	TargetPropertyName *string `json:"targetPropertyName,omitempty"`
}
