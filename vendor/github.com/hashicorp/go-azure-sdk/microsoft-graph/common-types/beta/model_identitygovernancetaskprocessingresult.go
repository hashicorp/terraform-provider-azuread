package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceTaskProcessingResult{}

type IdentityGovernanceTaskProcessingResult struct {
	// The date time when taskProcessingResult execution ended. Value is null if task execution is still in
	// progress.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The date time when the taskProcessingResult was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Describes why the taskProcessingResult has failed.
	FailureReason nullable.Type[string] `json:"failureReason,omitempty"`

	ProcessingStatus *IdentityGovernanceLifecycleWorkflowProcessingStatus `json:"processingStatus,omitempty"`

	// The date time when taskProcessingResult execution started. Value is null if task execution has not yet
	// started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	StartedDateTime nullable.Type[string] `json:"startedDateTime,omitempty"`

	Subject *User                   `json:"subject,omitempty"`
	Task    *IdentityGovernanceTask `json:"task,omitempty"`

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

func (s IdentityGovernanceTaskProcessingResult) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceTaskProcessingResult{}

func (s IdentityGovernanceTaskProcessingResult) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceTaskProcessingResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceTaskProcessingResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceTaskProcessingResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.taskProcessingResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceTaskProcessingResult: %+v", err)
	}

	return encoded, nil
}
