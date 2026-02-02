package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceHealthScriptParameter = DeviceHealthScriptBooleanParameter{}

type DeviceHealthScriptBooleanParameter struct {
	// The default value of boolean param
	DefaultValue *bool `json:"defaultValue,omitempty"`

	// Fields inherited from DeviceHealthScriptParameter

	// Whether Apply DefaultValue When Not Assigned
	ApplyDefaultValueWhenNotAssigned *bool `json:"applyDefaultValueWhenNotAssigned,omitempty"`

	// The description of the param
	Description nullable.Type[string] `json:"description,omitempty"`

	// Whether the param is required
	IsRequired *bool `json:"isRequired,omitempty"`

	// The name of the param
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceHealthScriptBooleanParameter) DeviceHealthScriptParameter() BaseDeviceHealthScriptParameterImpl {
	return BaseDeviceHealthScriptParameterImpl{
		ApplyDefaultValueWhenNotAssigned: s.ApplyDefaultValueWhenNotAssigned,
		Description:                      s.Description,
		IsRequired:                       s.IsRequired,
		Name:                             s.Name,
		ODataId:                          s.ODataId,
		ODataType:                        s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScriptBooleanParameter{}

func (s DeviceHealthScriptBooleanParameter) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScriptBooleanParameter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScriptBooleanParameter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptBooleanParameter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptBooleanParameter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScriptBooleanParameter: %+v", err)
	}

	return encoded, nil
}
