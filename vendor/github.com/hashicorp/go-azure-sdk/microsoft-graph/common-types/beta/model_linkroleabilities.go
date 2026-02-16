package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkRoleAbilities struct {
	// Indicates if the current user can add existing external user recipients to this sharing link.
	AddExistingExternalUsers *SharingOperationStatus `json:"addExistingExternalUsers,omitempty"`

	// Indicates if the current user can add new external user recipients to this sharing link.
	AddNewExternalUsers *SharingOperationStatus `json:"addNewExternalUsers,omitempty"`

	// Indicates the status of the potential sharing link variants. If selected, it generates a separate sharing link from
	// the sharing link that would otherwise be generated without the variant, yet with an identical role and scope.
	ApplyVariants *SharingLinkVariants `json:"applyVariants,omitempty"`

	// Indicates if links of this role can be created.
	CreateLink *SharingOperationStatus `json:"createLink,omitempty"`

	// Indicates if links of this role can be deleted.
	DeleteLink *SharingOperationStatus `json:"deleteLink,omitempty"`

	// Indicates if this link can include external users.
	LinkAllowsExternalUsers *SharingOperationStatus `json:"linkAllowsExternalUsers,omitempty"`

	// Indicates if links must expire, meaning the link is no longer usable after a specified time. If link expiration is
	// enabled, a default link expiration time is provided.
	LinkExpiration *SharingLinkExpirationStatus `json:"linkExpiration,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates if links of this role can be retrieved.
	RetrieveLink *SharingOperationStatus `json:"retrieveLink,omitempty"`

	// Indicates if links of this role can be updated.
	UpdateLink *SharingOperationStatus `json:"updateLink,omitempty"`
}
