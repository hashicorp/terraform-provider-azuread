package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityInformationProtectionAction = SecurityApplyLabelAction{}

type SecurityApplyLabelAction struct {
	ActionSource *SecurityActionSource `json:"actionSource,omitempty"`

	// The collection of actions that should be implemented by the caller.
	Actions *[]SecurityInformationProtectionAction `json:"actions,omitempty"`

	// If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that
	// resulted in the returned label.
	ResponsibleSensitiveTypeIds *[]string `json:"responsibleSensitiveTypeIds,omitempty"`

	SensitivityLabelId nullable.Type[string] `json:"sensitivityLabelId,omitempty"`

	// Fields inherited from SecurityInformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityApplyLabelAction) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return BaseSecurityInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityApplyLabelAction{}

func (s SecurityApplyLabelAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityApplyLabelAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityApplyLabelAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityApplyLabelAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.applyLabelAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityApplyLabelAction: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityApplyLabelAction{}

func (s *SecurityApplyLabelAction) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActionSource                *SecurityActionSource `json:"actionSource,omitempty"`
		ResponsibleSensitiveTypeIds *[]string             `json:"responsibleSensitiveTypeIds,omitempty"`
		SensitivityLabelId          nullable.Type[string] `json:"sensitivityLabelId,omitempty"`
		ODataId                     *string               `json:"@odata.id,omitempty"`
		ODataType                   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActionSource = decoded.ActionSource
	s.ResponsibleSensitiveTypeIds = decoded.ResponsibleSensitiveTypeIds
	s.SensitivityLabelId = decoded.SensitivityLabelId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityApplyLabelAction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Actions into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityInformationProtectionAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityInformationProtectionActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Actions' for 'SecurityApplyLabelAction': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Actions = &output
	}

	return nil
}
