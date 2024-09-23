package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsSearchSettings struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Enables the developer to define the appearance of the content and configure conditions that dictate when the template
	// should be displayed. Maximum of 2 search result templates per connection.
	SearchResultTemplates *[]ExternalConnectorsDisplayTemplate `json:"searchResultTemplates,omitempty"`
}
