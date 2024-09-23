package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCAccountManagerPolicy struct {
	// Possible values for when accounts are deleted on a shared PC.
	AccountDeletionPolicy *SharedPCAccountDeletionPolicyType `json:"accountDeletionPolicy,omitempty"`

	// Sets the percentage of available disk space a PC should have before it stops deleting cached shared PC accounts. Only
	// applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to
	// 100
	CacheAccountsAboveDiskFreePercentage nullable.Type[int64] `json:"cacheAccountsAboveDiskFreePercentage,omitempty"`

	// Specifies when the accounts will start being deleted when they have not been logged on during the specified period,
	// given as number of days. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or
	// DiskSpaceThresholdOrInactiveThreshold.
	InactiveThresholdDays nullable.Type[int64] `json:"inactiveThresholdDays,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Sets the percentage of disk space remaining on a PC before cached accounts will be deleted to free disk space.
	// Accounts that have been inactive the longest will be deleted first. Only applies when AccountDeletionPolicy is
	// DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
	RemoveAccountsBelowDiskFreePercentage nullable.Type[int64] `json:"removeAccountsBelowDiskFreePercentage,omitempty"`
}
