package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSAssociatedDomainsItem struct {
	// The application identifier of the app to associate domains with.
	ApplicationIdentifier *string `json:"applicationIdentifier,omitempty"`

	// Determines whether data should be downloaded directly or via a CDN.
	DirectDownloadsEnabled *bool `json:"directDownloadsEnabled,omitempty"`

	// The list of domains to associate.
	Domains *[]string `json:"domains,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
