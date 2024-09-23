package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OperationApprovalRequest{}

type OperationApprovalRequest struct {
	// Indicates the justification for approving or rejecting the request. Maximum length of justification is 1024
	// characters. For example: 'Approved per Change 23423 - needed for Feb 2023 application baseline updates.' Read-only.
	// This property is read-only.
	ApprovalJustification nullable.Type[string] `json:"approvalJustification,omitempty"`

	// The identity of the approver as an Identity Set. Optionally contains the application ID, the device ID and the User
	// ID. See information about this type here:
	// https://learn.microsoft.com/graph/api/resources/identityset?view=graph-rest-1.0. Read-only. This property is
	// read-only.
	Approver *IdentitySet `json:"approver,omitempty"`

	// Indicates the DateTime when any action on the approval request is no longer permitted. The value cannot be modified
	// and is automatically populated when the request is created using expiration offset values defined in the service
	// controllers. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default.
	// Read-only. This property is read-only.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Indicates the last DateTime that the request was modified. The value cannot be modified and is automatically
	// populated whenever values in the request are updated. For example, when the 'status' property changes from
	// needsApproval to approved. The Timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned
	// by default. Read-only. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Indicates the DateTime that the request was made. The value cannot be modified and is automatically populated when
	// the request is created. The Timestamp type represents date and time information using ISO 8601 format and is always
	// in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by
	// default. Read-only. This property is read-only.
	RequestDateTime *string `json:"requestDateTime,omitempty"`

	// Indicates the justification for creating the request. Maximum length of justification is 1024 characters. For
	// example: 'Needed for Feb 2023 application baseline updates.' Read-only. This property is read-only.
	RequestJustification nullable.Type[string] `json:"requestJustification,omitempty"`

	// The identity of the requestor as an Identity Set. Optionally contains the application ID, the device ID and the User
	// ID. See information about this type here:
	// https://learn.microsoft.com/graph/api/resources/identityset?view=graph-rest-1.0. Read-only. This property is
	// read-only.
	Requestor *IdentitySet `json:"requestor,omitempty"`

	// Indicates the approval policy types required by the request in order for the request to be approved or rejected.
	// Read-only. This property is read-only.
	RequiredOperationApprovalPolicyTypes *[]OperationApprovalPolicyType `json:"requiredOperationApprovalPolicyTypes,omitempty"`

	// Indicates the status of the Approval Request. The status of a request will change when an action is successfully
	// performed on it, such as when it is `approved` or `rejected`, or when the request's expiration DateTime passes and
	// the result is `expired`.
	Status *OperationApprovalRequestStatus `json:"status,omitempty"`

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

func (s OperationApprovalRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OperationApprovalRequest{}

func (s OperationApprovalRequest) MarshalJSON() ([]byte, error) {
	type wrapper OperationApprovalRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OperationApprovalRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationApprovalRequest: %+v", err)
	}

	delete(decoded, "approvalJustification")
	delete(decoded, "approver")
	delete(decoded, "expirationDateTime")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "requestDateTime")
	delete(decoded, "requestJustification")
	delete(decoded, "requestor")
	delete(decoded, "requiredOperationApprovalPolicyTypes")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.operationApprovalRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OperationApprovalRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OperationApprovalRequest{}

func (s *OperationApprovalRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApprovalJustification                nullable.Type[string]           `json:"approvalJustification,omitempty"`
		ExpirationDateTime                   *string                         `json:"expirationDateTime,omitempty"`
		LastModifiedDateTime                 *string                         `json:"lastModifiedDateTime,omitempty"`
		RequestDateTime                      *string                         `json:"requestDateTime,omitempty"`
		RequestJustification                 nullable.Type[string]           `json:"requestJustification,omitempty"`
		RequiredOperationApprovalPolicyTypes *[]OperationApprovalPolicyType  `json:"requiredOperationApprovalPolicyTypes,omitempty"`
		Status                               *OperationApprovalRequestStatus `json:"status,omitempty"`
		Id                                   *string                         `json:"id,omitempty"`
		ODataId                              *string                         `json:"@odata.id,omitempty"`
		ODataType                            *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApprovalJustification = decoded.ApprovalJustification
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.RequestDateTime = decoded.RequestDateTime
	s.RequestJustification = decoded.RequestJustification
	s.RequiredOperationApprovalPolicyTypes = decoded.RequiredOperationApprovalPolicyTypes
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OperationApprovalRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["approver"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Approver' for 'OperationApprovalRequest': %+v", err)
		}
		s.Approver = &impl
	}

	if v, ok := temp["requestor"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Requestor' for 'OperationApprovalRequest': %+v", err)
		}
		s.Requestor = &impl
	}

	return nil
}
