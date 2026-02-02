package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordSingleSignOnSettings struct {
	// The fields to capture to fill the user credentials for password-based single sign-on.
	Fields *[]PasswordSingleSignOnField `json:"fields,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
