package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionScope struct {
	// A description of the delegated permissions, intended to be read by an administrator granting the permission on behalf
	// of all users. This text appears in tenant-wide admin consent experiences.
	AdminConsentDescription nullable.Type[string] `json:"adminConsentDescription,omitempty"`

	// The permission's title, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName nullable.Type[string] `json:"adminConsentDisplayName,omitempty"`

	// Unique delegated permission identifier inside the collection of delegated permissions defined for a resource
	// application.
	Id *string `json:"id,omitempty"`

	// When you create or update a permission, this property must be set to true (which is the default). To delete a
	// permission, this property must first be set to false. At that point, in a subsequent call, the permission may be
	// removed.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Origin nullable.Type[string] `json:"origin,omitempty"`

	// The possible values are: User and Admin. Specifies whether this delegated permission should be considered safe for
	// non-admin users to consent to on behalf of themselves, or whether an administrator consent should always be required.
	// While Microsoft Graph defines the default consent requirement for each permission, the tenant administrator may
	// override the behavior in their organization (by allowing, restricting, or limiting user consent to this delegated
	// permission). For more information, see Configure how users consent to applications.
	Type nullable.Type[string] `json:"type,omitempty"`

	// A description of the delegated permissions, intended to be read by a user granting the permission on their own
	// behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
	UserConsentDescription nullable.Type[string] `json:"userConsentDescription,omitempty"`

	// A title for the permission, intended to be read by a user granting the permission on their own behalf. This text
	// appears in consent experiences where the user is consenting only on behalf of themselves.
	UserConsentDisplayName nullable.Type[string] `json:"userConsentDisplayName,omitempty"`

	// Specifies the value to include in the scp (scope) claim in access tokens. Must not exceed 120 characters in length.
	// Allowed characters are : ! # $ % & ' ( ) * + , - . / : ; = ? @ [ ] ^ + _ { } ~, and characters in the ranges 0-9, A-Z
	// and a-z. Any other character, including the space character, aren't allowed. May not begin with ..
	Value nullable.Type[string] `json:"value,omitempty"`
}
