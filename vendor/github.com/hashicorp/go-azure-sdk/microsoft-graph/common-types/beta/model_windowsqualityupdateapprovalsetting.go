package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateApprovalSetting struct {
	// Enum type to describe the approval type for different type of quality updates.
	ApprovalMethodType *WindowsQualityUpdatePolicyApprovalMethodType `json:"approvalMethodType,omitempty"`

	// The deferral days for auto approval type, not applicable for manual approve
	DeferredDeploymentInDay nullable.Type[int64] `json:"deferredDeploymentInDay,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The publishing cadence of the quality update. Possible values are: monthly, outOfBand. This property cannot be
	// modified and is automatically populated when the catalog is created.
	WindowsQualityUpdateCadence *WindowsQualityUpdateCadence `json:"windowsQualityUpdateCadence,omitempty"`

	// Windows quality update category
	WindowsQualityUpdateCategory *WindowsQualityUpdateCategory `json:"windowsQualityUpdateCategory,omitempty"`
}
