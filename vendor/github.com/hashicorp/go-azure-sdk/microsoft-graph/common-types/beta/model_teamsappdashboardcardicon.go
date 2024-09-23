package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDashboardCardIcon struct {
	// The icon for the card, displayed in the toolbox and card bar, is represented as a URL. The preferred size for raster
	// images is 28x28 pixels. If this property has a value, the officeFabricIconFontName property is ignored.
	IconUrl nullable.Type[string] `json:"iconUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The friendly name of the Office UI Fabric/Fluent UI icon for the card that is used when the iconUrl property isn't
	// specified. For example, 'officeUIFabricIconName': 'Search'.
	OfficeUIFabricIconName nullable.Type[string] `json:"officeUIFabricIconName,omitempty"`
}
