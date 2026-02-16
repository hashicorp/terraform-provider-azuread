package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingLinkVariants struct {
	// Indicates the most permissive role with which an address bar link can be created. The possible values are: none,
	// view, edit, manageList, review, restrictedView, submitOnly, unknownFutureValue.
	AddressBarLinkPermission *SharingRole `json:"addressBarLinkPermission,omitempty"`

	// Indicates whether a link can be embedded.
	AllowEmbed *SharingOperationStatus `json:"allowEmbed,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether a link can be password protected, meaning that link users would need to enter a password to access
	// the item for which the sharing link is produced. Creating a passwordProtected link for the first time requires
	// providing a password.
	PasswordProtected *SharingOperationStatus `json:"passwordProtected,omitempty"`

	// Indicates whether a link requires identity authentication for recipients. Users can be verified through either an
	// email address or identity.
	RequiresAuthentication *SharingOperationStatus `json:"requiresAuthentication,omitempty"`
}
