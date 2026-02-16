package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtensionOperationStatus string

const (
	IdentityGovernanceCustomTaskExtensionOperationStatus_Completed IdentityGovernanceCustomTaskExtensionOperationStatus = "completed"
	IdentityGovernanceCustomTaskExtensionOperationStatus_Failed    IdentityGovernanceCustomTaskExtensionOperationStatus = "failed"
)

func PossibleValuesForIdentityGovernanceCustomTaskExtensionOperationStatus() []string {
	return []string{
		string(IdentityGovernanceCustomTaskExtensionOperationStatus_Completed),
		string(IdentityGovernanceCustomTaskExtensionOperationStatus_Failed),
	}
}

func (s *IdentityGovernanceCustomTaskExtensionOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceCustomTaskExtensionOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceCustomTaskExtensionOperationStatus(input string) (*IdentityGovernanceCustomTaskExtensionOperationStatus, error) {
	vals := map[string]IdentityGovernanceCustomTaskExtensionOperationStatus{
		"completed": IdentityGovernanceCustomTaskExtensionOperationStatus_Completed,
		"failed":    IdentityGovernanceCustomTaskExtensionOperationStatus_Failed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceCustomTaskExtensionOperationStatus(input)
	return &out, nil
}
