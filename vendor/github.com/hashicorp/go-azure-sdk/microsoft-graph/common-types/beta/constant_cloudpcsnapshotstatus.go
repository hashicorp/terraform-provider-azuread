package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSnapshotStatus string

const (
	CloudPCSnapshotStatus_Ready CloudPCSnapshotStatus = "ready"
)

func PossibleValuesForCloudPCSnapshotStatus() []string {
	return []string{
		string(CloudPCSnapshotStatus_Ready),
	}
}

func (s *CloudPCSnapshotStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSnapshotStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSnapshotStatus(input string) (*CloudPCSnapshotStatus, error) {
	vals := map[string]CloudPCSnapshotStatus{
		"ready": CloudPCSnapshotStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSnapshotStatus(input)
	return &out, nil
}
