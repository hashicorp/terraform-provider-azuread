package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateDeploymentStatus string

const (
	ManagedTenantsManagementTemplateDeploymentStatus_Completed  ManagedTenantsManagementTemplateDeploymentStatus = "completed"
	ManagedTenantsManagementTemplateDeploymentStatus_Failed     ManagedTenantsManagementTemplateDeploymentStatus = "failed"
	ManagedTenantsManagementTemplateDeploymentStatus_InProgress ManagedTenantsManagementTemplateDeploymentStatus = "inProgress"
	ManagedTenantsManagementTemplateDeploymentStatus_Ineligible ManagedTenantsManagementTemplateDeploymentStatus = "ineligible"
	ManagedTenantsManagementTemplateDeploymentStatus_Unknown    ManagedTenantsManagementTemplateDeploymentStatus = "unknown"
)

func PossibleValuesForManagedTenantsManagementTemplateDeploymentStatus() []string {
	return []string{
		string(ManagedTenantsManagementTemplateDeploymentStatus_Completed),
		string(ManagedTenantsManagementTemplateDeploymentStatus_Failed),
		string(ManagedTenantsManagementTemplateDeploymentStatus_InProgress),
		string(ManagedTenantsManagementTemplateDeploymentStatus_Ineligible),
		string(ManagedTenantsManagementTemplateDeploymentStatus_Unknown),
	}
}

func (s *ManagedTenantsManagementTemplateDeploymentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementTemplateDeploymentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementTemplateDeploymentStatus(input string) (*ManagedTenantsManagementTemplateDeploymentStatus, error) {
	vals := map[string]ManagedTenantsManagementTemplateDeploymentStatus{
		"completed":  ManagedTenantsManagementTemplateDeploymentStatus_Completed,
		"failed":     ManagedTenantsManagementTemplateDeploymentStatus_Failed,
		"inprogress": ManagedTenantsManagementTemplateDeploymentStatus_InProgress,
		"ineligible": ManagedTenantsManagementTemplateDeploymentStatus_Ineligible,
		"unknown":    ManagedTenantsManagementTemplateDeploymentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateDeploymentStatus(input)
	return &out, nil
}
