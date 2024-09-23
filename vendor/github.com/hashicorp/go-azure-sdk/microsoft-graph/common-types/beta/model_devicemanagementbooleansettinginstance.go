package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementSettingInstance = DeviceManagementBooleanSettingInstance{}

type DeviceManagementBooleanSettingInstance struct {
	// The boolean value
	Value nullable.Type[bool] `json:"value,omitempty"`

	// Fields inherited from DeviceManagementSettingInstance

	// The ID of the setting definition for this instance
	DefinitionId *string `json:"definitionId,omitempty"`

	// JSON representation of the value
	ValueJson nullable.Type[string] `json:"valueJson,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementBooleanSettingInstance) DeviceManagementSettingInstance() BaseDeviceManagementSettingInstanceImpl {
	return BaseDeviceManagementSettingInstanceImpl{
		DefinitionId: s.DefinitionId,
		ValueJson:    s.ValueJson,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s DeviceManagementBooleanSettingInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementBooleanSettingInstance{}

func (s DeviceManagementBooleanSettingInstance) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementBooleanSettingInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementBooleanSettingInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementBooleanSettingInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementBooleanSettingInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementBooleanSettingInstance: %+v", err)
	}

	return encoded, nil
}
