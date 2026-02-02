package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConstraint = DeviceManagementSettingAbstractImplementationConstraint{}

type DeviceManagementSettingAbstractImplementationConstraint struct {
	// List of value which means not configured for the setting
	AllowedAbstractImplementationDefinitionIds *[]string `json:"allowedAbstractImplementationDefinitionIds,omitempty"`

	// Fields inherited from DeviceManagementConstraint

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementSettingAbstractImplementationConstraint) DeviceManagementConstraint() BaseDeviceManagementConstraintImpl {
	return BaseDeviceManagementConstraintImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementSettingAbstractImplementationConstraint{}

func (s DeviceManagementSettingAbstractImplementationConstraint) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementSettingAbstractImplementationConstraint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementSettingAbstractImplementationConstraint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingAbstractImplementationConstraint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingAbstractImplementationConstraint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementSettingAbstractImplementationConstraint: %+v", err)
	}

	return encoded, nil
}
