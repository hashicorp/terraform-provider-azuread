package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationApprovalRequestEntityStatus struct {
	// The status of the Entity connected to the OperationApprovalRequest in regard to changes, whether further requests are
	// allowed or if the Entity is locked. When true, a lock is present on the Entity and no approval requests can be
	// currently made for it. When false, the Entity is not locked and approval requests are allowed. Default value is
	// false. Read-only. This property is read-only.
	EntityLocked *bool `json:"entityLocked,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the DateTime when any action on the OperationApprovalRequest is no longer permitted. The value cannot be
	// modified and is automatically populated when the request is created using expiration offset values defined in the
	// service controllers. The Timestamp type represents date and time information using ISO 8601 format and is always in
	// UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default.
	// Read-only. This property is read-only.
	RequestExpirationDateTime nullable.Type[string] `json:"requestExpirationDateTime,omitempty"`

	// The unique identifier of the OperationApprovalRequest. This property cannot be modified and is required when the
	// entity status is created. Read-only. This property is read-only.
	RequestId nullable.Type[string] `json:"requestId,omitempty"`

	// Indicates the status of the Approval Request. The status of a request will change when an action is successfully
	// performed on it, such as when it is `approved` or `rejected`, or when the request's expiration DateTime passes and
	// the result is `expired`.
	RequestStatus *OperationApprovalRequestStatus `json:"requestStatus,omitempty"`
}

var _ json.Marshaler = OperationApprovalRequestEntityStatus{}

func (s OperationApprovalRequestEntityStatus) MarshalJSON() ([]byte, error) {
	type wrapper OperationApprovalRequestEntityStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OperationApprovalRequestEntityStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationApprovalRequestEntityStatus: %+v", err)
	}

	delete(decoded, "entityLocked")
	delete(decoded, "requestExpirationDateTime")
	delete(decoded, "requestId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OperationApprovalRequestEntityStatus: %+v", err)
	}

	return encoded, nil
}
