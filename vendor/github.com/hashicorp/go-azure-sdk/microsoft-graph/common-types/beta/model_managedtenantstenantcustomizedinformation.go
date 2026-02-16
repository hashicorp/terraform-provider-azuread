package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsTenantCustomizedInformation{}

type ManagedTenantsTenantCustomizedInformation struct {
	// Describes the relationship between the Managed Services Provider and the managed tenant; for example, Managed,
	// Co-managed, Licensing. The maximum length is 250 characters. Optional.
	BusinessRelationship nullable.Type[string] `json:"businessRelationship,omitempty"`

	// Contains the compliance requirements for the customer tenant; for example, HIPPA, NIST, CMMC. The maximum length is
	// 250 characters per compliance requirement. Optional.
	ComplianceRequirements *[]string `json:"complianceRequirements,omitempty"`

	// The collection of contacts for the managed tenant. Optional.
	Contacts *[]ManagedTenantsTenantContactInformation `json:"contacts,omitempty"`

	// The display name for the managed tenant. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// This is the Managed Services Plans for the customer tenant that the Managed Services Provider manages. The maximum
	// length is 250 characters per managed service plan. Optional.
	ManagedServicesPlans *[]string `json:"managedServicesPlans,omitempty"`

	// A field for the Managed Services Provider technician to input custom text to share notes between technicians within
	// the Managed Service Providers. The maximum length is 5000 characters. Optional.
	Note nullable.Type[string] `json:"note,omitempty"`

	// The date on which the note field of this entity was last modified. Optional.
	NoteLastModifiedDateTime nullable.Type[string] `json:"noteLastModifiedDateTime,omitempty"`

	// The list of Entra user IDs for users in the Managed Services Provider that manage the relationship with the managed
	// tenant. Optional.
	PartnerRelationshipManagerUserIds *[]string `json:"partnerRelationshipManagerUserIds,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
	TenantId *string `json:"tenantId,omitempty"`

	// The website for the managed tenant. Required.
	Website nullable.Type[string] `json:"website,omitempty"`

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

func (s ManagedTenantsTenantCustomizedInformation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsTenantCustomizedInformation{}

func (s ManagedTenantsTenantCustomizedInformation) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTenantCustomizedInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTenantCustomizedInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTenantCustomizedInformation: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.tenantCustomizedInformation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTenantCustomizedInformation: %+v", err)
	}

	return encoded, nil
}
