package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageLayoutConfiguration struct {
	// Option to show the footer on the sign-in page.
	IsFooterShown nullable.Type[bool] `json:"isFooterShown,omitempty"`

	// Option to show the header on the sign-in page.
	IsHeaderShown nullable.Type[bool] `json:"isHeaderShown,omitempty"`

	// Represents the layout template to be displayed on the login page for a tenant. The possible values are default -
	// Represents the default Microsoft layout with a centered lightbox. verticalSplit - Represents a layout with a
	// background on the left side and a full-height lightbox to the right. unknownFutureValue - Evolvable enumeration
	// sentinel value. Don't use.
	LayoutTemplateType *LayoutTemplateType `json:"layoutTemplateType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
