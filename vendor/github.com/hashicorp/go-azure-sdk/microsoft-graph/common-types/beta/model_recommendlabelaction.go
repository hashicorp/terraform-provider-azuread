package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InformationProtectionAction = RecommendLabelAction{}

type RecommendLabelAction struct {
	ActionSource *ActionSource `json:"actionSource,omitempty"`

	// Actions to take if the label is accepted by the user.
	Actions *[]InformationProtectionAction `json:"actions,omitempty"`

	// The label that is being recommended.
	Label *LabelDetails `json:"label,omitempty"`

	// The sensitive information type GUIDs that caused the recommendation to be given.
	ResponsibleSensitiveTypeIds *[]string `json:"responsibleSensitiveTypeIds,omitempty"`

	// Fields inherited from InformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s RecommendLabelAction) InformationProtectionAction() BaseInformationProtectionActionImpl {
	return BaseInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RecommendLabelAction{}

func (s RecommendLabelAction) MarshalJSON() ([]byte, error) {
	type wrapper RecommendLabelAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RecommendLabelAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RecommendLabelAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.recommendLabelAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RecommendLabelAction: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &RecommendLabelAction{}

func (s *RecommendLabelAction) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActionSource                *ActionSource `json:"actionSource,omitempty"`
		Label                       *LabelDetails `json:"label,omitempty"`
		ResponsibleSensitiveTypeIds *[]string     `json:"responsibleSensitiveTypeIds,omitempty"`
		ODataId                     *string       `json:"@odata.id,omitempty"`
		ODataType                   *string       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActionSource = decoded.ActionSource
	s.Label = decoded.Label
	s.ResponsibleSensitiveTypeIds = decoded.ResponsibleSensitiveTypeIds
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RecommendLabelAction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Actions into list []json.RawMessage: %+v", err)
		}

		output := make([]InformationProtectionAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalInformationProtectionActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Actions' for 'RecommendLabelAction': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Actions = &output
	}

	return nil
}
