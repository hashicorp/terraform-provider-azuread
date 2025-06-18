package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApprovalItemRequest{}

type ApprovalItemRequest struct {
	// The identity set of the principal assigned to this request.
	Approver *ApprovalIdentitySet `json:"approver,omitempty"`

	// Creation date and time for the request.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates whether a request was reassigned.
	IsReassigned nullable.Type[bool] `json:"isReassigned,omitempty"`

	// The identity set of the principal who reassigned the request.
	ReassignedFrom *ApprovalIdentitySet `json:"reassignedFrom,omitempty"`

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

func (s ApprovalItemRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApprovalItemRequest{}

func (s ApprovalItemRequest) MarshalJSON() ([]byte, error) {
	type wrapper ApprovalItemRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApprovalItemRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApprovalItemRequest: %+v", err)
	}

	delete(decoded, "approver")
	delete(decoded, "createdDateTime")
	delete(decoded, "isReassigned")
	delete(decoded, "reassignedFrom")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.approvalItemRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApprovalItemRequest: %+v", err)
	}

	return encoded, nil
}
