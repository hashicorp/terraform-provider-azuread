package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MfaDetail struct {
	// Indicates the MFA auth detail for the corresponding sign-in activity when the MFA Required is 'Yes'.
	AuthDetail nullable.Type[string] `json:"authDetail,omitempty"`

	// Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the values) for the corresponding sign-in
	// activity when the MFA Required field is 'Yes'.
	AuthMethod nullable.Type[string] `json:"authMethod,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
