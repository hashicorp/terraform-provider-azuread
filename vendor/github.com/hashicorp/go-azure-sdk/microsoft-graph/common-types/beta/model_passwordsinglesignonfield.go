package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordSingleSignOnField struct {
	// Title/label override for customization.
	CustomizedLabel nullable.Type[string] `json:"customizedLabel,omitempty"`

	// Label that would be used if no customizedLabel is provided. Read only.
	DefaultLabel nullable.Type[string] `json:"defaultLabel,omitempty"`

	// Id used to identity the field type. This is an internal ID and possible values are param1, param2, paramuserName,
	// parampassword.
	FieldId nullable.Type[string] `json:"fieldId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of the credential. The values can be text, password.
	Type nullable.Type[string] `json:"type,omitempty"`
}
