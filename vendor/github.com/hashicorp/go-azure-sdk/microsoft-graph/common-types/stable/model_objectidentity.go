package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectIdentity struct {
	// Specifies the issuer of the identity, for example facebook.com. 512 character limit. For local accounts (where
	// signInType isn't federated), this property is the local default domain name for the tenant, for example contoso.com.
	// For guests from other Microsoft Entra organizations, this is the domain of the federated organization, for example
	// contoso.com. For more information about filtering behavior for this property, see Filtering on the identities
	// property of a user.
	Issuer nullable.Type[string] `json:"issuer,omitempty"`

	// Specifies the unique identifier assigned to the user by the issuer. 64 character limit. The combination of issuer and
	// issuerAssignedId must be unique within the organization. Represents the sign-in name for the user, when signInType is
	// set to emailAddress or userName (also known as local accounts).When signInType is set to: emailAddress (or a custom
	// string that starts with emailAddress like emailAddress1), issuerAssignedId must be a valid email addressuserName,
	// issuerAssignedId must begin with an alphabetical character or number, and can only contain alphanumeric characters
	// and the following symbols: - or _ For more information about filtering behavior for this property, see Filtering on
	// the identities property of a user.
	IssuerAssignedId nullable.Type[string] `json:"issuerAssignedId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the user sign-in types in your directory, such as emailAddress, userName, federated, or userPrincipalName.
	// federated represents a unique identifier for a user from an issuer that can be in any format chosen by the issuer.
	// Setting or updating a userPrincipalName identity updates the value of the userPrincipalName property on the user
	// object. The validations performed on the userPrincipalName property on the user object, for example, verified domains
	// and acceptable characters, are performed when setting or updating a userPrincipalName identity. Extra validation is
	// enforced on issuerAssignedId when the sign-in type is set to emailAddress or userName. This property can also be set
	// to any custom string. For more information about filtering behavior for this property, see Filtering on the
	// identities property of a user.
	SignInType nullable.Type[string] `json:"signInType,omitempty"`
}
