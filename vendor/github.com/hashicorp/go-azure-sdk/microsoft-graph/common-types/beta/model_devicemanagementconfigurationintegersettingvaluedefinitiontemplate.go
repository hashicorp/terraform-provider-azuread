package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate struct {
	// Integer Setting Maximum Value. Valid values -2147483648 to 2147483647
	MaxValue *int64 `json:"maxValue,omitempty"`

	// Integer Setting Minimum Value. Valid values -2147483648 to 2147483647
	MinValue *int64 `json:"minValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
