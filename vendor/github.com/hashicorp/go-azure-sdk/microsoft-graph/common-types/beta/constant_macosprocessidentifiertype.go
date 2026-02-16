package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSProcessIdentifierType string

const (
	MacOSProcessIdentifierType_BundleId MacOSProcessIdentifierType = "bundleID"
	MacOSProcessIdentifierType_Path     MacOSProcessIdentifierType = "path"
)

func PossibleValuesForMacOSProcessIdentifierType() []string {
	return []string{
		string(MacOSProcessIdentifierType_BundleId),
		string(MacOSProcessIdentifierType_Path),
	}
}

func (s *MacOSProcessIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSProcessIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSProcessIdentifierType(input string) (*MacOSProcessIdentifierType, error) {
	vals := map[string]MacOSProcessIdentifierType{
		"bundleid": MacOSProcessIdentifierType_BundleId,
		"path":     MacOSProcessIdentifierType_Path,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSProcessIdentifierType(input)
	return &out, nil
}
