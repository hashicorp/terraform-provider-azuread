package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationDependentOn struct {
	// Identifier of parent setting/ parent setting option dependent on
	DependentOn nullable.Type[string] `json:"dependentOn,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Identifier of parent setting/ parent setting id dependent on
	ParentSettingId nullable.Type[string] `json:"parentSettingId,omitempty"`
}
