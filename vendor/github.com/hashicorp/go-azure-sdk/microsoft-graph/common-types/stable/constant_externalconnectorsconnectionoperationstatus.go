package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsConnectionOperationStatus string

const (
	ExternalConnectorsConnectionOperationStatus_Completed   ExternalConnectorsConnectionOperationStatus = "completed"
	ExternalConnectorsConnectionOperationStatus_Failed      ExternalConnectorsConnectionOperationStatus = "failed"
	ExternalConnectorsConnectionOperationStatus_Inprogress  ExternalConnectorsConnectionOperationStatus = "inprogress"
	ExternalConnectorsConnectionOperationStatus_Unspecified ExternalConnectorsConnectionOperationStatus = "unspecified"
)

func PossibleValuesForExternalConnectorsConnectionOperationStatus() []string {
	return []string{
		string(ExternalConnectorsConnectionOperationStatus_Completed),
		string(ExternalConnectorsConnectionOperationStatus_Failed),
		string(ExternalConnectorsConnectionOperationStatus_Inprogress),
		string(ExternalConnectorsConnectionOperationStatus_Unspecified),
	}
}

func (s *ExternalConnectorsConnectionOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsConnectionOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsConnectionOperationStatus(input string) (*ExternalConnectorsConnectionOperationStatus, error) {
	vals := map[string]ExternalConnectorsConnectionOperationStatus{
		"completed":   ExternalConnectorsConnectionOperationStatus_Completed,
		"failed":      ExternalConnectorsConnectionOperationStatus_Failed,
		"inprogress":  ExternalConnectorsConnectionOperationStatus_Inprogress,
		"unspecified": ExternalConnectorsConnectionOperationStatus_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsConnectionOperationStatus(input)
	return &out, nil
}
