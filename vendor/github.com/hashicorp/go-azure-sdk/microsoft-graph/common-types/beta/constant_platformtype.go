package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlatformType string

const (
	PlatformType_Android    PlatformType = "android"
	PlatformType_IOS        PlatformType = "iOS"
	PlatformType_Linux      PlatformType = "linux"
	PlatformType_MacOS      PlatformType = "macOS"
	PlatformType_None       PlatformType = "none"
	PlatformType_Windows10  PlatformType = "windows10"
	PlatformType_Windows10X PlatformType = "windows10X"
)

func PossibleValuesForPlatformType() []string {
	return []string{
		string(PlatformType_Android),
		string(PlatformType_IOS),
		string(PlatformType_Linux),
		string(PlatformType_MacOS),
		string(PlatformType_None),
		string(PlatformType_Windows10),
		string(PlatformType_Windows10X),
	}
}

func (s *PlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlatformType(input string) (*PlatformType, error) {
	vals := map[string]PlatformType{
		"android":    PlatformType_Android,
		"ios":        PlatformType_IOS,
		"linux":      PlatformType_Linux,
		"macos":      PlatformType_MacOS,
		"none":       PlatformType_None,
		"windows10":  PlatformType_Windows10,
		"windows10x": PlatformType_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformType(input)
	return &out, nil
}
