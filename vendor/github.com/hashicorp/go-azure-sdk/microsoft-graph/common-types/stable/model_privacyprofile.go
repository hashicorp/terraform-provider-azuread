package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivacyProfile struct {
	// A valid smtp email address for the privacy statement contact. Not required.
	ContactEmail nullable.Type[string] `json:"contactEmail,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to
	// the company's privacy statement. Not required.
	StatementUrl nullable.Type[string] `json:"statementUrl,omitempty"`
}
