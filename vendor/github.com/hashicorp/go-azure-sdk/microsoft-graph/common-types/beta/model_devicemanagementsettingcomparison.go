package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingComparison struct {
	// Setting comparison result type
	ComparisonResult *DeviceManagementComparisonResult `json:"comparisonResult,omitempty"`

	// JSON representation of current intent (or) template setting's value
	CurrentValueJson nullable.Type[string] `json:"currentValueJson,omitempty"`

	// The ID of the setting definition for this instance
	DefinitionId nullable.Type[string] `json:"definitionId,omitempty"`

	// The setting's display name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The setting ID
	Id nullable.Type[string] `json:"id,omitempty"`

	// JSON representation of new template setting's value
	NewValueJson nullable.Type[string] `json:"newValueJson,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
