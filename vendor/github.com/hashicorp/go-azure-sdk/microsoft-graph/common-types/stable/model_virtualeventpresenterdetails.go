package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventPresenterDetails struct {
	Bio                   *ItemBody             `json:"bio,omitempty"`
	Company               nullable.Type[string] `json:"company,omitempty"`
	JobTitle              nullable.Type[string] `json:"jobTitle,omitempty"`
	LinkedInProfileWebUrl nullable.Type[string] `json:"linkedInProfileWebUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PersonalSiteWebUrl   nullable.Type[string] `json:"personalSiteWebUrl,omitempty"`
	Photo                nullable.Type[string] `json:"photo,omitempty"`
	TwitterProfileWebUrl nullable.Type[string] `json:"twitterProfileWebUrl,omitempty"`
}
