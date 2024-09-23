package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppRevokeLicensesActionResult struct {
	// Possible types of reasons for an Apple Volume Purchase Program token action failure.
	ActionFailureReason *VppTokenActionFailureReason `json:"actionFailureReason,omitempty"`

	// Action name
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	ActionState *ActionState `json:"actionState,omitempty"`

	// A count of the number of licenses for which revoke failed.
	FailedLicensesCount *int64 `json:"failedLicensesCount,omitempty"`

	// Time the action state was last updated
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// DeviceId associated with the action.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time the action was initiated
	StartDateTime *string `json:"startDateTime,omitempty"`

	// A count of the number of licenses for which revoke was attempted.
	TotalLicensesCount *int64 `json:"totalLicensesCount,omitempty"`

	// UserId associated with the action.
	UserId nullable.Type[string] `json:"userId,omitempty"`
}
