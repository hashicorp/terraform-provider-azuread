package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionOperationStatus string

const (
	ConnectionOperationStatus_Completed   ConnectionOperationStatus = "completed"
	ConnectionOperationStatus_Failed      ConnectionOperationStatus = "failed"
	ConnectionOperationStatus_Inprogress  ConnectionOperationStatus = "inprogress"
	ConnectionOperationStatus_Unspecified ConnectionOperationStatus = "unspecified"
)

func PossibleValuesForConnectionOperationStatus() []string {
	return []string{
		string(ConnectionOperationStatus_Completed),
		string(ConnectionOperationStatus_Failed),
		string(ConnectionOperationStatus_Inprogress),
		string(ConnectionOperationStatus_Unspecified),
	}
}

func (s *ConnectionOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionOperationStatus(input string) (*ConnectionOperationStatus, error) {
	vals := map[string]ConnectionOperationStatus{
		"completed":   ConnectionOperationStatus_Completed,
		"failed":      ConnectionOperationStatus_Failed,
		"inprogress":  ConnectionOperationStatus_Inprogress,
		"unspecified": ConnectionOperationStatus_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionOperationStatus(input)
	return &out, nil
}
