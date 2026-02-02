package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedLanguageFile struct {
	// The contents of the uploaded ADML file.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The file name of the uploaded ADML file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// Key of the entity.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The language code of the uploaded ADML file.
	LanguageCode nullable.Type[string] `json:"languageCode,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
