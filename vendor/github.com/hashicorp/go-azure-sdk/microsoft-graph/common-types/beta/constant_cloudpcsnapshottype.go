package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSnapshotType string

const (
	CloudPCSnapshotType_Automatic CloudPCSnapshotType = "automatic"
	CloudPCSnapshotType_Manual    CloudPCSnapshotType = "manual"
)

func PossibleValuesForCloudPCSnapshotType() []string {
	return []string{
		string(CloudPCSnapshotType_Automatic),
		string(CloudPCSnapshotType_Manual),
	}
}

func (s *CloudPCSnapshotType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSnapshotType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSnapshotType(input string) (*CloudPCSnapshotType, error) {
	vals := map[string]CloudPCSnapshotType{
		"automatic": CloudPCSnapshotType_Automatic,
		"manual":    CloudPCSnapshotType_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSnapshotType(input)
	return &out, nil
}
