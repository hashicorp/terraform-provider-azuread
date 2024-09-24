package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebPartData struct {
	// Audience information of the web part. By using this property, specific content is prioritized to specific audiences.
	Audiences *[]string `json:"audiences,omitempty"`

	// Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers
	// to a different property structure.
	DataVersion nullable.Type[string] `json:"dataVersion,omitempty"`

	// Description of the web part.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains collections of data that can be processed by server side services like search index and link fixup.
	ServerProcessedContent *ServerProcessedContent `json:"serverProcessedContent,omitempty"`

	// Title of the web part.
	Title nullable.Type[string] `json:"title,omitempty"`
}
