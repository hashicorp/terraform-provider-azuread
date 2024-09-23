package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSetContent struct {
	// Content type information of the file.
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`

	// Name of the file in resource folder that should be added as a default content or a template in the document set
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// Folder name in which the file will be placed when a new document set is created in the library.
	FolderName nullable.Type[string] `json:"folderName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
