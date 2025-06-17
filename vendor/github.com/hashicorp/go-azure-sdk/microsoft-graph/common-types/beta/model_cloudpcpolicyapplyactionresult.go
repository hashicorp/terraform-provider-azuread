package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCPolicyApplyActionResult{}

type CloudPCPolicyApplyActionResult struct {
	// The date and time when the operation finished.
	FinishDateTime nullable.Type[string] `json:"finishDateTime,omitempty"`

	// The date and time when the operation was applied.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	Status *CloudPCPolicyApplyActionStatus `json:"status,omitempty"`

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

func (s CloudPCPolicyApplyActionResult) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCPolicyApplyActionResult{}

func (s CloudPCPolicyApplyActionResult) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCPolicyApplyActionResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCPolicyApplyActionResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCPolicyApplyActionResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcPolicyApplyActionResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCPolicyApplyActionResult: %+v", err)
	}

	return encoded, nil
}
