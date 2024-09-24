package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumeType string

const (
	VolumeType_FixedDataVolume       VolumeType = "fixedDataVolume"
	VolumeType_OperatingSystemVolume VolumeType = "operatingSystemVolume"
	VolumeType_RemovableDataVolume   VolumeType = "removableDataVolume"
)

func PossibleValuesForVolumeType() []string {
	return []string{
		string(VolumeType_FixedDataVolume),
		string(VolumeType_OperatingSystemVolume),
		string(VolumeType_RemovableDataVolume),
	}
}

func (s *VolumeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVolumeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVolumeType(input string) (*VolumeType, error) {
	vals := map[string]VolumeType{
		"fixeddatavolume":       VolumeType_FixedDataVolume,
		"operatingsystemvolume": VolumeType_OperatingSystemVolume,
		"removabledatavolume":   VolumeType_RemovableDataVolume,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VolumeType(input)
	return &out, nil
}
