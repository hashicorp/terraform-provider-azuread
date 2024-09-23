package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequirements struct {
	// Indicates whether the requestor is allowed to set a custom schedule.
	AllowCustomAssignmentSchedule nullable.Type[bool] `json:"allowCustomAssignmentSchedule,omitempty"`

	// Indicates whether a request to add must be approved by an approver.
	IsApprovalRequiredForAdd nullable.Type[bool] `json:"isApprovalRequiredForAdd,omitempty"`

	// Indicates whether a request to update must be approved by an approver.
	IsApprovalRequiredForUpdate nullable.Type[bool] `json:"isApprovalRequiredForUpdate,omitempty"`

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

	Questions *[]AccessPackageQuestion `json:"questions,omitempty"`

	// Schedule restrictions enforced, if any.
	Schedule *EntitlementManagementSchedule `json:"schedule,omitempty"`
}

var _ json.Unmarshaler = &AccessPackageAssignmentRequestRequirements{}

func (s *AccessPackageAssignmentRequestRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowCustomAssignmentSchedule nullable.Type[bool]            `json:"allowCustomAssignmentSchedule,omitempty"`
		IsApprovalRequiredForAdd      nullable.Type[bool]            `json:"isApprovalRequiredForAdd,omitempty"`
		IsApprovalRequiredForUpdate   nullable.Type[bool]            `json:"isApprovalRequiredForUpdate,omitempty"`
		ODataId                       *string                        `json:"@odata.id,omitempty"`
		ODataType                     *string                        `json:"@odata.type,omitempty"`
		PolicyDescription             nullable.Type[string]          `json:"policyDescription,omitempty"`
		PolicyDisplayName             nullable.Type[string]          `json:"policyDisplayName,omitempty"`
		PolicyId                      nullable.Type[string]          `json:"policyId,omitempty"`
		Schedule                      *EntitlementManagementSchedule `json:"schedule,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowCustomAssignmentSchedule = decoded.AllowCustomAssignmentSchedule
	s.IsApprovalRequiredForAdd = decoded.IsApprovalRequiredForAdd
	s.IsApprovalRequiredForUpdate = decoded.IsApprovalRequiredForUpdate
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

	return nil
}
