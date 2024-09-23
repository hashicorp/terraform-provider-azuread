package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConfigurationValidation struct {
	// Errors in the validation result of a customAuthenticationExtension.
	Errors *[]GenericError `json:"errors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Warnings in the validation result of a customAuthenticationExtension.
	Warnings *[]GenericError `json:"warnings,omitempty"`
}
