package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSet struct {
	// Content types allowed in document set.
	AllowedContentTypes *[]ContentTypeInfo `json:"allowedContentTypes,omitempty"`

	// Default contents of document set.
	DefaultContents *[]DocumentSetContent `json:"defaultContents,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether to add the name of the document set to each file name.
	PropagateWelcomePageChanges nullable.Type[bool] `json:"propagateWelcomePageChanges,omitempty"`

	SharedColumns *[]ColumnDefinition `json:"sharedColumns,omitempty"`

	// Add the name of the Document Set to each file name.
	ShouldPrefixNameToFile nullable.Type[bool] `json:"shouldPrefixNameToFile,omitempty"`

	WelcomePageColumns *[]ColumnDefinition `json:"welcomePageColumns,omitempty"`

	// Welcome page absolute URL.
	WelcomePageUrl nullable.Type[string] `json:"welcomePageUrl,omitempty"`
}
