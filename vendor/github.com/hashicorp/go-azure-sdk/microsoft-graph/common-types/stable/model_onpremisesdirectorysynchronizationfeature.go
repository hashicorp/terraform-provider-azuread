package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesDirectorySynchronizationFeature struct {
	// Used to block cloud object takeover via source anchor hard match if enabled.
	BlockCloudObjectTakeoverThroughHardMatchEnabled nullable.Type[bool] `json:"blockCloudObjectTakeoverThroughHardMatchEnabled,omitempty"`

	// Use to block soft match for all objects if enabled for the tenant. Customers are encouraged to enable this feature
	// and keep it enabled until soft matching is required again for their tenancy. This flag should be enabled again after
	// any soft matching has been completed and is no longer needed.
	BlockSoftMatchEnabled nullable.Type[bool] `json:"blockSoftMatchEnabled,omitempty"`

	// When true, persists the values of Mobile and OtherMobile in on-premises AD during sync cycles instead of values of
	// MobilePhone or AlternateMobilePhones in Microsoft Entra ID.
	BypassDirSyncOverridesEnabled nullable.Type[bool] `json:"bypassDirSyncOverridesEnabled,omitempty"`

	// Used to indicate that cloud password policy applies to users whose passwords are synchronized from on-premises.
	CloudPasswordPolicyForPasswordSyncedUsersEnabled nullable.Type[bool] `json:"cloudPasswordPolicyForPasswordSyncedUsersEnabled,omitempty"`

	// Used to enable concurrent user credentials update in OrgId.
	ConcurrentCredentialUpdateEnabled nullable.Type[bool] `json:"concurrentCredentialUpdateEnabled,omitempty"`

	// Used to enable concurrent user creation in OrgId.
	ConcurrentOrgIdProvisioningEnabled nullable.Type[bool] `json:"concurrentOrgIdProvisioningEnabled,omitempty"`

	// Used to indicate that device write-back is enabled.
	DeviceWritebackEnabled nullable.Type[bool] `json:"deviceWritebackEnabled,omitempty"`

	// Used to indicate that directory extensions are being synced from on-premises AD to Microsoft Entra ID.
	DirectoryExtensionsEnabled nullable.Type[bool] `json:"directoryExtensionsEnabled,omitempty"`

	// Used to indicate that for a Microsoft Forefront Online Protection for Exchange (FOPE) migrated tenant, the
	// conflicting proxy address should be migrated over.
	FopeConflictResolutionEnabled nullable.Type[bool] `json:"fopeConflictResolutionEnabled,omitempty"`

	// Used to enable object-level group writeback feature for additional group types.
	GroupWriteBackEnabled nullable.Type[bool] `json:"groupWriteBackEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Used to indicate on-premise password synchronization is enabled.
	PasswordSyncEnabled nullable.Type[bool] `json:"passwordSyncEnabled,omitempty"`

	// Used to indicate that writeback of password resets from Microsoft Entra ID to on-premises AD is enabled. This
	// property isn't in use and updating it isn't supported.
	PasswordWritebackEnabled nullable.Type[bool] `json:"passwordWritebackEnabled,omitempty"`

	// Used to indicate that we should quarantine objects with conflicting proxy address.
	QuarantineUponProxyAddressesConflictEnabled nullable.Type[bool] `json:"quarantineUponProxyAddressesConflictEnabled,omitempty"`

	// Used to indicate that we should quarantine objects conflicting with duplicate userPrincipalName.
	QuarantineUponUpnConflictEnabled nullable.Type[bool] `json:"quarantineUponUpnConflictEnabled,omitempty"`

	// Used to indicate that we should soft match objects based on userPrincipalName.
	SoftMatchOnUpnEnabled nullable.Type[bool] `json:"softMatchOnUpnEnabled,omitempty"`

	// Used to indicate that we should synchronize userPrincipalName objects for managed users with licenses.
	SynchronizeUpnForManagedUsersEnabled nullable.Type[bool] `json:"synchronizeUpnForManagedUsersEnabled,omitempty"`

	// Used to indicate that Microsoft 365 Group write-back is enabled.
	UnifiedGroupWritebackEnabled nullable.Type[bool] `json:"unifiedGroupWritebackEnabled,omitempty"`

	// Used to indicate that feature to force password change for a user on logon is enabled while synchronizing on-premise
	// credentials.
	UserForcePasswordChangeOnLogonEnabled nullable.Type[bool] `json:"userForcePasswordChangeOnLogonEnabled,omitempty"`

	// Used to indicate that user writeback is enabled.
	UserWritebackEnabled nullable.Type[bool] `json:"userWritebackEnabled,omitempty"`
}
