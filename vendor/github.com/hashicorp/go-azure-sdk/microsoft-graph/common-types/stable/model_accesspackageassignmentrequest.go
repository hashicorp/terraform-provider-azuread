package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageAssignmentRequest{}

type AccessPackageAssignmentRequest struct {
	// The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of
	// resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.
	// Supports $expand.
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`

	// Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
	Answers *[]AccessPackageAnswer `json:"answers,omitempty"`

	// For a requestType of userAdd or adminAdd, this is an access package assignment requested to be created. For a
	// requestType of userRemove, adminRemove or systemRemove, this has the id property of an existing assignment to be
	// removed. Supports $expand.
	Assignment *AccessPackageAssignment `json:"assignment,omitempty"`

	// The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Read-only.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Information about all the custom extension calls that were made during the access package assignment workflow.
	CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`

	// The type of the request. The possible values are: notSpecified, userAdd, UserExtend, userUpdate, userRemove,
	// adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd (not supported),
	// unknownFutureValue. Requests from the user have a requestType of userAdd, userUpdate, or userRemove. This property
	// can't be changed once set.
	RequestType *AccessPackageRequestType `json:"requestType,omitempty"`

	// The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
	Requestor *AccessPackageSubject `json:"requestor,omitempty"`

	// The range of dates that access is to be assigned to the requestor. This property can't be changed once set.
	Schedule *EntitlementManagementSchedule `json:"schedule,omitempty"`

	// The state of the request. The possible values are: submitted, pendingApproval, delivering, delivered, deliveryFailed,
	// denied, scheduled, canceled, partiallyDelivered, unknownFutureValue. Read-only. Supports $filter (eq).
	State *AccessPackageRequestState `json:"state,omitempty"`

	// More information on the request processing status. Read-only.
	Status nullable.Type[string] `json:"status,omitempty"`

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

func (s AccessPackageAssignmentRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAssignmentRequest{}

func (s AccessPackageAssignmentRequest) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAssignmentRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAssignmentRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAssignmentRequest: %+v", err)
	}

	delete(decoded, "accessPackage")
	delete(decoded, "completedDateTime")
	delete(decoded, "createdDateTime")
	delete(decoded, "requestor")
	delete(decoded, "state")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAssignmentRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAssignmentRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessPackageAssignmentRequest{}

func (s *AccessPackageAssignmentRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessPackage                   *AccessPackage                    `json:"accessPackage,omitempty"`
		Assignment                      *AccessPackageAssignment          `json:"assignment,omitempty"`
		CompletedDateTime               nullable.Type[string]             `json:"completedDateTime,omitempty"`
		CreatedDateTime                 nullable.Type[string]             `json:"createdDateTime,omitempty"`
		CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`
		RequestType                     *AccessPackageRequestType         `json:"requestType,omitempty"`
		Requestor                       *AccessPackageSubject             `json:"requestor,omitempty"`
		Schedule                        *EntitlementManagementSchedule    `json:"schedule,omitempty"`
		State                           *AccessPackageRequestState        `json:"state,omitempty"`
		Status                          nullable.Type[string]             `json:"status,omitempty"`
		Id                              *string                           `json:"id,omitempty"`
		ODataId                         *string                           `json:"@odata.id,omitempty"`
		ODataType                       *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessPackage = decoded.AccessPackage
	s.Assignment = decoded.Assignment
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomExtensionCalloutInstances = decoded.CustomExtensionCalloutInstances
	s.RequestType = decoded.RequestType
	s.Requestor = decoded.Requestor
	s.Schedule = decoded.Schedule
	s.State = decoded.State
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageAssignmentRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["answers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Answers into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessPackageAnswer, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessPackageAnswerImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Answers' for 'AccessPackageAssignmentRequest': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Answers = &output
	}

	return nil
}
