package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFileStatus string

const (
	GroupPolicyUploadedDefinitionFileStatus_Assigned          GroupPolicyUploadedDefinitionFileStatus = "assigned"
	GroupPolicyUploadedDefinitionFileStatus_Available         GroupPolicyUploadedDefinitionFileStatus = "available"
	GroupPolicyUploadedDefinitionFileStatus_None              GroupPolicyUploadedDefinitionFileStatus = "none"
	GroupPolicyUploadedDefinitionFileStatus_RemovalFailed     GroupPolicyUploadedDefinitionFileStatus = "removalFailed"
	GroupPolicyUploadedDefinitionFileStatus_RemovalInProgress GroupPolicyUploadedDefinitionFileStatus = "removalInProgress"
	GroupPolicyUploadedDefinitionFileStatus_UploadFailed      GroupPolicyUploadedDefinitionFileStatus = "uploadFailed"
	GroupPolicyUploadedDefinitionFileStatus_UploadInProgress  GroupPolicyUploadedDefinitionFileStatus = "uploadInProgress"
)

func PossibleValuesForGroupPolicyUploadedDefinitionFileStatus() []string {
	return []string{
		string(GroupPolicyUploadedDefinitionFileStatus_Assigned),
		string(GroupPolicyUploadedDefinitionFileStatus_Available),
		string(GroupPolicyUploadedDefinitionFileStatus_None),
		string(GroupPolicyUploadedDefinitionFileStatus_RemovalFailed),
		string(GroupPolicyUploadedDefinitionFileStatus_RemovalInProgress),
		string(GroupPolicyUploadedDefinitionFileStatus_UploadFailed),
		string(GroupPolicyUploadedDefinitionFileStatus_UploadInProgress),
	}
}

func (s *GroupPolicyUploadedDefinitionFileStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyUploadedDefinitionFileStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyUploadedDefinitionFileStatus(input string) (*GroupPolicyUploadedDefinitionFileStatus, error) {
	vals := map[string]GroupPolicyUploadedDefinitionFileStatus{
		"assigned":          GroupPolicyUploadedDefinitionFileStatus_Assigned,
		"available":         GroupPolicyUploadedDefinitionFileStatus_Available,
		"none":              GroupPolicyUploadedDefinitionFileStatus_None,
		"removalfailed":     GroupPolicyUploadedDefinitionFileStatus_RemovalFailed,
		"removalinprogress": GroupPolicyUploadedDefinitionFileStatus_RemovalInProgress,
		"uploadfailed":      GroupPolicyUploadedDefinitionFileStatus_UploadFailed,
		"uploadinprogress":  GroupPolicyUploadedDefinitionFileStatus_UploadInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyUploadedDefinitionFileStatus(input)
	return &out, nil
}
