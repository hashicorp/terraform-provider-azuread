package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceLifecycleWorkflowProcessingStatus string

const (
	IdentityGovernanceLifecycleWorkflowProcessingStatus_Canceled            IdentityGovernanceLifecycleWorkflowProcessingStatus = "canceled"
	IdentityGovernanceLifecycleWorkflowProcessingStatus_Completed           IdentityGovernanceLifecycleWorkflowProcessingStatus = "completed"
	IdentityGovernanceLifecycleWorkflowProcessingStatus_CompletedWithErrors IdentityGovernanceLifecycleWorkflowProcessingStatus = "completedWithErrors"
	IdentityGovernanceLifecycleWorkflowProcessingStatus_Failed              IdentityGovernanceLifecycleWorkflowProcessingStatus = "failed"
	IdentityGovernanceLifecycleWorkflowProcessingStatus_InProgress          IdentityGovernanceLifecycleWorkflowProcessingStatus = "inProgress"
	IdentityGovernanceLifecycleWorkflowProcessingStatus_Queued              IdentityGovernanceLifecycleWorkflowProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceLifecycleWorkflowProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceLifecycleWorkflowProcessingStatus_Canceled),
		string(IdentityGovernanceLifecycleWorkflowProcessingStatus_Completed),
		string(IdentityGovernanceLifecycleWorkflowProcessingStatus_CompletedWithErrors),
		string(IdentityGovernanceLifecycleWorkflowProcessingStatus_Failed),
		string(IdentityGovernanceLifecycleWorkflowProcessingStatus_InProgress),
		string(IdentityGovernanceLifecycleWorkflowProcessingStatus_Queued),
	}
}

func (s *IdentityGovernanceLifecycleWorkflowProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceLifecycleWorkflowProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceLifecycleWorkflowProcessingStatus(input string) (*IdentityGovernanceLifecycleWorkflowProcessingStatus, error) {
	vals := map[string]IdentityGovernanceLifecycleWorkflowProcessingStatus{
		"canceled":            IdentityGovernanceLifecycleWorkflowProcessingStatus_Canceled,
		"completed":           IdentityGovernanceLifecycleWorkflowProcessingStatus_Completed,
		"completedwitherrors": IdentityGovernanceLifecycleWorkflowProcessingStatus_CompletedWithErrors,
		"failed":              IdentityGovernanceLifecycleWorkflowProcessingStatus_Failed,
		"inprogress":          IdentityGovernanceLifecycleWorkflowProcessingStatus_InProgress,
		"queued":              IdentityGovernanceLifecycleWorkflowProcessingStatus_Queued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceLifecycleWorkflowProcessingStatus(input)
	return &out, nil
}
