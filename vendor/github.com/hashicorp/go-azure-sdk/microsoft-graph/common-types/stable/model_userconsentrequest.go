package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Request = UserConsentRequest{}

type UserConsentRequest struct {
	// Approval decisions associated with a request.
	Approval *Approval `json:"approval,omitempty"`

	// The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
	Reason nullable.Type[string] `json:"reason,omitempty"`

	// Fields inherited from Request

	// The identifier of the approval of the request.
	ApprovalId nullable.Type[string] `json:"approvalId,omitempty"`

	// The request completion date time.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The principal that created the request.
	CreatedBy IdentitySet `json:"createdBy"`

	// The request creation date time.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Free text field to define any custom data for the request. Not used.
	CustomData nullable.Type[string] `json:"customData,omitempty"`

	// The status of the request. Not nullable. The possible values are: Canceled, Denied, Failed, Granted,
	// PendingAdminDecision, PendingApproval, PendingProvisioning, PendingScheduleCreation, Provisioned, Revoked, and
	// ScheduleCreated. Not nullable.
	Status *string `json:"status,omitempty"`

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

func (s UserConsentRequest) Request() BaseRequestImpl {
	return BaseRequestImpl{
		ApprovalId:        s.ApprovalId,
		CompletedDateTime: s.CompletedDateTime,
		CreatedBy:         s.CreatedBy,
		CreatedDateTime:   s.CreatedDateTime,
		CustomData:        s.CustomData,
		Status:            s.Status,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s UserConsentRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserConsentRequest{}

func (s UserConsentRequest) MarshalJSON() ([]byte, error) {
	type wrapper UserConsentRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserConsentRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserConsentRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userConsentRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserConsentRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UserConsentRequest{}

func (s *UserConsentRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Approval          *Approval             `json:"approval,omitempty"`
		Reason            nullable.Type[string] `json:"reason,omitempty"`
		ApprovalId        nullable.Type[string] `json:"approvalId,omitempty"`
		CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`
		CreatedDateTime   nullable.Type[string] `json:"createdDateTime,omitempty"`
		CustomData        nullable.Type[string] `json:"customData,omitempty"`
		Status            *string               `json:"status,omitempty"`
		Id                *string               `json:"id,omitempty"`
		ODataId           *string               `json:"@odata.id,omitempty"`
		ODataType         *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Approval = decoded.Approval
	s.Reason = decoded.Reason
	s.ApprovalId = decoded.ApprovalId
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomData = decoded.CustomData
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UserConsentRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'UserConsentRequest': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
