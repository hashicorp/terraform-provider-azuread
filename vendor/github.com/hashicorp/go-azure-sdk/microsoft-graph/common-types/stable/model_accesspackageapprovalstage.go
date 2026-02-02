package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageApprovalStage struct {
	// The number of days that a request can be pending a response before it is automatically denied.
	DurationBeforeAutomaticDenial nullable.Type[string] `json:"durationBeforeAutomaticDenial,omitempty"`

	// If escalation is required, the time a request can be pending a response from a primary approver.
	DurationBeforeEscalation nullable.Type[string] `json:"durationBeforeEscalation,omitempty"`

	// If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers
	// are the users who will be asked to approve requests.
	EscalationApprovers *[]SubjectSet `json:"escalationApprovers,omitempty"`

	// The subjects, typically users, who are the fallback escalation approvers.
	FallbackEscalationApprovers *[]SubjectSet `json:"fallbackEscalationApprovers,omitempty"`

	// The subjects, typically users, who are the fallback primary approvers.
	FallbackPrimaryApprovers *[]SubjectSet `json:"fallbackPrimaryApprovers,omitempty"`

	// Indicates whether the approver is required to provide a justification for approving a request.
	IsApproverJustificationRequired nullable.Type[bool] `json:"isApproverJustificationRequired,omitempty"`

	// If true, then one or more escalationApprovers are configured in this approval stage.
	IsEscalationEnabled nullable.Type[bool] `json:"isEscalationEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The subjects, typically users, who will be asked to approve requests. A collection of singleUser, groupMembers,
	// requestorManager, internalSponsors, externalSponsors, or targetUserSponsors.
	PrimaryApprovers *[]SubjectSet `json:"primaryApprovers,omitempty"`
}

var _ json.Unmarshaler = &AccessPackageApprovalStage{}

func (s *AccessPackageApprovalStage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DurationBeforeAutomaticDenial   nullable.Type[string] `json:"durationBeforeAutomaticDenial,omitempty"`
		DurationBeforeEscalation        nullable.Type[string] `json:"durationBeforeEscalation,omitempty"`
		IsApproverJustificationRequired nullable.Type[bool]   `json:"isApproverJustificationRequired,omitempty"`
		IsEscalationEnabled             nullable.Type[bool]   `json:"isEscalationEnabled,omitempty"`
		ODataId                         *string               `json:"@odata.id,omitempty"`
		ODataType                       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DurationBeforeAutomaticDenial = decoded.DurationBeforeAutomaticDenial
	s.DurationBeforeEscalation = decoded.DurationBeforeEscalation
	s.IsApproverJustificationRequired = decoded.IsApproverJustificationRequired
	s.IsEscalationEnabled = decoded.IsEscalationEnabled
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageApprovalStage into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'EscalationApprovers' for 'AccessPackageApprovalStage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EscalationApprovers = &output
	}

	if v, ok := temp["fallbackEscalationApprovers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling FallbackEscalationApprovers into list []json.RawMessage: %+v", err)
		}

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'FallbackEscalationApprovers' for 'AccessPackageApprovalStage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.FallbackEscalationApprovers = &output
	}

	if v, ok := temp["fallbackPrimaryApprovers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling FallbackPrimaryApprovers into list []json.RawMessage: %+v", err)
		}

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'FallbackPrimaryApprovers' for 'AccessPackageApprovalStage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.FallbackPrimaryApprovers = &output
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
				return fmt.Errorf("unmarshaling index %d field 'PrimaryApprovers' for 'AccessPackageApprovalStage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PrimaryApprovers = &output
	}

	return nil
}
