package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternal struct {
	// Represents an onboarded AWS account, Azure subscription, or GCP project that Microsoft Entra Permissions Management
	// will collect and analyze permissions and actions on.
	AuthorizationSystems *[]AuthorizationSystem `json:"authorizationSystems,omitempty"`

	Connections  *[]ExternalConnectorsExternalConnection `json:"connections,omitempty"`
	IndustryData *IndustryDataIndustryDataRoot           `json:"industryData,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &ExternalConnectorsExternal{}

func (s *ExternalConnectorsExternal) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Connections  *[]ExternalConnectorsExternalConnection `json:"connections,omitempty"`
		IndustryData *IndustryDataIndustryDataRoot           `json:"industryData,omitempty"`
		ODataId      *string                                 `json:"@odata.id,omitempty"`
		ODataType    *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Connections = decoded.Connections
	s.IndustryData = decoded.IndustryData
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExternalConnectorsExternal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystems"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AuthorizationSystems into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthorizationSystem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthorizationSystemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AuthorizationSystems' for 'ExternalConnectorsExternal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AuthorizationSystems = &output
	}

	return nil
}
