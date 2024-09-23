package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsCredentialUserRegistrationsSummary{}

type ManagedTenantsCredentialUserRegistrationsSummary struct {
	// Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The number of users that are capable of performing multi-factor authentication or self service password reset.
	// Optional. Read-only.
	MfaAndSsprCapableUserCount nullable.Type[int64] `json:"mfaAndSsprCapableUserCount,omitempty"`

	// The state of a conditional access policy that enforces multi-factor authentication. Optional. Read-only.
	MfaConditionalAccessPolicyState nullable.Type[string] `json:"mfaConditionalAccessPolicyState,omitempty"`

	// The number of users in the multi-factor authentication exclusion security group (Microsoft 365 Lighthouse - MFA
	// exclusions). Optional. Read-only.
	MfaExcludedUserCount nullable.Type[int64] `json:"mfaExcludedUserCount,omitempty"`

	// The number of users registered for multi-factor authentication. Optional. Read-only.
	MfaRegisteredUserCount nullable.Type[int64] `json:"mfaRegisteredUserCount,omitempty"`

	// A flag indicating whether Identity Security Defaults is enabled. Optional. Read-only.
	SecurityDefaultsEnabled nullable.Type[bool] `json:"securityDefaultsEnabled,omitempty"`

	// The number of users enabled for self service password reset. Optional. Read-only.
	SsprEnabledUserCount nullable.Type[int64] `json:"ssprEnabledUserCount,omitempty"`

	// The number of users registered for self service password reset. Optional. Read-only.
	SsprRegisteredUserCount nullable.Type[int64] `json:"ssprRegisteredUserCount,omitempty"`

	// The display name for the managed tenant. Required. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Required. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The license type associated with the tenant; for example, AADFree, AADPremium1, AADPremium2.
	TenantLicenseType nullable.Type[string] `json:"tenantLicenseType,omitempty"`

	// The total number of users in the given managed tenant. Optional. Read-only.
	TotalUserCount nullable.Type[int64] `json:"totalUserCount,omitempty"`

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

func (s ManagedTenantsCredentialUserRegistrationsSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsCredentialUserRegistrationsSummary{}

func (s ManagedTenantsCredentialUserRegistrationsSummary) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsCredentialUserRegistrationsSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsCredentialUserRegistrationsSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsCredentialUserRegistrationsSummary: %+v", err)
	}

	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "mfaAndSsprCapableUserCount")
	delete(decoded, "mfaConditionalAccessPolicyState")
	delete(decoded, "mfaExcludedUserCount")
	delete(decoded, "mfaRegisteredUserCount")
	delete(decoded, "securityDefaultsEnabled")
	delete(decoded, "ssprEnabledUserCount")
	delete(decoded, "ssprRegisteredUserCount")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")
	delete(decoded, "totalUserCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.credentialUserRegistrationsSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsCredentialUserRegistrationsSummary: %+v", err)
	}

	return encoded, nil
}
