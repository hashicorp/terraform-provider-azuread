package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingValueTemplateReference struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting value template id
	SettingValueTemplateId *string `json:"settingValueTemplateId,omitempty"`

	// Indicates whether to update policy setting value to match template setting default value
	UseTemplateDefault *bool `json:"useTemplateDefault,omitempty"`
}
