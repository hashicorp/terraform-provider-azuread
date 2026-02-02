package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAssociatedIdentities struct {
	All               *[]AzureIdentity        `json:"all,omitempty"`
	ManagedIdentities *[]AzureManagedIdentity `json:"managedIdentities,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ServicePrincipals *[]AzureServicePrincipal `json:"servicePrincipals,omitempty"`
	Users             *[]AzureUser             `json:"users,omitempty"`
}

var _ json.Unmarshaler = &AzureAssociatedIdentities{}

func (s *AzureAssociatedIdentities) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ManagedIdentities *[]AzureManagedIdentity  `json:"managedIdentities,omitempty"`
		ODataId           *string                  `json:"@odata.id,omitempty"`
		ODataType         *string                  `json:"@odata.type,omitempty"`
		ServicePrincipals *[]AzureServicePrincipal `json:"servicePrincipals,omitempty"`
		Users             *[]AzureUser             `json:"users,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ManagedIdentities = decoded.ManagedIdentities
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ServicePrincipals = decoded.ServicePrincipals
	s.Users = decoded.Users

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AzureAssociatedIdentities into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["all"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling All into list []json.RawMessage: %+v", err)
		}

		output := make([]AzureIdentity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAzureIdentityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'All' for 'AzureAssociatedIdentities': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.All = &output
	}

	return nil
}
