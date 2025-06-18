package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FederatedIdentityExpression struct {
	// Indicated the language version to be used. Should always be set to 1. Required.
	LanguageVersion int64 `json:"languageVersion"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the configured expression. Required.
	Value string `json:"value"`
}
