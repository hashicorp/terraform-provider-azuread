package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectSharingAbilities struct {
	// Indicates whether the current user can add existing guest recipients to this item using direct sharing.
	AddExistingExternalUsers *SharingOperationStatus `json:"addExistingExternalUsers,omitempty"`

	// Indicates whether the current user can add internal recipients to this item using direct sharing.
	AddInternalUsers *SharingOperationStatus `json:"addInternalUsers,omitempty"`

	// Indicates whether the current user can add new guest recipients to this item using direct sharing.
	AddNewExternalUsers *SharingOperationStatus `json:"addNewExternalUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether the user querying this endpoint can request access for the user or on behalf of other users, after
	// which, site admins, can approve or deny the creation of a potential sharing link.
	RequestGrantAccess *SharingOperationStatus `json:"requestGrantAccess,omitempty"`
}
