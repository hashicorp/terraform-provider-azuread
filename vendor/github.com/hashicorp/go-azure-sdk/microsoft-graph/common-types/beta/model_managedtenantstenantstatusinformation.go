package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformation struct {
	// The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible
	// values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges,
	// delegatedAndGranularDelegetedAdminPrivileges. Use the Prefer: include-unknown-enum-members request header to get the
	// following values from this evolvable enum: granularDelegatedAdminPrivileges ,
	// delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
	DelegatedPrivilegeStatus *ManagedTenantsDelegatedPrivilegeStatus `json:"delegatedPrivilegeStatus,omitempty"`

	// The date and time the delegated admin privileges status was updated. Optional. Read-only.
	LastDelegatedPrivilegeRefreshDateTime nullable.Type[string] `json:"lastDelegatedPrivilegeRefreshDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier for the account that offboarded the managed tenant. Optional. Read-only.
	OffboardedByUserId nullable.Type[string] `json:"offboardedByUserId,omitempty"`

	// The date and time when the managed tenant was offboarded. Optional. Read-only.
	OffboardedDateTime nullable.Type[string] `json:"offboardedDateTime,omitempty"`

	// The identifier for the account that onboarded the managed tenant. Optional. Read-only.
	OnboardedByUserId nullable.Type[string] `json:"onboardedByUserId,omitempty"`

	// The date and time when the managed tenant was onboarded. Optional. Read-only.
	OnboardedDateTime nullable.Type[string] `json:"onboardedDateTime,omitempty"`

	// The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive,
	// unknownFutureValue. Optional. Read-only.
	OnboardingStatus *ManagedTenantsTenantOnboardingStatus `json:"onboardingStatus,omitempty"`

	// Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType,
	// delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
	TenantOnboardingEligibilityReason *ManagedTenantsTenantOnboardingEligibilityReason `json:"tenantOnboardingEligibilityReason,omitempty"`

	// The collection of workload statues for the managed tenant. Optional. Read-only.
	WorkloadStatuses *[]ManagedTenantsWorkloadStatus `json:"workloadStatuses,omitempty"`
}

var _ json.Marshaler = ManagedTenantsTenantStatusInformation{}

func (s ManagedTenantsTenantStatusInformation) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTenantStatusInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTenantStatusInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTenantStatusInformation: %+v", err)
	}

	delete(decoded, "delegatedPrivilegeStatus")
	delete(decoded, "lastDelegatedPrivilegeRefreshDateTime")
	delete(decoded, "offboardedByUserId")
	delete(decoded, "offboardedDateTime")
	delete(decoded, "onboardedByUserId")
	delete(decoded, "onboardedDateTime")
	delete(decoded, "onboardingStatus")
	delete(decoded, "tenantOnboardingEligibilityReason")
	delete(decoded, "workloadStatuses")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTenantStatusInformation: %+v", err)
	}

	return encoded, nil
}
