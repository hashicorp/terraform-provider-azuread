package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedApprovalStage struct {
	// The number of days that a request can be pending a response before it is automatically denied.
	ApprovalStageTimeOutInDays nullable.Type[int64] `json:"approvalStageTimeOutInDays,omitempty"`

	// The escalation approvers for this stage when the primary approvers don't respond.
	EscalationApprovers *[]SubjectSet `json:"escalationApprovers,omitempty"`

	// The time a request can be pending a response from a primary approver before it can be escalated to the escalation
	// approvers.
	EscalationTimeInMinutes nullable.Type[int64] `json:"escalationTimeInMinutes,omitempty"`

	// Indicates whether the approver must provide justification for their reponse.
	IsApproverJustificationRequired nullable.Type[bool] `json:"isApproverJustificationRequired,omitempty"`

	// Indicates whether escalation if enabled.
	IsEscalationEnabled nullable.Type[bool] `json:"isEscalationEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The primary approvers of this stage.
	PrimaryApprovers *[]SubjectSet `json:"primaryApprovers,omitempty"`
}

var _ json.Unmarshaler = &UnifiedApprovalStage{}

func (s *UnifiedApprovalStage) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling UnifiedApprovalStage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["escalationApprovers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EscalationApprovers into list []json.RawMessage: %+v", err)
		}

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EscalationApprovers' for 'UnifiedApprovalStage': %+v", i, err)
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

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PrimaryApprovers' for 'UnifiedApprovalStage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PrimaryApprovers = &output
	}

	return nil
}
