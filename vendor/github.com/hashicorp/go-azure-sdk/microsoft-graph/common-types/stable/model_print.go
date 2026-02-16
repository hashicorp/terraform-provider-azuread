package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Print struct {
	// The list of available print connectors.
	Connectors *[]PrintConnector `json:"connectors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of print long running operations.
	Operations *[]PrintOperation `json:"operations,omitempty"`

	// The list of printers registered in the tenant.
	Printers *[]Printer `json:"printers,omitempty"`

	// The list of available Universal Print service endpoints.
	Services *[]PrintService `json:"services,omitempty"`

	// Tenant-wide settings for the Universal Print service.
	Settings *PrintSettings `json:"settings,omitempty"`

	// The list of printer shares registered in the tenant.
	Shares *[]PrinterShare `json:"shares,omitempty"`

	// List of abstract definition for a task that can be triggered when various events occur within Universal Print.
	TaskDefinitions *[]PrintTaskDefinition `json:"taskDefinitions,omitempty"`
}

var _ json.Unmarshaler = &Print{}

func (s *Print) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Connectors      *[]PrintConnector      `json:"connectors,omitempty"`
		ODataId         *string                `json:"@odata.id,omitempty"`
		ODataType       *string                `json:"@odata.type,omitempty"`
		Printers        *[]Printer             `json:"printers,omitempty"`
		Services        *[]PrintService        `json:"services,omitempty"`
		Settings        *PrintSettings         `json:"settings,omitempty"`
		Shares          *[]PrinterShare        `json:"shares,omitempty"`
		TaskDefinitions *[]PrintTaskDefinition `json:"taskDefinitions,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Connectors = decoded.Connectors
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Printers = decoded.Printers
	s.Services = decoded.Services
	s.Settings = decoded.Settings
	s.Shares = decoded.Shares
	s.TaskDefinitions = decoded.TaskDefinitions

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Print into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]PrintOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPrintOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'Print': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}
