package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentCustomizedSetting struct {
	// JSON representation of the customized value, if different from default
	CustomizedJson nullable.Type[string] `json:"customizedJson,omitempty"`

	// JSON representation of the default value from the template
	DefaultJson nullable.Type[string] `json:"defaultJson,omitempty"`

	// The ID of the setting definition for this setting
	DefinitionId nullable.Type[string] `json:"definitionId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
