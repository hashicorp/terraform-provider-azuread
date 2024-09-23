package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataPolicyOperationStatus string

const (
	DataPolicyOperationStatus_Complete   DataPolicyOperationStatus = "complete"
	DataPolicyOperationStatus_Failed     DataPolicyOperationStatus = "failed"
	DataPolicyOperationStatus_NotStarted DataPolicyOperationStatus = "notStarted"
	DataPolicyOperationStatus_Running    DataPolicyOperationStatus = "running"
)

func PossibleValuesForDataPolicyOperationStatus() []string {
	return []string{
		string(DataPolicyOperationStatus_Complete),
		string(DataPolicyOperationStatus_Failed),
		string(DataPolicyOperationStatus_NotStarted),
		string(DataPolicyOperationStatus_Running),
	}
}

func (s *DataPolicyOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataPolicyOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataPolicyOperationStatus(input string) (*DataPolicyOperationStatus, error) {
	vals := map[string]DataPolicyOperationStatus{
		"complete":   DataPolicyOperationStatus_Complete,
		"failed":     DataPolicyOperationStatus_Failed,
		"notstarted": DataPolicyOperationStatus_NotStarted,
		"running":    DataPolicyOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataPolicyOperationStatus(input)
	return &out, nil
}
