package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskType string

const (
	DiskType_Hdd     DiskType = "hdd"
	DiskType_Ssd     DiskType = "ssd"
	DiskType_Unknown DiskType = "unknown"
)

func PossibleValuesForDiskType() []string {
	return []string{
		string(DiskType_Hdd),
		string(DiskType_Ssd),
		string(DiskType_Unknown),
	}
}

func (s *DiskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiskType(input string) (*DiskType, error) {
	vals := map[string]DiskType{
		"hdd":     DiskType_Hdd,
		"ssd":     DiskType_Ssd,
		"unknown": DiskType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskType(input)
	return &out, nil
}
