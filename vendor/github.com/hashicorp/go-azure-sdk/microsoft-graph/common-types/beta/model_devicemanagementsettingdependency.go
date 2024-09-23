package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingDependency struct {
	// Collection of constraints for the dependency setting value
	Constraints *[]DeviceManagementConstraint `json:"constraints,omitempty"`

	// The setting definition ID of the setting depended on
	DefinitionId *string `json:"definitionId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &DeviceManagementSettingDependency{}

func (s *DeviceManagementSettingDependency) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DefinitionId *string `json:"definitionId,omitempty"`
		ODataId      *string `json:"@odata.id,omitempty"`
		ODataType    *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DefinitionId = decoded.DefinitionId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementSettingDependency into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["constraints"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Constraints into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConstraint, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConstraintImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Constraints' for 'DeviceManagementSettingDependency': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Constraints = &output
	}

	return nil
}
