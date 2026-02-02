package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsUrlMatchInfo struct {
	// A list of the URL prefixes that must match URLs to be processed by this URL-to-item-resolver.
	BaseUrls *[]string `json:"baseUrls,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A regular expression that will be matched towards the URL that is processed by this URL-to-item-resolver. The
	// ECMAScript specification for regular expressions (ECMA-262) is used for the evaluation. The named groups defined by
	// the regular expression will be used later to extract values from the URL.
	UrlPattern nullable.Type[string] `json:"urlPattern,omitempty"`
}
