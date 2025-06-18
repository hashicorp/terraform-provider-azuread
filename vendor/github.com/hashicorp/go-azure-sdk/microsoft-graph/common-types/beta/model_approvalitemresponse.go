package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApprovalItemResponse{}

type ApprovalItemResponse struct {
	// The comment made by the approver.
	Comments nullable.Type[string] `json:"comments,omitempty"`

	// The identity set of the approver.
	CreatedBy *ApprovalIdentitySet `json:"createdBy,omitempty"`

	// Creation date and time of the response.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The identity set of the principal who owns the approval item.
	Owners *[]ApprovalIdentitySet `json:"owners,omitempty"`

	// Approver response based on the response options. The default response options are 'Approved' and 'Rejected'. The
	// approval item creator can also define custom response options during approval item creation.
	Response nullable.Type[string] `json:"response,omitempty"`

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

func (s ApprovalItemResponse) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApprovalItemResponse{}

func (s ApprovalItemResponse) MarshalJSON() ([]byte, error) {
	type wrapper ApprovalItemResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApprovalItemResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApprovalItemResponse: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "owners")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.approvalItemResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApprovalItemResponse: %+v", err)
	}

	return encoded, nil
}
