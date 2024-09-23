package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDependedOnBy struct {
	// Identifier of child setting that is dependent on the current setting
	DependedOnBy nullable.Type[string] `json:"dependedOnBy,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value that determines if the child setting is required based on the parent setting's selection
	Required nullable.Type[bool] `json:"required,omitempty"`
}
