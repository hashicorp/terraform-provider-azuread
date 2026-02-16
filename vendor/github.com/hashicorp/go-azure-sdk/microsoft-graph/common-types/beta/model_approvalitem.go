package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApprovalItem{}

type ApprovalItem struct {
	// Indicates whether the approval item can be canceled.
	AllowCancel nullable.Type[bool] `json:"allowCancel,omitempty"`

	// Indicates whether email notification is enabled.
	AllowEmailNotification nullable.Type[bool] `json:"allowEmailNotification,omitempty"`

	// The workflow type of the approval item. The possible values are: basic, basicAwaitAll, custom, customAwaitAll.
	// Required.
	ApprovalType ApprovalItemType `json:"approvalType"`

	// The identity of the principals to whom the approval item was initially assigned. Required.
	Approvers []ApprovalIdentitySet `json:"approvers"`

	// Approval request completion date and time. Read-only.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// Creation date and time of the approval request. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description of the approval request.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The displayName of the approval request. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identity set of the principal who owns the approval item. Only provide a value for this property when creating an
	// approval item on behalf of the principal. If the owner field isn't provided, the user information from the user
	// context is used.
	Owner *ApprovalIdentitySet `json:"owner,omitempty"`

	// A collection of requests created for each approver on the approval item.
	Requests *[]ApprovalItemRequest `json:"requests,omitempty"`

	// Approval response prompts. Only provide a value for this property when creating a custom approval item. For custom
	// approval items, supply two response prompt strings. The default response prompts are 'Approve' and 'Reject'.
	ResponsePrompts *[]string `json:"responsePrompts,omitempty"`

	// A collection of responses created for the approval item.
	Responses *[]ApprovalItemResponse `json:"responses,omitempty"`

	// The result field is only populated once the approval item is in its final state. The result of the approval item is
	// based on the approvalType. For basic approval items, the result is either 'Approved' or 'Rejected'. For custom
	// approval items, the result could either be a single response or multiple responses separated by a semi-colon.
	// Read-only.
	Result nullable.Type[string] `json:"result,omitempty"`

	// The approval item state. The possible values are: canceled, created, pending, completed. Read-only.
	State *ApprovalItemState `json:"state,omitempty"`

	// Represents user viewpoints data on the ApprovalItem. The data includes the users roles regarding the approval item.
	// Read-only.
	ViewPoint *ApprovalItemViewPoint `json:"viewPoint,omitempty"`

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

func (s ApprovalItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApprovalItem{}

func (s ApprovalItem) MarshalJSON() ([]byte, error) {
	type wrapper ApprovalItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApprovalItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApprovalItem: %+v", err)
	}

	delete(decoded, "allowCancel")
	delete(decoded, "completedDateTime")
	delete(decoded, "createdDateTime")
	delete(decoded, "owner")
	delete(decoded, "result")
	delete(decoded, "state")
	delete(decoded, "viewPoint")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.approvalItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApprovalItem: %+v", err)
	}

	return encoded, nil
}
