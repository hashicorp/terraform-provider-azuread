package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingLabel struct {
	ApplicationMode             *ApplicationMode      `json:"applicationMode,omitempty"`
	Description                 nullable.Type[string] `json:"description,omitempty"`
	DisplayName                 nullable.Type[string] `json:"displayName,omitempty"`
	Id                          nullable.Type[string] `json:"id,omitempty"`
	IsEndpointProtectionEnabled nullable.Type[bool]   `json:"isEndpointProtectionEnabled,omitempty"`
	LabelActions                *[]LabelActionBase    `json:"labelActions,omitempty"`
	Name                        nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PolicyTip nullable.Type[string] `json:"policyTip,omitempty"`
	Priority  nullable.Type[int64]  `json:"priority,omitempty"`
	ToolTip   nullable.Type[string] `json:"toolTip,omitempty"`
}

var _ json.Unmarshaler = &MatchingLabel{}

func (s *MatchingLabel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicationMode             *ApplicationMode      `json:"applicationMode,omitempty"`
		Description                 nullable.Type[string] `json:"description,omitempty"`
		DisplayName                 nullable.Type[string] `json:"displayName,omitempty"`
		Id                          nullable.Type[string] `json:"id,omitempty"`
		IsEndpointProtectionEnabled nullable.Type[bool]   `json:"isEndpointProtectionEnabled,omitempty"`
		Name                        nullable.Type[string] `json:"name,omitempty"`
		ODataId                     *string               `json:"@odata.id,omitempty"`
		ODataType                   *string               `json:"@odata.type,omitempty"`
		PolicyTip                   nullable.Type[string] `json:"policyTip,omitempty"`
		Priority                    nullable.Type[int64]  `json:"priority,omitempty"`
		ToolTip                     nullable.Type[string] `json:"toolTip,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplicationMode = decoded.ApplicationMode
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.IsEndpointProtectionEnabled = decoded.IsEndpointProtectionEnabled
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PolicyTip = decoded.PolicyTip
	s.Priority = decoded.Priority
	s.ToolTip = decoded.ToolTip

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MatchingLabel into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'LabelActions' for 'MatchingLabel': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.LabelActions = &output
	}

	return nil
}
