package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsSetting struct {
	// The display name for the setting. Required. Read-only.
	DisplayName *string `json:"displayName,omitempty"`

	// The value for the setting serialized as string of JSON. Required. Read-only.
	JsonValue *string `json:"jsonValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
	OverwriteAllowed *bool `json:"overwriteAllowed,omitempty"`

	SettingId nullable.Type[string]                       `json:"settingId,omitempty"`
	ValueType *ManagedTenantsManagementParameterValueType `json:"valueType,omitempty"`
}

var _ json.Marshaler = ManagedTenantsSetting{}

func (s ManagedTenantsSetting) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsSetting: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "jsonValue")
	delete(decoded, "overwriteAllowed")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsSetting: %+v", err)
	}

	return encoded, nil
}
