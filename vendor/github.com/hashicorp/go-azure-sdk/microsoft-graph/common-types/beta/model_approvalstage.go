package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalStage struct {
	// The number of days that a request can be pending a response before it is automatically denied.
	ApprovalStageTimeOutInDays nullable.Type[int64] `json:"approvalStageTimeOutInDays,omitempty"`

	// The users who are asked to approve requests if escalation is enabled and the primary approvers don't respond before
	// the escalation time. This property can be a collection of singleUser, groupMembers, requestorManager,
	// internalSponsors, and externalSponsors. When you create or update a policy, if there are no escalation approvers, or
	// escalation approvers aren't required for the stage, assign an empty collection to this property.
	EscalationApprovers *[]UserSet `json:"escalationApprovers,omitempty"`

	// If escalation is required, the time a request can be pending a response from a primary approver.
	EscalationTimeInMinutes nullable.Type[int64] `json:"escalationTimeInMinutes,omitempty"`

	// Indicates whether the approver is required to provide a justification for approving a request.
	IsApproverJustificationRequired nullable.Type[bool] `json:"isApproverJustificationRequired,omitempty"`

	// If true, then one or more escalation approvers are configured in this approval stage.
	IsEscalationEnabled nullable.Type[bool] `json:"isEscalationEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The users who are asked to approve requests. A collection of singleUser, groupMembers, requestorManager,
	// internalSponsors, externalSponsors, and targetUserSponsors. When creating or updating a policy, include at least one
	// userSet in this collection.
	PrimaryApprovers *[]UserSet `json:"primaryApprovers,omitempty"`
}

var _ json.Unmarshaler = &ApprovalStage{}

func (s *ApprovalStage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApprovalStageTimeOutInDays      nullable.Type[int64] `json:"approvalStageTimeOutInDays,omitempty"`
		EscalationTimeInMinutes         nullable.Type[int64] `json:"escalationTimeInMinutes,omitempty"`
		IsApproverJustificationRequired nullable.Type[bool]  `json:"isApproverJustificationRequired,omitempty"`
		IsEscalationEnabled             nullable.Type[bool]  `json:"isEscalationEnabled,omitempty"`
		ODataId                         *string              `json:"@odata.id,omitempty"`
		ODataType                       *string              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApprovalStageTimeOutInDays = decoded.ApprovalStageTimeOutInDays
	s.EscalationTimeInMinutes = decoded.EscalationTimeInMinutes
	s.IsApproverJustificationRequired = decoded.IsApproverJustificationRequired
	s.IsEscalationEnabled = decoded.IsEscalationEnabled
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ApprovalStage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["escalationApprovers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EscalationApprovers into list []json.RawMessage: %+v", err)
		}

		output := make([]UserSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUserSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EscalationApprovers' for 'ApprovalStage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EscalationApprovers = &output
	}

	if v, ok := temp["primaryApprovers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PrimaryApprovers into list []json.RawMessage: %+v", err)
		}

		output := make([]UserSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUserSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PrimaryApprovers' for 'ApprovalStage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PrimaryApprovers = &output
	}

	return nil
}
