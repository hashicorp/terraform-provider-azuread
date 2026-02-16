package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberTransitionDetails struct {
	// Role of the tenant in the multitenant organization. The possible values are: owner, member, unknownFutureValue.
	DesiredRole *MultiTenantOrganizationMemberRole `json:"desiredRole,omitempty"`

	// State of the tenant in the multitenant organization currently being processed. The possible values are: pending,
	// active, removed, unknownFutureValue. Read-only.
	DesiredState *MultiTenantOrganizationMemberState `json:"desiredState,omitempty"`

	// Details that explain the processing status if any. Read-only.
	Details nullable.Type[string] `json:"details,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed,
	// unknownFutureValue. Read-only.
	Status *MultiTenantOrganizationMemberProcessingStatus `json:"status,omitempty"`
}

var _ json.Marshaler = MultiTenantOrganizationMemberTransitionDetails{}

func (s MultiTenantOrganizationMemberTransitionDetails) MarshalJSON() ([]byte, error) {
	type wrapper MultiTenantOrganizationMemberTransitionDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiTenantOrganizationMemberTransitionDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiTenantOrganizationMemberTransitionDetails: %+v", err)
	}

	delete(decoded, "desiredState")
	delete(decoded, "details")
	delete(decoded, "status")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiTenantOrganizationMemberTransitionDetails: %+v", err)
	}

	return encoded, nil
}
