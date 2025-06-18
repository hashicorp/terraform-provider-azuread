package beta

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

	// For a requestType of UserAdd or AdminAdd, an access package assignment requested to be created. For a requestType of
	// UserRemove, AdminRemove, or SystemRemove, this property has the id property of an existing assignment to be removed.
	// Supports $expand.
	AccessPackageAssignment *AccessPackageAssignment `json:"accessPackageAssignment,omitempty"`

	// Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
	Answers *[]AccessPackageAnswer `json:"answers,omitempty"`

	// The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Read-only.
	CompletedDate nullable.Type[string] `json:"completedDate,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Information about all the custom extension calls that were made during the access package assignment request
	// workflow.
	CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`

	// A collection of custom workflow extension instances being run on an assignment request. Read-only.
	CustomExtensionHandlerInstances *[]CustomExtensionHandlerInstance `json:"customExtensionHandlerInstances,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	History *[]RequestActivity `json:"history,omitempty"`

	// True if the request isn't to be processed for assignment.
	IsValidationOnly nullable.Type[bool] `json:"isValidationOnly,omitempty"`

	// The requestor's supplied justification.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// One of PendingApproval, Canceled, Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted, or
	// Scheduled. Read-only.
	RequestState nullable.Type[string] `json:"requestState,omitempty"`

	// More information on the request processing status. Read-only.
	RequestStatus nullable.Type[string] `json:"requestStatus,omitempty"`

	// One of UserAdd, UserExtend, UserUpdate, UserRemove, AdminAdd, AdminRemove, or SystemRemove. A request from the user
	// has a requestType of UserAdd, UserUpdate, or UserRemove. Read-only.
	RequestType nullable.Type[string] `json:"requestType,omitempty"`

	// The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
	Requestor *AccessPackageSubject `json:"requestor,omitempty"`

	// The range of dates that access is to be assigned to the requestor. Read-only.
	Schedule *RequestSchedule `json:"schedule,omitempty"`

	// The details of the verifiable credential that the requestor presented, such as the issuer and claims. Read-only.
	VerifiedCredentialsData *[]VerifiedCredentialData `json:"verifiedCredentialsData,omitempty"`

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
	delete(decoded, "completedDate")
	delete(decoded, "createdDateTime")
	delete(decoded, "customExtensionHandlerInstances")
	delete(decoded, "requestState")
	delete(decoded, "requestStatus")
	delete(decoded, "requestType")
	delete(decoded, "requestor")
	delete(decoded, "schedule")
	delete(decoded, "verifiedCredentialsData")

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
		AccessPackageAssignment         *AccessPackageAssignment          `json:"accessPackageAssignment,omitempty"`
		CompletedDate                   nullable.Type[string]             `json:"completedDate,omitempty"`
		CreatedDateTime                 nullable.Type[string]             `json:"createdDateTime,omitempty"`
		CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`
		CustomExtensionHandlerInstances *[]CustomExtensionHandlerInstance `json:"customExtensionHandlerInstances,omitempty"`
		ExpirationDateTime              nullable.Type[string]             `json:"expirationDateTime,omitempty"`
		History                         *[]RequestActivity                `json:"history,omitempty"`
		IsValidationOnly                nullable.Type[bool]               `json:"isValidationOnly,omitempty"`
		Justification                   nullable.Type[string]             `json:"justification,omitempty"`
		RequestState                    nullable.Type[string]             `json:"requestState,omitempty"`
		RequestStatus                   nullable.Type[string]             `json:"requestStatus,omitempty"`
		RequestType                     nullable.Type[string]             `json:"requestType,omitempty"`
		Requestor                       *AccessPackageSubject             `json:"requestor,omitempty"`
		Schedule                        *RequestSchedule                  `json:"schedule,omitempty"`
		VerifiedCredentialsData         *[]VerifiedCredentialData         `json:"verifiedCredentialsData,omitempty"`
		Id                              *string                           `json:"id,omitempty"`
		ODataId                         *string                           `json:"@odata.id,omitempty"`
		ODataType                       *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessPackage = decoded.AccessPackage
	s.AccessPackageAssignment = decoded.AccessPackageAssignment
	s.CompletedDate = decoded.CompletedDate
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomExtensionCalloutInstances = decoded.CustomExtensionCalloutInstances
	s.CustomExtensionHandlerInstances = decoded.CustomExtensionHandlerInstances
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.History = decoded.History
	s.IsValidationOnly = decoded.IsValidationOnly
	s.Justification = decoded.Justification
	s.RequestState = decoded.RequestState
	s.RequestStatus = decoded.RequestStatus
	s.RequestType = decoded.RequestType
	s.Requestor = decoded.Requestor
	s.Schedule = decoded.Schedule
	s.VerifiedCredentialsData = decoded.VerifiedCredentialsData
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
