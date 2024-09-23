package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SensitivityLabel{}

type SensitivityLabel struct {
	ApplicableTo                *SensitivityLabelTarget `json:"applicableTo,omitempty"`
	ApplicationMode             *ApplicationMode        `json:"applicationMode,omitempty"`
	AssignedPolicies            *[]LabelPolicy          `json:"assignedPolicies,omitempty"`
	AutoLabeling                *AutoLabeling           `json:"autoLabeling,omitempty"`
	Description                 nullable.Type[string]   `json:"description,omitempty"`
	DisplayName                 nullable.Type[string]   `json:"displayName,omitempty"`
	IsDefault                   nullable.Type[bool]     `json:"isDefault,omitempty"`
	IsEndpointProtectionEnabled nullable.Type[bool]     `json:"isEndpointProtectionEnabled,omitempty"`
	LabelActions                *[]LabelActionBase      `json:"labelActions,omitempty"`
	Name                        nullable.Type[string]   `json:"name,omitempty"`
	Priority                    nullable.Type[int64]    `json:"priority,omitempty"`
	Sublabels                   *[]SensitivityLabel     `json:"sublabels,omitempty"`
	ToolTip                     nullable.Type[string]   `json:"toolTip,omitempty"`

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

func (s SensitivityLabel) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SensitivityLabel{}

func (s SensitivityLabel) MarshalJSON() ([]byte, error) {
	type wrapper SensitivityLabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SensitivityLabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SensitivityLabel: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sensitivityLabel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SensitivityLabel: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SensitivityLabel{}

func (s *SensitivityLabel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicableTo                *SensitivityLabelTarget `json:"applicableTo,omitempty"`
		ApplicationMode             *ApplicationMode        `json:"applicationMode,omitempty"`
		AssignedPolicies            *[]LabelPolicy          `json:"assignedPolicies,omitempty"`
		AutoLabeling                *AutoLabeling           `json:"autoLabeling,omitempty"`
		Description                 nullable.Type[string]   `json:"description,omitempty"`
		DisplayName                 nullable.Type[string]   `json:"displayName,omitempty"`
		IsDefault                   nullable.Type[bool]     `json:"isDefault,omitempty"`
		IsEndpointProtectionEnabled nullable.Type[bool]     `json:"isEndpointProtectionEnabled,omitempty"`
		Name                        nullable.Type[string]   `json:"name,omitempty"`
		Priority                    nullable.Type[int64]    `json:"priority,omitempty"`
		Sublabels                   *[]SensitivityLabel     `json:"sublabels,omitempty"`
		ToolTip                     nullable.Type[string]   `json:"toolTip,omitempty"`
		Id                          *string                 `json:"id,omitempty"`
		ODataId                     *string                 `json:"@odata.id,omitempty"`
		ODataType                   *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplicableTo = decoded.ApplicableTo
	s.ApplicationMode = decoded.ApplicationMode
	s.AssignedPolicies = decoded.AssignedPolicies
	s.AutoLabeling = decoded.AutoLabeling
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsDefault = decoded.IsDefault
	s.IsEndpointProtectionEnabled = decoded.IsEndpointProtectionEnabled
	s.Name = decoded.Name
	s.Priority = decoded.Priority
	s.Sublabels = decoded.Sublabels
	s.ToolTip = decoded.ToolTip
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SensitivityLabel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["labelActions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling LabelActions into list []json.RawMessage: %+v", err)
		}

		output := make([]LabelActionBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLabelActionBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'LabelActions' for 'SensitivityLabel': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.LabelActions = &output
	}

	return nil
}
