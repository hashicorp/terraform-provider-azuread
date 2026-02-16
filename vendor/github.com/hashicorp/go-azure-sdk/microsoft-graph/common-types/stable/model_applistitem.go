package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppListItem struct {
	// The application or bundle identifier of the application
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The Store URL of the application
	AppStoreUrl nullable.Type[string] `json:"appStoreUrl,omitempty"`

	// The application name
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The publisher of the application
	Publisher nullable.Type[string] `json:"publisher,omitempty"`
}
