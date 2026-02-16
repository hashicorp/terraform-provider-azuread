package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentDiagnosticResultType string

const (
	CloudPCAgentDiagnosticResultType_CommunicationUnhealthy CloudPCAgentDiagnosticResultType = "communicationUnhealthy"
	CloudPCAgentDiagnosticResultType_FunctionalityDefect    CloudPCAgentDiagnosticResultType = "functionalityDefect"
	CloudPCAgentDiagnosticResultType_Healthy                CloudPCAgentDiagnosticResultType = "healthy"
	CloudPCAgentDiagnosticResultType_UnknownError           CloudPCAgentDiagnosticResultType = "unknownError"
	CloudPCAgentDiagnosticResultType_VersionOutdated        CloudPCAgentDiagnosticResultType = "versionOutdated"
)

func PossibleValuesForCloudPCAgentDiagnosticResultType() []string {
	return []string{
		string(CloudPCAgentDiagnosticResultType_CommunicationUnhealthy),
		string(CloudPCAgentDiagnosticResultType_FunctionalityDefect),
		string(CloudPCAgentDiagnosticResultType_Healthy),
		string(CloudPCAgentDiagnosticResultType_UnknownError),
		string(CloudPCAgentDiagnosticResultType_VersionOutdated),
	}
}

func (s *CloudPCAgentDiagnosticResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAgentDiagnosticResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAgentDiagnosticResultType(input string) (*CloudPCAgentDiagnosticResultType, error) {
	vals := map[string]CloudPCAgentDiagnosticResultType{
		"communicationunhealthy": CloudPCAgentDiagnosticResultType_CommunicationUnhealthy,
		"functionalitydefect":    CloudPCAgentDiagnosticResultType_FunctionalityDefect,
		"healthy":                CloudPCAgentDiagnosticResultType_Healthy,
		"unknownerror":           CloudPCAgentDiagnosticResultType_UnknownError,
		"versionoutdated":        CloudPCAgentDiagnosticResultType_VersionOutdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAgentDiagnosticResultType(input)
	return &out, nil
}
