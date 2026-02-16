package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConstraint = DeviceManagementSettingCollectionConstraint{}

type DeviceManagementSettingCollectionConstraint struct {
	// The maximum number of elements in the collection
	MaximumLength nullable.Type[int64] `json:"maximumLength,omitempty"`

	// The minimum number of elements in the collection
	MinimumLength nullable.Type[int64] `json:"minimumLength,omitempty"`

	// Fields inherited from DeviceManagementConstraint

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementSettingCollectionConstraint) DeviceManagementConstraint() BaseDeviceManagementConstraintImpl {
	return BaseDeviceManagementConstraintImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementSettingCollectionConstraint{}

func (s DeviceManagementSettingCollectionConstraint) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementSettingCollectionConstraint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementSettingCollectionConstraint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingCollectionConstraint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingCollectionConstraint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementSettingCollectionConstraint: %+v", err)
	}

	return encoded, nil
}
