package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberProcessingStatus string

const (
	MultiTenantOrganizationMemberProcessingStatus_Failed     MultiTenantOrganizationMemberProcessingStatus = "failed"
	MultiTenantOrganizationMemberProcessingStatus_NotStarted MultiTenantOrganizationMemberProcessingStatus = "notStarted"
	MultiTenantOrganizationMemberProcessingStatus_Running    MultiTenantOrganizationMemberProcessingStatus = "running"
	MultiTenantOrganizationMemberProcessingStatus_Succeeded  MultiTenantOrganizationMemberProcessingStatus = "succeeded"
)

func PossibleValuesForMultiTenantOrganizationMemberProcessingStatus() []string {
	return []string{
		string(MultiTenantOrganizationMemberProcessingStatus_Failed),
		string(MultiTenantOrganizationMemberProcessingStatus_NotStarted),
		string(MultiTenantOrganizationMemberProcessingStatus_Running),
		string(MultiTenantOrganizationMemberProcessingStatus_Succeeded),
	}
}

func (s *MultiTenantOrganizationMemberProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationMemberProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationMemberProcessingStatus(input string) (*MultiTenantOrganizationMemberProcessingStatus, error) {
	vals := map[string]MultiTenantOrganizationMemberProcessingStatus{
		"failed":     MultiTenantOrganizationMemberProcessingStatus_Failed,
		"notstarted": MultiTenantOrganizationMemberProcessingStatus_NotStarted,
		"running":    MultiTenantOrganizationMemberProcessingStatus_Running,
		"succeeded":  MultiTenantOrganizationMemberProcessingStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberProcessingStatus(input)
	return &out, nil
}
