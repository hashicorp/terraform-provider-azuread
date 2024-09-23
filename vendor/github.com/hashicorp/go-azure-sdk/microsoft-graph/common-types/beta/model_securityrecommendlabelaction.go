package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityInformationProtectionAction = SecurityRecommendLabelAction{}

type SecurityRecommendLabelAction struct {
	ActionSource *SecurityActionSource `json:"actionSource,omitempty"`

	// Actions to take if the label is accepted by the user.
	Actions *[]SecurityInformationProtectionAction `json:"actions,omitempty"`

	// The sensitive information type GUIDs that caused the recommendation to be given.
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

func (s SecurityRecommendLabelAction) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return BaseSecurityInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityRecommendLabelAction{}

func (s SecurityRecommendLabelAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityRecommendLabelAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityRecommendLabelAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityRecommendLabelAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.recommendLabelAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityRecommendLabelAction: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityRecommendLabelAction{}

func (s *SecurityRecommendLabelAction) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling SecurityRecommendLabelAction into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Actions' for 'SecurityRecommendLabelAction': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Actions = &output
	}

	return nil
}
