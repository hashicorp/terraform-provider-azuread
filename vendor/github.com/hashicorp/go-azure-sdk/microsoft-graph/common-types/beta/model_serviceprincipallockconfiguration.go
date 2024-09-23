package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalLockConfiguration struct {
	// Enables locking all sensitive properties. The sensitive properties are keyCredentials, passwordCredentials, and
	// tokenEncryptionKeyId.
	AllProperties nullable.Type[bool] `json:"allProperties,omitempty"`

	// Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Sign.
	CredentialsWithUsageSign nullable.Type[bool] `json:"credentialsWithUsageSign,omitempty"`

	// Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Verify.
	// This locks OAuth service principals.
	CredentialsWithUsageVerify nullable.Type[bool] `json:"credentialsWithUsageVerify,omitempty"`

	// Enables or disables service principal lock configuration. To allow the sensitive properties to be updated, update
	// this property to false to disable the lock on the service principal.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Locks the tokenEncryptionKeyId property for modification on the service principal.
	TokenEncryptionKeyId nullable.Type[bool] `json:"tokenEncryptionKeyId,omitempty"`
}
