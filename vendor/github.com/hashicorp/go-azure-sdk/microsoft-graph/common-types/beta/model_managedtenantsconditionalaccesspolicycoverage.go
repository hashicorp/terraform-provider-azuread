package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsConditionalAccessPolicyCoverage{}

type ManagedTenantsConditionalAccessPolicyCoverage struct {
	// The state for the conditional access policy. Possible values are: enabled, disabled,
	// enabledForReportingButNotEnforced. Required. Read-only.
	ConditionalAccessPolicyState nullable.Type[string] `json:"conditionalAccessPolicyState,omitempty"`

	// The date and time the conditional access policy was last modified. Required. Read-only.
	LatestPolicyModifiedDateTime nullable.Type[string] `json:"latestPolicyModifiedDateTime,omitempty"`

	// A flag indicating whether the conditional access policy requires device compliance. Required. Read-only.
	RequiresDeviceCompliance nullable.Type[bool] `json:"requiresDeviceCompliance,omitempty"`

	// The display name for the managed tenant. Required. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

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

func (s ManagedTenantsConditionalAccessPolicyCoverage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsConditionalAccessPolicyCoverage{}

func (s ManagedTenantsConditionalAccessPolicyCoverage) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsConditionalAccessPolicyCoverage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsConditionalAccessPolicyCoverage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsConditionalAccessPolicyCoverage: %+v", err)
	}

	delete(decoded, "conditionalAccessPolicyState")
	delete(decoded, "latestPolicyModifiedDateTime")
	delete(decoded, "requiresDeviceCompliance")
	delete(decoded, "tenantDisplayName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.conditionalAccessPolicyCoverage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsConditionalAccessPolicyCoverage: %+v", err)
	}

	return encoded, nil
}
