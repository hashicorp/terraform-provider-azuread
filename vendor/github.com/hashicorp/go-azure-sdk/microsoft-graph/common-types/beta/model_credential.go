package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Credential struct {
	// The name of the field for this credential. e.g, username or password or phoneNumber. This is defined by the
	// application. Must match what is in the html field on singleSignOnSettings/password object.
	FieldId nullable.Type[string] `json:"fieldId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type for this credential. Valid values: username, password, or other.
	Type nullable.Type[string] `json:"type,omitempty"`

	// The value for this credential. e.g, mysuperhiddenpassword. Note the value for passwords is write-only, the value can
	// never be read back.
	Value nullable.Type[string] `json:"value,omitempty"`
}
