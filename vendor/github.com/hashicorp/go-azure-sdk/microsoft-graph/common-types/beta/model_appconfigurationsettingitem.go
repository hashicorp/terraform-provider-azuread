package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConfigurationSettingItem struct {
	// app configuration key.
	AppConfigKey *string `json:"appConfigKey,omitempty"`

	// App configuration key types.
	AppConfigKeyType *MdmAppConfigKeyType `json:"appConfigKeyType,omitempty"`

	// app configuration key value.
	AppConfigKeyValue *string `json:"appConfigKeyValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
