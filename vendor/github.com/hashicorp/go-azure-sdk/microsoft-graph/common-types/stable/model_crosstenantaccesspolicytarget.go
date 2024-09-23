package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTarget struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the target for cross-tenant access policy settings and can have one of the following values: The unique
	// identifier of the user, group, or application AllUsers AllApplications - Refers to any Microsoft cloud application.
	// Office365 - Includes the applications mentioned as part of the Office 365 suite.
	Target nullable.Type[string] `json:"target,omitempty"`

	// The type of resource that you want to target. The possible values are: user, group, application, unknownFutureValue.
	TargetType *CrossTenantAccessPolicyTargetType `json:"targetType,omitempty"`
}
