package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConstraint = DeviceManagementSettingRegexConstraint{}

type DeviceManagementSettingRegexConstraint struct {
	// The RegEx pattern to match against
	Regex *string `json:"regex,omitempty"`

	// Fields inherited from DeviceManagementConstraint

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementSettingRegexConstraint) DeviceManagementConstraint() BaseDeviceManagementConstraintImpl {
	return BaseDeviceManagementConstraintImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementSettingRegexConstraint{}

func (s DeviceManagementSettingRegexConstraint) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementSettingRegexConstraint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementSettingRegexConstraint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingRegexConstraint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingRegexConstraint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementSettingRegexConstraint: %+v", err)
	}

	return encoded, nil
}
