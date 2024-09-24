package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConstraint = DeviceManagementSettingRequiredConstraint{}

type DeviceManagementSettingRequiredConstraint struct {
	// List of value which means not configured for the setting
	NotConfiguredValue nullable.Type[string] `json:"notConfiguredValue,omitempty"`

	// Fields inherited from DeviceManagementConstraint

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementSettingRequiredConstraint) DeviceManagementConstraint() BaseDeviceManagementConstraintImpl {
	return BaseDeviceManagementConstraintImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementSettingRequiredConstraint{}

func (s DeviceManagementSettingRequiredConstraint) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementSettingRequiredConstraint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementSettingRequiredConstraint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingRequiredConstraint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingRequiredConstraint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementSettingRequiredConstraint: %+v", err)
	}

	return encoded, nil
}
