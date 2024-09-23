package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadStatus struct {
	// The display name for the workload. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time the workload was offboarded. Optional. Read-only.
	OffboardedDateTime nullable.Type[string] `json:"offboardedDateTime,omitempty"`

	// The date and time the workload was onboarded. Optional. Read-only.
	OnboardedDateTime nullable.Type[string] `json:"onboardedDateTime,omitempty"`

	OnboardingStatus *ManagedTenantsWorkloadOnboardingStatus `json:"onboardingStatus,omitempty"`
}

var _ json.Marshaler = ManagedTenantsWorkloadStatus{}

func (s ManagedTenantsWorkloadStatus) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsWorkloadStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsWorkloadStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsWorkloadStatus: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "offboardedDateTime")
	delete(decoded, "onboardedDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsWorkloadStatus: %+v", err)
	}

	return encoded, nil
}
