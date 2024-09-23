package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetonationObservables struct {
	// The list of all contacted IPs in the detonation.
	ContactedIps *[]string `json:"contactedIps,omitempty"`

	// The list of all URLs found in the detonation.
	ContactedUrls *[]string `json:"contactedUrls,omitempty"`

	// The list of all dropped files in the detonation.
	Droppedfiles *[]string `json:"droppedfiles,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
