package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsAggregatedPolicyCompliance{}

type ManagedTenantsAggregatedPolicyCompliance struct {
	// Identifier for the device compliance policy. Optional. Read-only.
	CompliancePolicyId nullable.Type[string] `json:"compliancePolicyId,omitempty"`

	// Name of the device compliance policy. Optional. Read-only.
	CompliancePolicyName nullable.Type[string] `json:"compliancePolicyName,omitempty"`

	// Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81,
	// windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
	CompliancePolicyPlatform nullable.Type[string] `json:"compliancePolicyPlatform,omitempty"`

	// The type of compliance policy. Optional. Read-only.
	CompliancePolicyType nullable.Type[string] `json:"compliancePolicyType,omitempty"`

	// Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The number of devices that are in a compliant status. Optional. Read-only.
	NumberOfCompliantDevices nullable.Type[int64] `json:"numberOfCompliantDevices,omitempty"`

	// The number of devices that are in an error status. Optional. Read-only.
	NumberOfErrorDevices nullable.Type[int64] `json:"numberOfErrorDevices,omitempty"`

	// The number of device that are in a non-compliant status. Optional. Read-only.
	NumberOfNonCompliantDevices nullable.Type[int64] `json:"numberOfNonCompliantDevices,omitempty"`

	// The date and time the device policy was last modified. Optional. Read-only.
	PolicyModifiedDateTime nullable.Type[string] `json:"policyModifiedDateTime,omitempty"`

	// The display name for the managed tenant. Optional. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

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

func (s ManagedTenantsAggregatedPolicyCompliance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsAggregatedPolicyCompliance{}

func (s ManagedTenantsAggregatedPolicyCompliance) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsAggregatedPolicyCompliance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsAggregatedPolicyCompliance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsAggregatedPolicyCompliance: %+v", err)
	}

	delete(decoded, "compliancePolicyId")
	delete(decoded, "compliancePolicyName")
	delete(decoded, "compliancePolicyPlatform")
	delete(decoded, "compliancePolicyType")
	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "numberOfCompliantDevices")
	delete(decoded, "numberOfErrorDevices")
	delete(decoded, "numberOfNonCompliantDevices")
	delete(decoded, "policyModifiedDateTime")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.aggregatedPolicyCompliance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsAggregatedPolicyCompliance: %+v", err)
	}

	return encoded, nil
}
