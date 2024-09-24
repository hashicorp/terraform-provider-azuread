package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipRuleProcessingStatus struct {
	// Detailed error message if dynamic group processing ran into an error. Optional. Read-only.
	ErrorMessage nullable.Type[string] `json:"errorMessage,omitempty"`

	// Most recent date and time when membership of a dynamic group was updated. Optional. Read-only.
	LastMembershipUpdated nullable.Type[string] `json:"lastMembershipUpdated,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and
	// UnknownFutureValue. Required. Read-only.
	Status *MembershipRuleProcessingStatusDetails `json:"status,omitempty"`
}

var _ json.Marshaler = MembershipRuleProcessingStatus{}

func (s MembershipRuleProcessingStatus) MarshalJSON() ([]byte, error) {
	type wrapper MembershipRuleProcessingStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MembershipRuleProcessingStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MembershipRuleProcessingStatus: %+v", err)
	}

	delete(decoded, "errorMessage")
	delete(decoded, "lastMembershipUpdated")
	delete(decoded, "status")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MembershipRuleProcessingStatus: %+v", err)
	}

	return encoded, nil
}
