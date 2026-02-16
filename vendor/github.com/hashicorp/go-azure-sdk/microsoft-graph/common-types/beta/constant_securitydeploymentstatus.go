package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeploymentStatus string

const (
	SecurityDeploymentStatus_Disconnected  SecurityDeploymentStatus = "disconnected"
	SecurityDeploymentStatus_NotConfigured SecurityDeploymentStatus = "notConfigured"
	SecurityDeploymentStatus_Outdated      SecurityDeploymentStatus = "outdated"
	SecurityDeploymentStatus_StartFailure  SecurityDeploymentStatus = "startFailure"
	SecurityDeploymentStatus_Syncing       SecurityDeploymentStatus = "syncing"
	SecurityDeploymentStatus_Unreachable   SecurityDeploymentStatus = "unreachable"
	SecurityDeploymentStatus_UpToDate      SecurityDeploymentStatus = "upToDate"
	SecurityDeploymentStatus_UpdateFailed  SecurityDeploymentStatus = "updateFailed"
	SecurityDeploymentStatus_Updating      SecurityDeploymentStatus = "updating"
)

func PossibleValuesForSecurityDeploymentStatus() []string {
	return []string{
		string(SecurityDeploymentStatus_Disconnected),
		string(SecurityDeploymentStatus_NotConfigured),
		string(SecurityDeploymentStatus_Outdated),
		string(SecurityDeploymentStatus_StartFailure),
		string(SecurityDeploymentStatus_Syncing),
		string(SecurityDeploymentStatus_Unreachable),
		string(SecurityDeploymentStatus_UpToDate),
		string(SecurityDeploymentStatus_UpdateFailed),
		string(SecurityDeploymentStatus_Updating),
	}
}

func (s *SecurityDeploymentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeploymentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeploymentStatus(input string) (*SecurityDeploymentStatus, error) {
	vals := map[string]SecurityDeploymentStatus{
		"disconnected":  SecurityDeploymentStatus_Disconnected,
		"notconfigured": SecurityDeploymentStatus_NotConfigured,
		"outdated":      SecurityDeploymentStatus_Outdated,
		"startfailure":  SecurityDeploymentStatus_StartFailure,
		"syncing":       SecurityDeploymentStatus_Syncing,
		"unreachable":   SecurityDeploymentStatus_Unreachable,
		"uptodate":      SecurityDeploymentStatus_UpToDate,
		"updatefailed":  SecurityDeploymentStatus_UpdateFailed,
		"updating":      SecurityDeploymentStatus_Updating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeploymentStatus(input)
	return &out, nil
}
