package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantRelationship struct {
	// The customer who has a delegated admin relationship with a Microsoft partner.
	DelegatedAdminCustomers *[]DelegatedAdminCustomer `json:"delegatedAdminCustomers,omitempty"`

	// The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
	DelegatedAdminRelationships *[]DelegatedAdminRelationship `json:"delegatedAdminRelationships,omitempty"`

	// The operations available to interact with the multi-tenant management platform.
	ManagedTenants *ManagedTenantsManagedTenant `json:"managedTenants,omitempty"`

	// Defines an organization with more than one instance of Microsoft Entra ID.
	MultiTenantOrganization *MultiTenantOrganization `json:"multiTenantOrganization,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &TenantRelationship{}

func (s *TenantRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DelegatedAdminCustomers *[]DelegatedAdminCustomer    `json:"delegatedAdminCustomers,omitempty"`
		ManagedTenants          *ManagedTenantsManagedTenant `json:"managedTenants,omitempty"`
		MultiTenantOrganization *MultiTenantOrganization     `json:"multiTenantOrganization,omitempty"`
		ODataId                 *string                      `json:"@odata.id,omitempty"`
		ODataType               *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DelegatedAdminCustomers = decoded.DelegatedAdminCustomers
	s.ManagedTenants = decoded.ManagedTenants
	s.MultiTenantOrganization = decoded.MultiTenantOrganization
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TenantRelationship into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["delegatedAdminRelationships"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DelegatedAdminRelationships into list []json.RawMessage: %+v", err)
		}

		output := make([]DelegatedAdminRelationship, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDelegatedAdminRelationshipImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DelegatedAdminRelationships' for 'TenantRelationship': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DelegatedAdminRelationships = &output
	}

	return nil
}
