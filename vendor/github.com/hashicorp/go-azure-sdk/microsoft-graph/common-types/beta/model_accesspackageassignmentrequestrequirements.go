package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequirements struct {
	// Answers that have already been provided.
	ExistingAnswers *[]AccessPackageAnswer `json:"existingAnswers,omitempty"`

	// Indicates whether a request must be approved by an approver.
	IsApprovalRequired nullable.Type[bool] `json:"isApprovalRequired,omitempty"`

	// Indicates whether approval is required when a user tries to extend their access.
	IsApprovalRequiredForExtension nullable.Type[bool] `json:"isApprovalRequiredForExtension,omitempty"`

	// Indicates whether the requestor is allowed to set a custom schedule.
	IsCustomAssignmentScheduleAllowed nullable.Type[bool] `json:"isCustomAssignmentScheduleAllowed,omitempty"`

	// Indicates whether a requestor must supply justification when submitting an assignment request.
	IsRequestorJustificationRequired nullable.Type[bool] `json:"isRequestorJustificationRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The description of the policy that the user is trying to request access using.
	PolicyDescription nullable.Type[string] `json:"policyDescription,omitempty"`

	// The display name of the policy that the user is trying to request access using.
	PolicyDisplayName nullable.Type[string] `json:"policyDisplayName,omitempty"`

	// The identifier of the policy that these requirements are associated with. This identifier can be used when creating a
	// new assignment request.
	PolicyId nullable.Type[string] `json:"policyId,omitempty"`

	// Questions that are configured on the policy. The questions can be required or optional; callers can determine whether
	// a question is required or optional based on the isRequired property on accessPackageQuestion.
	Questions *[]AccessPackageQuestion `json:"questions,omitempty"`

	// Schedule restrictions enforced, if any.
	Schedule *RequestSchedule `json:"schedule,omitempty"`

	// The status of the process to process the verifiable credential, if any.
	VerifiableCredentialRequirementStatus VerifiableCredentialRequirementStatus `json:"verifiableCredentialRequirementStatus"`
}

var _ json.Unmarshaler = &AccessPackageAssignmentRequestRequirements{}

func (s *AccessPackageAssignmentRequestRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsApprovalRequired                nullable.Type[bool]   `json:"isApprovalRequired,omitempty"`
		IsApprovalRequiredForExtension    nullable.Type[bool]   `json:"isApprovalRequiredForExtension,omitempty"`
		IsCustomAssignmentScheduleAllowed nullable.Type[bool]   `json:"isCustomAssignmentScheduleAllowed,omitempty"`
		IsRequestorJustificationRequired  nullable.Type[bool]   `json:"isRequestorJustificationRequired,omitempty"`
		ODataId                           *string               `json:"@odata.id,omitempty"`
		ODataType                         *string               `json:"@odata.type,omitempty"`
		PolicyDescription                 nullable.Type[string] `json:"policyDescription,omitempty"`
		PolicyDisplayName                 nullable.Type[string] `json:"policyDisplayName,omitempty"`
		PolicyId                          nullable.Type[string] `json:"policyId,omitempty"`
		Schedule                          *RequestSchedule      `json:"schedule,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsApprovalRequired = decoded.IsApprovalRequired
	s.IsApprovalRequiredForExtension = decoded.IsApprovalRequiredForExtension
	s.IsCustomAssignmentScheduleAllowed = decoded.IsCustomAssignmentScheduleAllowed
	s.IsRequestorJustificationRequired = decoded.IsRequestorJustificationRequired
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PolicyDescription = decoded.PolicyDescription
	s.PolicyDisplayName = decoded.PolicyDisplayName
	s.PolicyId = decoded.PolicyId
	s.Schedule = decoded.Schedule

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageAssignmentRequestRequirements into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["existingAnswers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ExistingAnswers into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessPackageAnswer, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessPackageAnswerImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ExistingAnswers' for 'AccessPackageAssignmentRequestRequirements': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ExistingAnswers = &output
	}

	if v, ok := temp["questions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Questions into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessPackageQuestion, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessPackageQuestionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Questions' for 'AccessPackageAssignmentRequestRequirements': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Questions = &output
	}

	if v, ok := temp["verifiableCredentialRequirementStatus"]; ok {
		impl, err := UnmarshalVerifiableCredentialRequirementStatusImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'VerifiableCredentialRequirementStatus' for 'AccessPackageAssignmentRequestRequirements': %+v", err)
		}
		s.VerifiableCredentialRequirementStatus = impl
	}

	return nil
}
