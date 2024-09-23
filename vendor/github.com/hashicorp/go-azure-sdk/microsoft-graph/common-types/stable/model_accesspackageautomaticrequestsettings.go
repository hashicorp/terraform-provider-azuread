package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAutomaticRequestSettings struct {
	// The duration for which access must be retained before the target's access is revoked once they leave the allowed
	// target scope.
	GracePeriodBeforeAccessRemoval nullable.Type[string] `json:"gracePeriodBeforeAccessRemoval,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether automatic assignment must be removed for targets who move out of the allowed target scope.
	RemoveAccessWhenTargetLeavesAllowedTargets nullable.Type[bool] `json:"removeAccessWhenTargetLeavesAllowedTargets,omitempty"`

	// If set to true, automatic assignments will be created for targets in the allowed target scope.
	RequestAccessForAllowedTargets nullable.Type[bool] `json:"requestAccessForAllowedTargets,omitempty"`
}
