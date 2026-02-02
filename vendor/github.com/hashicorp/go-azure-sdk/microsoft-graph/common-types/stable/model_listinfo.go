package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListInfo struct {
	// If true, indicates that content types are enabled for this list.
	ContentTypesEnabled nullable.Type[bool] `json:"contentTypesEnabled,omitempty"`

	// If true, indicates that the list isn't normally visible in the SharePoint user experience.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// An enumerated value that represents the base list template used in creating the list. Possible values include
	// documentLibrary, genericList, task, survey, announcements, contacts, and more.
	Template nullable.Type[string] `json:"template,omitempty"`
}
