package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomAppManagementApplicationConfiguration struct {
	// Property to restrict creation or update of apps based on their target signInAudience types.
	Audiences *AudiencesConfiguration `json:"audiences,omitempty"`

	// Configuration for identifierUris restrictions.
	IdentifierUris *IdentifierUriConfiguration `json:"identifierUris,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
