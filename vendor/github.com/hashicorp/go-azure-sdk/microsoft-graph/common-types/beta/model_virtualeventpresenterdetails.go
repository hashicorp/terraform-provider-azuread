package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventPresenterDetails struct {
	// Bio of the presenter.
	Bio *ItemBody `json:"bio,omitempty"`

	// The presenter's company name.
	Company nullable.Type[string] `json:"company,omitempty"`

	// The presenter's job title.
	JobTitle nullable.Type[string] `json:"jobTitle,omitempty"`

	// The presenter's LinkedIn profile URL.
	LinkedInProfileWebUrl nullable.Type[string] `json:"linkedInProfileWebUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The presenter's personal website URL.
	PersonalSiteWebUrl nullable.Type[string] `json:"personalSiteWebUrl,omitempty"`

	// The content stream of the presenter's photo.
	Photo nullable.Type[string] `json:"photo,omitempty"`

	// The presenter's Twitter profile URL.
	TwitterProfileWebUrl nullable.Type[string] `json:"twitterProfileWebUrl,omitempty"`
}
