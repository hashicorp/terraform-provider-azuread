package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpAssociatedIdentities struct {
	All *[]GcpIdentity `json:"all,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ServiceAccounts *[]GcpServiceAccount `json:"serviceAccounts,omitempty"`
	Users           *[]GcpUser           `json:"users,omitempty"`
}

var _ json.Unmarshaler = &GcpAssociatedIdentities{}

func (s *GcpAssociatedIdentities) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId         *string              `json:"@odata.id,omitempty"`
		ODataType       *string              `json:"@odata.type,omitempty"`
		ServiceAccounts *[]GcpServiceAccount `json:"serviceAccounts,omitempty"`
		Users           *[]GcpUser           `json:"users,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ServiceAccounts = decoded.ServiceAccounts
	s.Users = decoded.Users

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GcpAssociatedIdentities into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["all"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling All into list []json.RawMessage: %+v", err)
		}

		output := make([]GcpIdentity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGcpIdentityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'All' for 'GcpAssociatedIdentities': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.All = &output
	}

	return nil
}
